<script lang="ts">
export type BaseItemEmits = {
  done: [id: string]
  change: [payload: { key: string, n: number }]
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  action: [payload: { action: string, id: string, data?: any }]
}

export interface BaseItemProps {
  item: ItemRecord
  compact?: boolean
  loading?: boolean
  disabled?: boolean
}
</script>

<script lang="ts" setup>
import { computed, resolveDynamicComponent } from 'vue'
import type { ItemRecord } from '#pocketbase-imports'

const props = withDefaults(defineProps<BaseItemProps>(), {
  compact: false,
  loading: false,
  disabled: false,
})

const emits = defineEmits<BaseItemEmits>()

const itemType = computed(() => {
  return props.item?.frontmatter?.type ?? 'none'
})

const itemComponent = computed(() => {
  const componentName = `item-${itemType.value}`
  try {
    return resolveDynamicComponent(componentName)
  }
  catch (error) {
    console.warn(`Component for type "${itemType.value}" not found, falling back to default`)
    return resolveDynamicComponent('item-none')
  }
})

const hasValidItem = computed(() => {
  return props.item && props.item.id && props.item.frontmatter
})

const handleDone = (id: string) => {
  emits('done', id)
}

const handleChange = (payload: { key: string, n: number }) => {
  emits('change', payload)
}

// eslint-disable-next-line @typescript-eslint/no-explicit-any
const handleAction = (payload: { action: string, id: string, data?: any }) => {
  emits('action', payload)
}
</script>

<template>
  <ItemEmpty
    v-if="!hasValidItem"
    :item-id="item.id"
  />

  <component
    :is="itemComponent"
    v-else
    :item="item"
    :compact="compact"
    :loading="loading"
    :disabled="disabled"
    @done="handleDone"
    @change="handleChange"
    @action="handleAction"
  />
</template>
