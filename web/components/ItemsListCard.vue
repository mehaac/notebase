<script lang="ts">
export interface ItemsListCardProps {
  title: string
  checked?: boolean
  icon: string
  id: string
  loading?: boolean
  disabled?: boolean
  ui?: {
    icon?: string
  }
  iconVariant?: {
    color?: 'primary' | 'secondary' | 'success' | 'warning' | 'danger' | 'info'
  }
}

export const iconTheme = {
  base: 'text-primary',
  variants: {
    color: {
      primary: 'text-primary',
      secondary: 'text-gray',
      success: 'text-green',
      warning: 'text-yellow',
      danger: 'text-red',
      info: 'text-blue',
    },
  },
  defaultVariants: {
    color: 'primary' as const,
  },
}
</script>

<script lang="ts" setup>
import { tv } from 'tailwind-variants'
import type { BaseItemEmits } from './BaseItem.vue'
import { computed } from 'vue'

const props = defineProps<ItemsListCardProps>()

const emits = defineEmits<BaseItemEmits>()

const iconClass = computed(() => {
  const tvInstance = tv(iconTheme)
  return tvInstance(props.iconVariant)
})
</script>

<template>
  <UCard>
    <slot name="header">
      <div class="flex items-center gap-2">
        <UCheckbox
          :loading="loading"
          :disabled="disabled || loading"
          :model-value="checked"
          @update:model-value="emits('done', id)"
        />
        <UIcon
          :name="icon"
          :class="[iconClass, props.ui?.icon]"
        />
        <ULink
          :to="`/items/${id}`"
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
