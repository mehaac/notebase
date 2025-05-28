<script lang="ts" setup>
import type { BaseItemEmits } from './BaseItem.vue'

defineProps<{
  title: string
  checked?: boolean
  icon: string
  id: string
  loading?: boolean
  disabled?: boolean
}>()

const emits = defineEmits<BaseItemEmits>()
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
          class="text-primary-500"
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
