<script setup lang="ts" generic="T extends ItemRecord & { frontmatter: TrackFrontmatter }">
import { onMounted, ref, watchDebounced } from '#imports'
import type { ItemRecord, TrackFrontmatter } from '#pocketbase-imports'

const {
  item,
  incrKey,
  isList = false,
  loading = false,
  disabled = false,
} = defineProps<{
  item: T
  incrKey: string
  isList?: boolean
  loading?: boolean
  disabled?: boolean
}>()

const emits = defineEmits<{
  change: [payload: { key: string, n: number }]
}>()

const num = ref(0)

watchDebounced(num, async (newValue, oldValue) => {
  if (oldValue === 0) return
  emits('change', { key: incrKey, n: newValue })
}, { debounce: 800 })

onMounted(() => {
  num.value = item.frontmatter[incrKey] as number
})
</script>

<template>
  <div
    class="flex flex-col items-center gap-2 px-4"
  >
    <label
      v-if="isList"
      class="text-sm text-start w-full font-medium text-(--ui-text-dimmed) capitalize"
    >
      {{ incrKey }}
    </label>
    <UInputNumber
      v-model="num"
      :loading="loading"
      :disabled="loading || disabled"
      :placeholder="incrKey"
      color="neutral"
      variant="outline"
      :min="0"
      increment-icon="i-lucide-arrow-right"
      decrement-icon="i-lucide-arrow-left"
      :increment="{
        color: 'info',
        variant: 'outline',
        size: 'xs',
      }"
      :decrement="{
        color: 'info',
        variant: 'outline',
        size: 'xs',
      }"
    />
  </div>
</template>
