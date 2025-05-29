<script setup lang="ts">
import type { GroceriesFrontmatter, ItemRecord } from '#pocketbase-imports'

import type { BaseItemEmits } from './BaseItem.vue'

const { item, compact, loading } = defineProps<{ item: ItemRecord & { frontmatter: GroceriesFrontmatter }, compact?: boolean, loading?: boolean }>()

const emits = defineEmits<BaseItemEmits>()
</script>

<template>
  <GroceriesCompact
    v-if="compact"
    :item="item"
    :loading="loading"
    @update-frontmatter="(payload) => emits('updateFrontmatter', payload)"
  />
  <GroceriesDetailed
    v-else
    :item="item"
    :loading="loading"
    @update-frontmatter="(payload) => emits('updateFrontmatter', payload)"
  />
</template>
