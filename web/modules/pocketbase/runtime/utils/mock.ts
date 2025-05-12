import type { RecordModel } from 'pocketbase'
import { useLocalStorage } from '@vueuse/core'
import type { BaseClient } from '../../types/types'
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

const useLocalItems = () => useLocalStorage<RecordModel[]>('items', [])

export function useMockClient(): BaseClient {
  const defaultReturn = createDefaultReturn()
  const localItems = useLocalItems()
  if (localItems.value.length === 0) {
    const items = [...tasks, ...debts, ...tracks, ...withContent].map((item) => {
      const defaults = createDefaultReturn()
      const { content, ...rest } = item
      const type = (item as { type?: string })?.type ?? 'base'

      if (!('title' in rest) && !('summary' in rest)) {
        (rest as { title?: string }).title = 'Mock Item'
      }
      return {
        ...defaults,
        content: content ?? '',
        path: `${type}/${defaults.id}.md`,
        frontmatter: {
          ...rest,
        },
      }
    })
    localItems.value = items
  }

  return {
    getItem: async (id: string) => {
      return localItems.value.find(item => item.id === id) || defaultReturn
    },
    toggleItem: async (id: string) => {
      const item = localItems.value.find(item => item.id === id)
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
      if (item) {
        item.frontmatter.transactions.push({
          amount,
          comment,
          date: new Date().toISOString(),
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
