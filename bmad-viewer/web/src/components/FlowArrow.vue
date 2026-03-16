<script setup lang="ts">
defineProps<{
  color?: string
}>()
</script>

<template>
  <span
    class="text-xl mx-1 shrink-0"
    :style="{ color: color || 'var(--border)' }"
  >→</span>
</template>
