<script setup lang="ts">
import type { GroceriesFrontmatter, ItemRecord } from '#pocketbase-imports'
import { useActivitiesUpdateItemMutation } from '~/composables/queries'
import type { BaseItemEmits } from './BaseItem.vue'

const { item, compact } = defineProps<{ item: ItemRecord & { frontmatter: GroceriesFrontmatter }, compact?: boolean }>()
const { mutateAsync, asyncStatus } = useActivitiesUpdateItemMutation()

const emits = defineEmits<BaseItemEmits>()

async function handleChange(payload: ItemRecord) {
  await mutateAsync(payload)
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
