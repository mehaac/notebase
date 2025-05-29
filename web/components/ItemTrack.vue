<script setup lang="ts" generic="T extends ItemRecord & { frontmatter: TrackFrontmatter }">
import type { ItemRecord, TrackFrontmatter } from '#pocketbase-imports'
import type { BaseItemEmits } from './BaseItem.vue'

const { item, compact, loading } = defineProps<{ item: T, compact?: boolean, loading?: boolean }>()

const emits = defineEmits<BaseItemEmits>()
</script>

<template>
  <TrackCompact
    v-if="compact"
    :item="item"
    :loading="loading"
    @update-frontmatter="(payload) => emits('updateFrontmatter', payload)"
  />
  <TrackDetailed
    v-else
    :item="item"
    :loading="loading"
    @update-frontmatter="(payload) => emits('updateFrontmatter', payload)"
  />
</template>
