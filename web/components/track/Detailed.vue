<script lang="ts" setup>
import type { ItemRecord, TrackFrontmatter } from '#pocketbase-imports'
import { computed } from 'vue'
import type { BaseItemEmits } from '../BaseItem.vue'
import { formatDateShort } from '#imports'

const {
  item,
  loading = false,
} = defineProps<{
  item: ItemRecord & { frontmatter: TrackFrontmatter }
  loading?: boolean
}>()

const emits = defineEmits<BaseItemEmits>()

const formattedDate = computed(() => {
  return item.frontmatter.next_episode
    ? formatDateShort(new Date(item.frontmatter.next_episode))
    : null
})

const formattedCreatedDate = computed(() => {
  return item.frontmatter.created && typeof item.frontmatter.created === 'string'
    ? formatDateShort(new Date(item.frontmatter.created))
    : null
})
</script>

<template>
  <div class="flex flex-col gap-6">
    <div class="flex-1">
      <div class="flex items-center mb-1">
        <UIcon
          name="i-lucide-tv-2"
          class="mr-2 text-primary-500"
        />
        <h1 class="text-2xl font-bold">
          {{ item.frontmatter.summary }}
        </h1>
      </div>

      <div class="mb-4 flex flex-wrap items-center gap-2">
        <UBadge
          v-if="item.frontmatter.rating"
          color="warning"
        >
          {{ item.frontmatter.rating }}
        </UBadge>
        <UBadge
          v-if="!item.frontmatter.completed"
          color="primary"
        >
          Ongoing
        </UBadge>
        <UBadge
          v-else
          color="success"
        >
          Completed
        </UBadge>

        <span
          v-if="formattedCreatedDate"
          class="text-xs"
        >Added: {{ formattedCreatedDate }}</span>
      </div>

      <div class="flex flex-col gap-4 mb-6">
        <div class="flex items-center">
          <UIcon
            name="i-lucide-play"
            class="mr-2 text-primary-500"
          />
          <ULink
            :to="item.frontmatter.url"
            target="_blank"
            external
            class="hover:underline"
          >
            Watch Online
          </ULink>
        </div>

        <div
          v-if="formattedDate"
          class="flex items-center"
        >
          <UIcon
            name="i-lucide-calendar"
            class="mr-2 text-primary-500"
          />
          <span class="text-sm">Next episode: {{ formattedDate }}</span>
        </div>

        <div
          v-if="item.frontmatter.tags && item.frontmatter.tags.length"
          class="flex items-start md:col-span-2"
        >
          <UIcon
            name="i-lucide-tags"
            class="mr-2 mt-0.5 text-primary-500"
          />
          <div class="flex flex-wrap gap-1">
            <UBadge
              v-for="tag in item.frontmatter.tags"
              :key="tag"
              color="neutral"
              variant="subtle"
              size="xs"
            >
              {{ tag }}
            </UBadge>
          </div>
        </div>

        <div
          v-if="item.frontmatter.rating"
          class="flex items-center md:col-span-2"
        >
          <UIcon
            name="i-lucide-star"
            class="mr-2 text-primary-500"
          />
          <div class="flex items-center">
            <span class="ml-2 text-sm">{{ item.frontmatter.rating }}</span>
          </div>
        </div>
      </div>
    </div>

    <div class="rounded-lg p-4 shadow-sm flex flex-col">
      <div class="flex items-center mb-4">
        <UIcon
          name="i-lucide-activity"
          class="mr-2 text-primary-500"
        />
        <h2 class="text-lg font-medium">
          Track Progress
        </h2>
      </div>

      <div class="space-y-4">
        <div class="flex items-center">
          <span class="w-20">Season:</span>
          <TrackIncrement
            :item="item"
            incr-key="season"
            :loading="loading"
            @update-frontmatter="(payload) => emits('updateFrontmatter', payload)"
          />
        </div>

        <div class="flex items-center">
          <span class="w-20">Episode:</span>
          <TrackIncrement
            :item="item"
            incr-key="episode"
            :loading="loading"
            @update-frontmatter="(payload) => emits('updateFrontmatter', payload)"
          />
        </div>
      </div>
    </div>
  </div>
</template>
