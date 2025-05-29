<script lang="ts">
export const iconColors = {
  primary: 'text-primary',
  secondary: 'text-gray-500',
  success: 'text-green-500',
  warning: 'text-yellow-500',
  danger: 'text-red-500',
}

export interface ItemCardProps {
  item: ItemRecord
  icon?: string
  compact?: boolean
  loading?: boolean
  disabled?: boolean
  iconColor?: keyof typeof iconColors
}
</script>

<script lang="ts" setup>
import { computed } from 'vue'
import type { ItemRecord } from '#pocketbase-imports'
import { toggleItem } from '#imports'

const { item, iconColor = 'primary', icon = 'i-lucide-notebook-pen' } = defineProps<ItemCardProps>()

const emits = defineEmits<{
  'toggle-completed': [item: ItemRecord]
}>()

const iconClass = computed(() => {
  return iconColors[iconColor]
})

const title = computed(() => item?.frontmatter?.title || item?.frontmatter?.summary || item.slug)

const checked = computed(() => {
  return Boolean(item?.frontmatter?.completed)
})

function handleToggleDone(item: ItemRecord) {
  emits('toggle-completed', toggleItem(item))
}
</script>

<template>
  <UCard>
    <slot name="header">
      <div class="flex items-center gap-2">
        <UCheckbox
          :loading="loading"
          :disabled="disabled || loading"
          :model-value="checked"
          @update:model-value="handleToggleDone(item)"
        />
        <UIcon
          :name="icon"
          :class="[
            iconClass,
          ]"
        />
        <ULink
          :to="`/items/${item.id}`"
          class="text-lg font-bold truncate"
        >
          {{ title }}
        </ULink>

        <div class="ml-auto">
          <slot name="actions" />
        </div>
      </div>
    </slot>
    <div class="py-4">
      <slot />
    </div>
  </UCard>
</template>
