<script lang="ts" setup>
import { useActivitiesUpdateItemMutation } from '~/composables/queries'
import type { DebtData, DebtProps, FormState } from '../ItemDebt.vue'
import { computed, h, parseDate, ref, resolveComponent } from '#imports'

import type { TableColumn } from '@nuxt/ui'

const { item, debtData } = defineProps<DebtProps & { debtData: DebtData }>()

const UDropdownMenu = resolveComponent('UDropdownMenu')
const UButton = resolveComponent('UButton')

const _emits = defineEmits<{
  done: [id: string]
}>()

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
// const isEditing = ref(false)
type Transaction = {
  id: string
  date: string
  amount: number
  comment: string | null | undefined
}
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
      return new Date(row.getValue('date')).toLocaleDateString('ru-RU', {
        year: 'numeric',
        month: '2-digit',
        day: '2-digit',
      })
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

const { mutateAsync } = useActivitiesUpdateItemMutation()

async function handleSuccess(payload: FormState) {
  const frontmatter = item.frontmatter
  frontmatter.transactions.push({
    created: payload.date ?? new Date().toISOString().split('.')[0]!,
    amount: payload.amount,
    comment: payload.comment,
  })

  await mutateAsync(item)
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
  await mutateAsync(item)
  expandCb()
}
</script>

<template>
  <div>
    <div
      class="flex items-center mb-1"
    >
      <UIcon
        name="i-lucide-credit-card"
        class="mr-2 text-primary-500"
      />
      <h1 class="text-2xl font-bold">
        {{ item.frontmatter.summary }}
      </h1>
    </div>

    <div>
      <h5>Всего: {{ formatCurrency(debtData.total) }}</h5>
      <h5>Возвращено: {{ formatCurrency(debtData.returned) }}</h5>
      <h5>Осталось: {{ formatCurrency(debtData.left) }}</h5>
      <DebtAddEditForm @success="handleSuccess" />
      <UTable
        v-model:sorting="sorting"
        :data="transactions"
        :columns="columns"
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
  </div>
</template>
