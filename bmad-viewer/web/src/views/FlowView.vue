<script setup lang="ts">
import { ref, computed, onMounted, watch, nextTick } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { fetchRoles, fetchDocuments, fetchDocument } from '@/api/client'
import type { RoleFlow, WorkflowStep, DocumentSummary, Document } from '@/types'
import RoleTab from '@/components/RoleTab.vue'
import FlowNode from '@/components/FlowNode.vue'
import FlowArrow from '@/components/FlowArrow.vue'
import DocMeta from '@/components/DocMeta.vue'
import DocRenderer from '@/components/DocRenderer.vue'

const route = useRoute()
const router = useRouter()

const roles = ref<RoleFlow[]>([])
const allDocs = ref<DocumentSummary[]>([])
const currentRole = ref(route.params.role as string)
const activeStep = ref<WorkflowStep | null>(null)
const visitedSteps = ref(new Set<string>())
const currentDoc = ref<Document | null>(null)
const docLoading = ref(false)
const isLoading = ref(true)
const docArea = ref<HTMLElement>()

const currentRoleFlow = computed(() =>
  roles.value.find(r => r.role === currentRole.value)
)

onMounted(async () => {
  try {
    const [rolesData, docsData] = await Promise.all([fetchRoles(), fetchDocuments()])
    roles.value = rolesData
    allDocs.value = docsData
  } catch (e) {
    console.error(e)
  } finally {
    isLoading.value = false
  }
})

watch(() => route.params.role, (newRole) => {
  if (newRole && newRole !== currentRole.value) {
    currentRole.value = newRole as string
    activeStep.value = null
    currentDoc.value = null
  }
})

function switchRole(role: string) {
  currentRole.value = role
  activeStep.value = null
  currentDoc.value = null
  router.replace({ name: 'flow', params: { role } })
}

function goHome() {
  router.push({ name: 'home' })
}

async function selectStep(step: WorkflowStep) {
  if (activeStep.value) {
    visitedSteps.value.add(activeStep.value.code)
  }
  activeStep.value = step
  router.replace({ name: 'flow', params: { role: currentRole.value, step: step.code } })

  // Find matching document
  const docPath = findDocumentForStep(step, allDocs.value)
  if (docPath) {
    docLoading.value = true
    try {
      currentDoc.value = await fetchDocument(docPath)
    } catch (e) {
      currentDoc.value = null
      console.error(e)
    } finally {
      docLoading.value = false
    }
  } else {
    currentDoc.value = null
  }

  nextTick(() => docArea.value?.scrollIntoView({ behavior: 'smooth' }))
}

function findDocumentForStep(step: WorkflowStep, docs: DocumentSummary[]): string | null {
  if (!step.outputs || step.outputs.length === 0) return null

  for (const output of step.outputs) {
    const keywords = output.toLowerCase().split(/\s+/)
    const match = docs.find(doc => {
      const target = (doc.path + ' ' + doc.title).toLowerCase()
      return keywords.some(kw => kw.length > 2 && target.includes(kw))
    })
    if (match) return match.path
  }
  return null
}
</script>

<template>
  <div class="min-h-screen flex flex-col" :class="{ 'cursor-wait': isLoading || docLoading }">
    <!-- Header + Tabs -->
    <div class="px-8 py-4" style="border-bottom: 1px solid var(--border);">
      <div class="flex items-center justify-between mb-3">
        <button
          class="text-lg font-bold cursor-pointer hover:opacity-80 transition-opacity duration-150"
          @click="goHome"
        >
          BMAD Viewer
        </button>
      </div>
      <RoleTab
        v-if="roles.length > 0"
        :roles="roles"
        :current-role="currentRole"
        @select="switchRole"
      />
    </div>

    <!-- Flow Chart Area -->
    <div
      v-if="currentRoleFlow"
      class="px-8 py-6 flex items-center justify-center overflow-x-auto"
      style="border-bottom: 1px solid var(--border);"
    >
      <div class="flex items-center gap-0">
        <!-- Upstream nodes -->
        <template v-if="currentRoleFlow.upstream.length > 0">
          <span class="text-[11px] mr-2" style="color: var(--text-secondary); writing-mode: vertical-lr; letter-spacing: 2px;">上游</span>
          <template v-for="(step, i) in currentRoleFlow.upstream" :key="'up-' + step.code">
            <FlowNode
              :step="step"
              :role-color="currentRoleFlow.roleColor"
              :is-upstream="true"
            />
            <FlowArrow v-if="i < currentRoleFlow.upstream.length - 1" />
          </template>
          <FlowArrow />
        </template>

        <!-- Core steps -->
        <template v-for="(step, i) in currentRoleFlow.steps" :key="step.code">
          <FlowNode
            :step="step"
            :role-color="currentRoleFlow.roleColor"
            :is-active="activeStep?.code === step.code"
            :is-visited="visitedSteps.has(step.code)"
            @click="selectStep(step)"
          />
          <FlowArrow
            v-if="i < currentRoleFlow.steps.length - 1"
            :color="currentRoleFlow.roleColor"
          />
        </template>

        <!-- Downstream nodes -->
        <template v-if="currentRoleFlow.downstream.length > 0">
          <FlowArrow />
          <template v-for="(step, i) in currentRoleFlow.downstream" :key="'down-' + step.code">
            <FlowNode
              :step="step"
              :role-color="currentRoleFlow.roleColor"
              :is-downstream="true"
            />
            <FlowArrow v-if="i < currentRoleFlow.downstream.length - 1" />
          </template>
          <span class="text-[11px] ml-2" style="color: var(--text-secondary); writing-mode: vertical-lr; letter-spacing: 2px;">下游</span>
        </template>
      </div>
    </div>

    <!-- Document Area -->
    <div ref="docArea" class="flex-1">
      <!-- No step selected -->
      <div v-if="!activeStep" class="flex items-center justify-center h-64">
        <p class="opacity-40">点击上方流程节点查看详情</p>
      </div>

      <!-- Step selected, document loaded -->
      <div v-else class="max-w-[800px] mx-auto px-12 py-8">
        <DocMeta
          v-if="currentRoleFlow"
          :step="activeStep"
          :role-color="currentRoleFlow.roleColor"
        />

        <div v-if="currentDoc">
          <DocRenderer :content="currentDoc.content" />
        </div>
        <div v-else-if="!docLoading" class="text-center py-12 opacity-40">
          该步骤暂无产出文档
        </div>
      </div>
    </div>
  </div>
</template>
