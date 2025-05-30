<script lang="ts">
import { z } from 'zod/v4-mini'

export const formSchema = z.object({
  date: z.optional(z.iso.datetime({ local: true })),
  amount: z.number({ error: 'amount is required' }),
  comment: z.optional(z.string().check(
    z.maxLength(256, { error: 'comment is too long' }),
  )),
})

export type FormState = z.infer<typeof formSchema>
export type Transaction = {
  id: string
  date: string
  amount: number
  comment: string | null | undefined
}
</script>

<script lang="ts" setup>
import type { DebtData, DebtProps } from '../ItemDebt.vue'

import type { BaseItemEmits } from '../BaseItem.vue'

const { item, debtData, loading } = defineProps<DebtProps & { debtData: DebtData, loading?: boolean }>()

const emits = defineEmits<BaseItemEmits>()

const formatCurrency = (amount: number) => {
  return new Intl.NumberFormat('ru-RU', {
    style: 'currency',
    currency: item.frontmatter.currency,
  }).format(amount)
}

async function handleSuccess(payload: FormState) {
  const frontmatter = item.frontmatter
  frontmatter.transactions.push({
    created: payload.date ?? new Date().toISOString().split('.')[0]!,
    amount: payload.amount,
    comment: payload.comment,
  })
  emits('updateFrontmatter', item)
}
</script>

<template>
  <ItemCard
    :item="item"
    :icon="'i-lucide-credit-card'"
    :loading="loading"
    @toggle-completed="(payload) => emits('updateFrontmatter', payload)"
  >
    <div>
      <h5>Всего: {{ formatCurrency(debtData.total) }}</h5>
      <h5>Возвращено: {{ formatCurrency(debtData.returned) }}</h5>
      <h5>Осталось: {{ formatCurrency(debtData.left) }}</h5>
      <DebtAddEditForm @success="handleSuccess" />
    </div>

    <DebtTransactionTable
      :item="item"
      :loading="loading"
      @update-frontmatter="(payload) => emits('updateFrontmatter', payload)"
    />
  </ItemCard>
</template>
