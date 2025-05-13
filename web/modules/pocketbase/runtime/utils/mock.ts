import type { RecordModel } from 'pocketbase'
import { useLocalStorage } from '@vueuse/core'
import type { BaseClient, Item } from '../../types/types'
import { frontmatterSchema, itemTypes } from '../../types/schema'
import tasks from '~/assets/mock/tasks.json'
import debts from '~/assets/mock/debts.json'
import tracks from '~/assets/mock/tracks.json'
import withContent from '~/assets/mock/withContent.json'

const createDefaultReturn = () => {
  const defaultReturn: RecordModel = {
    id: crypto.randomUUID().toString(),
    collectionId: 'pbc_boba',
    collectionName: 'files',
    path: '',
    slug: '',
    updated: '',
    content: '\n',
  }
  return defaultReturn
}

const useLocalItems = () => useLocalStorage<Item[]>('items', [])

export function useMockClient(): BaseClient {
  const defaultReturn = createDefaultReturn()
  const localItems = useLocalItems()
  if (localItems.value.length === 0) {
    const items = [...tasks, ...debts, ...tracks, ...withContent].map((item) => {
      const defaults = createDefaultReturn()
      const { content, ...rest } = item
      const frontmatter = frontmatterSchema.parse(rest)

      const _item: Item = {
        id: defaults.id,
        title: frontmatter.title ?? 'Mock Item',
        content: content ?? '',
        done: Math.random() > 0.5 ? true : false,
        type: frontmatter.type ?? itemTypes.none,
        frontmatter,
      }
      return _item
    })
    localItems.value = items
  }

  return {
    getItem: async (id: string) => {
      const item = localItems.value.find(item => item.id === id)
      if (!item) {
        throw new Error('Item not found')
      }
      return item
    },
    toggleItem: async (id: string) => {
      const item = localItems.value.find(item => item.id === id)
      if (!item) {
        throw new Error('Item not found')
      }
      if (item && item.frontmatter.completed) {
        item.frontmatter.completed = ''
      }
      else if (item) {
        item.frontmatter.completed = new Date().toISOString()
      }
      return item || defaultReturn
    },
    addDebtTransaction: async (id: string, amount: number, comment: string) => {
      const item = localItems.value.find(item => item.id === id)
      if (!item) {
        throw new Error('Item not found')
      }
      if (item.frontmatter.transactions) {
        item.frontmatter.transactions.push({
          amount,
          comment,
          created: new Date().toISOString(),
        })
      }
      return item || defaultReturn
    },
    getList: async (page: number, pageSize: number, _filter: string) => {
      const filteredItems = localItems.value.filter(() => true)
      const start = (page - 1) * pageSize
      const end = start + pageSize
      const paginatedItems = filteredItems.slice(start, end)
      return {
        items: paginatedItems,
        totalItems: filteredItems.length,
        totalPages: Math.ceil(filteredItems.length / pageSize),
        perPage: pageSize,
        page,
      }
    },
    isAuthenticated: async () => true,
    clearAuth: async () => {},
    authenticatedUser: async (_payload: { email: string, password: string }) => {
      return {
        record: {
          collectionId: '1',
          collectionName: 'items',
          id: '1',
          title: 'Mock Item',
          content: 'Mock Content',
        },
        token: '123',
        meta: {},
      }
    },
  }
}
