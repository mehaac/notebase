<script setup lang="ts" generic="T extends ItemRecord & { frontmatter: TrackFrontmatter }">
import { onMounted, ref, watchDebounced } from '#imports'
import type { ItemRecord, TrackFrontmatter } from '#pocketbase-imports'
import { useActivitiesUpdateItemMutation } from '~/composables/queries'

const { item, incrKey } = defineProps<{ item: T, incrKey: string }>()
const num = ref(0)
const { mutateAsync } = useActivitiesUpdateItemMutation()

const incr = async (n: number) => {
  num.value += n
}

watchDebounced(num, async (newValue, oldValue) => {
  if (oldValue === 0) return
  const newItem = item
  newItem.frontmatter[incrKey] = newValue
  await mutateAsync(newItem)
}, { debounce: 800 })

onMounted(() => {
  num.value = item.frontmatter[incrKey] as number
})
</script>

<template>
  <UButtonGroup class="mr-2">
    <UButton
      color="warning"
      variant="subtle"
      icon="material-symbols:remove"
      @click="() => incr(-1)"
    />
    <UInput
      v-model="num"
      :placeholder="incrKey"
      color="neutral"
      variant="outline"
      type="number"
      class="w-24"
    />

    <UButton
      color="success"
      variant="subtle"
      icon="material-symbols:add"
      @click="() => incr(1)"
    />
  </UButtonGroup>
</template>
