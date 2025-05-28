<script lang="ts" setup>
import type { ItemRecord, TrackFrontmatter } from '#pocketbase-imports'

const { item, loading = false } = defineProps<{
  item: ItemRecord & { frontmatter: TrackFrontmatter }
  loading?: boolean
}>()

const emits = defineEmits<{
  change: [payload: { key: string, n: number }]
}>()
</script>

<template>
  <div class="flex  justify-between gap-3  rounded-lg">
    <div class="flex items-end  gap-3">
      <TrackIncrement
        :item="item"
        incr-key="season"
        :is-list="true"
        :loading="loading"
        @change="(payload) => emits('change', payload)"
      />

      <TrackIncrement
        :item="item"
        incr-key="episode"
        :is-list="true"
        :loading="loading"
        @change="(payload) => emits('change', payload)"
      />

      <UButton
        v-if="item.frontmatter.url"
        color="primary"
        variant="soft"
        icon="i-lucide-play"
        aria-label="Watch"
        :to="item.frontmatter.url"
        target="_blank"
        external
        label="Watch"
        class="text-center justify-center min-w-20 absolute right-2 top-2"
      />
    </div>
  </div>
</template>
