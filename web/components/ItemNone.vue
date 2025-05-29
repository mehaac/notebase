<script lang="ts" setup>
import type { ItemRecord } from '#pocketbase-imports'
import { computed } from 'vue'

const { item } = defineProps<{
  item: ItemRecord
}>()

const title = computed(() => {
  return item.frontmatter?.title ?? item.frontmatter?.summary ?? item.id
})
</script>

<template>
  <ItemsListCard
    :id="item.id"
    :title="title"
    :icon="'i-lucide-alert-circle'"
    :icon-variant="{
      color: 'warning',
    }"
  >
    <div class="flex gap-2 items-center py-2">
      <p class="text-sm text-dimmed">
        This Item have data but no component
      </p>
    </div>
  </ItemsListCard>
</template>
