<script lang="ts" setup>
import type { ItemRecord } from '#pocketbase-imports'
import type { BaseItemEmits } from './BaseItem.vue'

const { item, loading = false } = defineProps<{
  item: ItemRecord
  loading?: boolean
}>()

const emits = defineEmits<BaseItemEmits>()
</script>

<template>
  <ItemsListCard
    :item="item"
    :icon="'i-lucide-alert-circle'"
    :icon-variant="{
      color: 'warning',
    }"
    :loading="loading"
    @toggle-completed="(payload) => emits('updateFrontmatter', payload)"
  >
    <div class="flex gap-2 items-center py-2">
      <p class="text-sm text-dimmed">
        Empty item
      </p>
    </div>
  </ItemsListCard>
</template>
