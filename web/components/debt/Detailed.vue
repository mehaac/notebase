<script lang="ts">
import { z } from 'zod/v4-mini'
import { formatDateShort, computed, h, parseDate, ref, resolveComponent } from '#imports'

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

import type { TableColumn } from '@nuxt/ui'
import type { BaseItemEmits } from '../BaseItem.vue'

const { item, debtData, loading } = defineProps<DebtProps & { debtData: DebtData, loading?: boolean }>()

const UDropdownMenu = resolveComponent('UDropdownMenu')
const UButton = resolveComponent('UButton')

const emits = defineEmits<BaseItemEmits>()

const formatCurrency = (amount: number) => {
  return new Intl.NumberFormat('ru-RU', {
    style: 'currency',
    currency: item.frontmatter.currency,
  }).format(amount)
}

const transactions = computed(() => {
  return item.frontmatter.transactions?.map((tx) => {
    return {
      id: `${tx.created}-${tx.amount}-${tx.comment ?? 'None'}`,
      amount: tx.amount,
      comment: tx.comment,
      date: tx.created,
    }
  }) ?? []
})

const columns: TableColumn<Transaction>[] = [
  {
    accessorKey: 'date',
    header: ({ column }) => {
      // TODO: fix sorting
      const isSorted = column.getIsSorted()

      return h(UButton, {
        color: 'neutral',
        variant: 'ghost',
        label: 'Date',
        icon: isSorted
          ? isSorted === 'asc'
            ? 'i-lucide-arrow-up-narrow-wide'
            : 'i-lucide-arrow-down-wide-narrow'
          : 'i-lucide-arrow-up-down',
        class: '-mx-2.5',
        onClick: () => column.toggleSorting(column.getIsSorted() === 'asc'),
      })
    },
    cell: ({ row }) => {
      return formatDateShort(new Date(row.getValue('date')))
    },
  },
  {
    accessorKey: 'amount',
    header: 'Amount',
    cell: ({ row }) => {
      return new Intl.NumberFormat('ru-RU', {
        style: 'currency',
        currency: item.frontmatter.currency,
      }).format(row.getValue('amount'))
    },
  },
  {
    accessorKey: 'comment',
    header: 'Comment',
    cell: ({ row }) => {
      return row.getValue('comment') ?? '-'
    },
  },
  {
    id: 'actions',
    cell: ({ row }) => {
      return h(
        'div',
        { class: 'text-right' },
        h(
          UDropdownMenu,
          {
            'content': {
              align: 'end',
            },
            'items': [
              {
                label: 'Edit',
                icon: 'i-lucide-edit',
                onClick: () => row.toggleExpanded(),
              },
            ],
            'aria-label': 'Actions dropdown',
          },
          () =>
            h(UButton, {
              'icon': 'i-lucide-ellipsis-vertical',
              'color': 'neutral',
              'variant': 'ghost',
              'class': 'ml-auto',
              'aria-label': 'Actions dropdown',
            }),
        ),
      )
    },
  },
]

const sorting = ref([
  {
    id: 'date',
    desc: false,
  },
])

async function handleSuccess(payload: FormState) {
  const frontmatter = item.frontmatter
  frontmatter.transactions.push({
    created: payload.date ?? new Date().toISOString().split('.')[0]!,
    amount: payload.amount,
    comment: payload.comment,
  })
  emits('updateFrontmatter', item)
}

async function handleEdit(originalCreated: string, payload: FormState, expandCb: () => void) {
  const frontmatter = item.frontmatter

  const parsedOriginalCreated = parseDate(originalCreated)
  const transaction = frontmatter.transactions.find(tx =>
    parseDate(tx.created).compare(parsedOriginalCreated) === 0,
  )
  console.log(transaction)
  if (!transaction) return
  transaction.amount = payload.amount ?? transaction.amount
  transaction.comment = payload.comment ?? transaction.comment
  transaction.created = payload.date ?? transaction.created
  emits('updateFrontmatter', item)
  expandCb()
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
      <UTable
        v-model:sorting="sorting"
        :data="transactions"
        :columns="columns"
        :loading="loading"
      >
        <template #expanded="{ row }">
          <LazyDebtAddEditForm
            type="edit"
            :defaults="{
              date: row.original.date,
              amount: row.original.amount,
              comment: row.original.comment ?? '',
            }"
            @success="(data) => handleEdit(row.original.date, data, row.toggleExpanded)"
          >
            <template #submit>
              update
            </template>
          </LazyDebtAddEditForm>
        </template>
      </UTable>
    </div>
  </ItemCard>
</template>
