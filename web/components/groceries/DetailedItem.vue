<script lang="ts" setup>
import type { GroceriesItem } from '#pocketbase-imports'

defineProps<{
  groceryItem: GroceriesItem
  loading?: boolean
}>()

const emits = defineEmits<{
  toggle: [id: string]
  remove: [id: string]
}>()

function toggleItem(id: string) {
  emits('toggle', id)
}

function removeItem(id: string) {
  emits('remove', id)
}
</script>

<template>
  <div

    :class="[
      'flex items-center gap-3 p-3 rounded-lg border border-(--ui-border)',
      groceryItem.done
        ? ''
        : '',
    ]"
  >
    <UCheckbox
      :model-value="groceryItem.done"
      :disabled="loading"
      @change="toggleItem(groceryItem.name)"
    />

    <span
      :class="[
        'flex-1',
        groceryItem.done ? 'line-through text-dimmed' : '',
      ]"
      class="truncate font-medium"
    >
      {{ groceryItem.name }}
    </span>

    <UButton
      color="neutral"
      variant="ghost"
      icon="i-lucide-x"
      size="xs"
      :disabled="loading"
      @click="removeItem(groceryItem.name)"
    />
  </div>
</template>
