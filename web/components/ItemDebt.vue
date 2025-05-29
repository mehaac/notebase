<script lang="ts">
import { computed } from 'vue'
import { clamp } from '@vueuse/core'
import type { ItemRecord, DebtFrontmatter } from '#pocketbase-imports'
import type { BaseItemEmits } from './BaseItem.vue'

export type DebtData = {
  total: number
  returned: number
  left: number
  progress: number
}

export type DebtProps = { item: ItemRecord & { frontmatter: DebtFrontmatter }, compact?: boolean, loading?: boolean }
</script>

<script setup lang="ts">
const { item, compact, loading } = defineProps<DebtProps>()

const emits = defineEmits<BaseItemEmits>()

const debtData = computed<DebtData>(() => {
  const result = item.frontmatter.transactions?.reduce((acc, tx) => {
    if (tx.amount > 0) {
      acc.total += tx.amount
    }
    else {
      acc.returned += Math.abs(tx.amount)
    }
    return acc
  }, { total: 0, returned: 0 }) ?? { total: 0, returned: 0 }

  const left = result.total - result.returned
  const progress = result.total > 0 ? Math.round((clamp(result.returned, 0, result.total) / result.total) * 100) : 0

  return {
    total: result.total,
    returned: result.returned,
    left,
    progress,
  }
})
</script>

<template>
  <DebtCompact
    v-if="compact"
    :item="item"
    :debt-data="debtData"
    @update-frontmatter="(payload) => emits('updateFrontmatter', payload)"
  />
  <DebtDetailed
    v-else
    :item="item"
    :debt-data="debtData"
    :loading="loading"
    @update-frontmatter="(payload) => emits('updateFrontmatter', payload)"
  />
</template>
