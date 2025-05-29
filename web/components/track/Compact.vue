<script lang="ts" setup>
import type { ItemRecord, TrackFrontmatter } from '#pocketbase-imports'
import type { BaseItemEmits } from '../BaseItem.vue'

const { item, loading = false } = defineProps<{
  item: ItemRecord & { frontmatter: TrackFrontmatter }
  loading?: boolean
}>()

const emits = defineEmits<BaseItemEmits>()
</script>

<template>
  <ItemsListCard
    :item="item"
    :icon="'i-lucide-tv'"
    :loading="loading"
    @toggle-completed="emits('updateFrontmatter', item)"
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
          @update-frontmatter="(payload) => emits('updateFrontmatter', payload)"
        />

        <TrackIncrement
          :item="item"
          incr-key="episode"
          :is-list="true"
          :loading="loading"
          @update-frontmatter="(payload) => emits('updateFrontmatter', payload)"
        />
      </div>
    </div>
  </ItemsListCard>
</template>
