<script lang="ts" setup>
import type { ItemRecord, TrackFrontmatter } from '#pocketbase-imports'
import { computed } from 'vue'

const { item, loading = false } = defineProps<{
  item: ItemRecord & { frontmatter: TrackFrontmatter }
  loading?: boolean
}>()

const checked = computed(() => {
  return Boolean(item.frontmatter?.completed)
})

const emits = defineEmits<{
  done: [id: string]
  change: [payload: { key: string, n: number }]
}>()

const title = computed(() =>
  (item.frontmatter.title || item.frontmatter.summary || 'None'),
)
</script>

<template>
  <ItemsListCard
    :id="item.id"
    :title="title"
    :icon="'i-lucide-tv'"
    :checked="checked"
    :loading="loading"
    @done="emits('done', item.id)"
  >
    <template #actions>
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
      />
    </template>
    <div class="flex flex-col gap-4">
      <div class="flex gap-3">
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
      </div>
    </div>
  </ItemsListCard>
</template>
