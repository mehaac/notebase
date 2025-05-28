<script lang="ts" setup>
import type { ItemRecord } from '#pocketbase-imports'
import { useActivitiesUpdateItemMutation } from '~/composables/queries'

const { items } = defineProps<{ items: ItemRecord[] }>()

const { mutateAsync } = useActivitiesUpdateItemMutation()

async function toggleItem(itemId: string) {
  const item = items.find(item => item.id === itemId)
  if (!item) return
  const frontmatter = item.frontmatter || {}
  if (frontmatter?.completed) {
    frontmatter.completed = ''
  }
  else {
    frontmatter.completed = new Date().toISOString().split('.')[0]!
  }

  await mutateAsync({
    ...item,
    frontmatter,
  })
}
</script>

<template>
  <ul
    class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4"
  >
    <LazyBaseItem
      v-for="item in items"
      :key="item.id"
      :item="item"
      compact
      @done="(val: string) => toggleItem(val)"
    />
  </ul>
</template>
