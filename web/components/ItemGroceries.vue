<script setup lang="ts" generic="T extends ItemRecord & { frontmatter: { title: string, created: string, items: Array<{ name: string, done: boolean }> } }">
import type { ItemRecord } from '#pocketbase-imports'
import { useActivitiesUpdateItemMutation } from '~/composables/queries'
import type { BaseItemEmits } from './BaseItem.vue'

const { item, compact } = defineProps<{ item: T, compact?: boolean }>()
const { mutateAsync, asyncStatus } = useActivitiesUpdateItemMutation()

const emits = defineEmits<BaseItemEmits>()

async function handleChange(payload: { index: number, done: boolean, name?: string, action: 'toggle' | 'add' | 'remove' }) {
  const newItem = { ...item }
  const { index, done, name, action } = payload

  switch (action) {
    case 'toggle':
      if (newItem.frontmatter.items[index]) {
        newItem.frontmatter.items[index].done = done
      }
      break
    case 'add':
      if (name) {
        newItem.frontmatter.items.push({ name, done: false })
      }
      break
    case 'remove':
      newItem.frontmatter.items = newItem.frontmatter.items.filter((_, i) => i !== index)
      break
  }

  await mutateAsync(newItem)
}
</script>

<template>
  <GroceriesCompact
    v-if="compact"
    :item="item"
    :loading="asyncStatus === 'loading'"
    @done="emits('done', item.id)"
  />
  <GroceriesDetailed
    v-else
    :item="item"
    :loading="asyncStatus === 'loading'"
    @change="handleChange"
    @done="emits('done', item.id)"
  />
</template>
