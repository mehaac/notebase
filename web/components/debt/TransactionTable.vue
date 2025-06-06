<script lang="ts" setup>
import type { TableColumn } from '@nuxt/ui'
import type { FormState, Transaction } from './Detailed.vue'
import { formatDateShort, computed, h, parseDate, ref, resolveComponent } from '#imports'
import type { BaseItemEmits } from '../BaseItem.vue'
import type { ItemRecord } from '#pocketbase-imports'

const emits = defineEmits<BaseItemEmits>()

const { item, loading } = defineProps<{
  item: ItemRecord
  loading?: boolean
}>()

const UDropdownMenu = resolveComponent('UDropdownMenu')
const UButton = resolveComponent('UButton')

const transactions = computed(() => {
  return item.frontmatter?.transactions?.map((tx) => {
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
        currency: item.frontmatter?.currency ?? 'RUB',
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

async function handleEdit(originalCreated: string, payload: FormState, expandCb: () => void) {
  const frontmatter = item.frontmatter

  const parsedOriginalCreated = parseDate(originalCreated)
  const transaction = frontmatter?.transactions?.find(tx =>
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
  <UTable
    v-model:sorting="sorting"
    :data="transactions"
    :columns="columns"
    :loading="loading"
    sticky
    class="h-96"
  >
    <template #expanded="{ row }">
      <LazyDebtAddEditForm
        table
        type="edit"
        :defaults="{
          date: row.original.date,
          amount: row.original.amount,
          comment: row.original.comment ?? '',
        }"
        @success="(data) => handleEdit(row.original.date, data, row.toggleExpanded)"
      >
        <template #submit>
          Update
        </template>
      </LazyDebtAddEditForm>
    </template>
  </UTable>
</template>
