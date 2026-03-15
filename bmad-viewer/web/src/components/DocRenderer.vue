<script setup lang="ts">
import { computed } from 'vue'
import MarkdownIt from 'markdown-it'

const props = defineProps<{
  content: string
}>()

const md = new MarkdownIt()

const renderedHtml = computed(() => md.render(props.content))
</script>

<template>
  <div
    class="prose prose-invert max-w-[800px] mx-auto px-12"
    v-html="renderedHtml"
  />
</template>
