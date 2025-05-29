<script lang="ts" setup>
import { computed } from 'vue'
import type { GroceriesFrontmatter, ItemRecord } from '#pocketbase-imports'
import type { BaseItemEmits } from '../BaseItem.vue'

const { item } = defineProps<{
  item: ItemRecord & { frontmatter: GroceriesFrontmatter }
  loading?: boolean
}>()

const emits = defineEmits<BaseItemEmits>()

const completedItems = computed(() => {
  return item.frontmatter.checklist.filter(item => item.done).length
})

const totalItems = computed(() => {
  return item.frontmatter.checklist.length
})
</script>

<template>
  <ItemCard
    :item="item"
    :icon="'i-lucide-shopping-cart'"
    :loading="loading"
    compact
    @toggle-completed="(payload) => emits('updateFrontmatter', payload)"
  >
    <template #actions>
      <UBadge
        color="primary"
        variant="subtle"
        class="ml-auto"
      >
        {{ completedItems }}/{{ totalItems }}
      </UBadge>
    </template>

    <div class="text-sm text-dimmed">
      <div
        v-if="item.frontmatter.checklist.length > 0"
        class="text-xs"
      >
        {{ item.frontmatter.checklist.slice(0, 3).map(item => item.name).join(', ') }}
        <template v-if="item.frontmatter.checklist.length > 3">
          ...
        </template>
      </div>
      <div
        v-else
        class="text-xs text-dimmed"
      >
        Empty list
      </div>
    </div>
  </ItemCard>
</template>
