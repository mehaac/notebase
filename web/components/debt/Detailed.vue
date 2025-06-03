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
    <div class="flex flex-col gap-2">
      <UProgress
        :model-value="debtData.progress"
        status
        class="pb-4"
      />
      <div class="flex flex-col space-y-3 max-w-sm">
        <div class="flex items-center justify-between py-2 px-1 border-b border-(--ui-border)">
          <span class="text-sm font-medium">Total</span>
          <span class="text-base font-semibold">
            {{ formatCurrency(debtData.total) }}
          </span>
        </div>
        <div class="flex items-center justify-between py-2 px-1 border-b border-(--ui-border)">
          <span class="text-sm font-medium">Returned</span>
          <span class="text-base font-semibold">
            {{ formatCurrency(debtData.returned) }}
          </span>
        </div>
        <div class="flex items-center justify-between py-2 px-1">
          <span class="text-sm font-medium">Left</span>
          <span class="text-base font-semibold">
            {{ formatCurrency(debtData.left) }}
          </span>
        </div>
      </div>
      <DebtAddEditForm @success="handleSuccess" />
    </div>

    <DebtTransactionTable
      :item="item"
      :loading="loading"
      @update-frontmatter="(payload) => emits('updateFrontmatter', payload)"
    />
  </ItemCard>
</template>
