<script setup lang="ts">
import type { GroceriesFrontmatter, ItemRecord } from '#pocketbase-imports'

import type { BaseItemEmits } from './BaseItem.vue'

const { item, compact, loading } = defineProps<{ item: ItemRecord & { frontmatter: GroceriesFrontmatter }, compact?: boolean, loading?: boolean }>()

const emits = defineEmits<BaseItemEmits>()

async function handleUpdateFrontmatter(payload: ItemRecord) {
  emits('updateFrontmatter', payload)
}
</script>

<template>
  <GroceriesCompact
    v-if="compact"
    :item="item"
    :loading="loading"
    @update-frontmatter="handleUpdateFrontmatter"
  />
  <GroceriesDetailed
    v-else
    :item="item"
    :loading="loading"
    @update-frontmatter="handleUpdateFrontmatter"
  />
</template>
