<script setup lang="ts" generic="T extends ItemRecord & { frontmatter: TrackFrontmatter }">
import type { ItemRecord, TrackFrontmatter } from '#pocketbase-imports'

import { useActivitiesUpdateItemMutation } from '~/composables/queries'

const { item, compact } = defineProps<{ item: T, compact?: boolean }>()
const { mutateAsync, asyncStatus } = useActivitiesUpdateItemMutation()

async function handleChange(payload: { key: string, n: number }) {
  const newItem = { ...item }
  newItem.frontmatter[payload.key] = payload.n
  await mutateAsync(newItem)
}
</script>

<template>
  <TrackCompact
    v-if="compact"
    :item="item"
    :loading="asyncStatus === 'loading'"
    @change="handleChange"
  />
  <TrackDetailed
    v-else
    :item="item"
    :loading="asyncStatus === 'loading'"
    @change="handleChange"
  />
</template>
