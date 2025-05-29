<script lang="ts">
export const iconColors = {
  primary: 'text-primary',
  secondary: 'text-gray-500',
  success: 'text-green-500',
  warning: 'text-yellow-500',
  danger: 'text-red-500',
}

export interface ItemsListCardProps {
  item: ItemRecord
  icon: string
  loading?: boolean
  disabled?: boolean
  iconColor?: keyof typeof iconColors
}
</script>

<script lang="ts" setup>
import { computed } from 'vue'
import type { ItemRecord } from '#pocketbase-imports'

const { item, iconColor = 'primary' } = defineProps<ItemsListCardProps>()

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

function toggleItem(item: ItemRecord) {
  const frontmatter = item.frontmatter || {}
  if (frontmatter?.completed) {
    frontmatter.completed = ''
  }
  else {
    frontmatter.completed = new Date().toISOString().split('.')[0]!
  }

  return {
    ...item,
    frontmatter,
  }
}

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
    <slot />
  </UCard>
</template>
