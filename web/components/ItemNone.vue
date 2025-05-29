<script lang="ts" setup>
import type { ItemRecord } from '#pocketbase-imports'

const { item } = defineProps<{
  item: ItemRecord
}>()
</script>

<template>
  <ItemsListCard
    :id="item.id"
    :title="item.slug"
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
