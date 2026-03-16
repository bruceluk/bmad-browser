<script setup lang="ts">
defineProps<{
  role: string
  roleColor: string
  icon: string
  label: string
  description: string
  stepCount: number
  totalDuration: string
}>()

defineEmits<{
  click: []
}>()
</script>

<template>
  <div
    class="w-[300px] p-7 rounded-2xl cursor-pointer transition-all duration-150"
    style="background-color: var(--surface); border: 1px solid var(--border);"
    @mouseenter="(e) => {
      const el = e.currentTarget as HTMLElement
      el.style.transform = 'translateY(-4px)'
      el.style.borderColor = roleColor
      el.style.boxShadow = `0 8px 32px ${roleColor}26`
    }"
    @mouseleave="(e) => {
      const el = e.currentTarget as HTMLElement
      el.style.transform = ''
      el.style.borderColor = 'var(--border)'
      el.style.boxShadow = ''
    }"
    @click="$emit('click')"
  >
    <div
      class="w-12 h-12 rounded-xl flex items-center justify-center text-2xl mb-4"
      :style="{ backgroundColor: roleColor + '26', color: roleColor }"
    >
      {{ icon }}
    </div>
    <h3 class="text-lg font-semibold mb-2">{{ label }}</h3>
    <p class="text-sm opacity-60 leading-relaxed mb-4">{{ description }}</p>
    <div class="text-xs opacity-50 flex items-center gap-3">
      <span>📄 {{ stepCount }} 个步骤</span>
      <span>⏱ {{ totalDuration }}</span>
    </div>
  </div>
</template>
