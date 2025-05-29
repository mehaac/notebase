<script lang="ts">
import { z } from 'zod/v4-mini'
import { computed } from 'vue'
import { clamp } from '@vueuse/core'
import type { ItemRecord, DebtFrontmatter } from '#pocketbase-imports'
import type { BaseItemEmits } from './BaseItem.vue'

export const formSchema = z.object({
  date: z.optional(z.iso.datetime({ local: true })),
  amount: z.number({ error: 'amount is required' }),
  comment: z.optional(z.string().check(
    z.maxLength(256, { error: 'comment is too long' }),
  )),
})

export type FormState = z.infer<typeof formSchema>
export type DebtData = {
  total: number
  returned: number
  left: number
  progress: number
}

export type DebtProps = { item: ItemRecord & { frontmatter: DebtFrontmatter }, compact?: boolean }
</script>

<script setup lang="ts">
const { item, compact } = defineProps<DebtProps>()

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
    @toggle-completed="emits('updateFrontmatter', item)"
  />
  <DebtDetailed
    v-else
    :item="item"
    :debt-data="debtData"
    @toggle-completed="emits('updateFrontmatter', item)"
  />
</template>
