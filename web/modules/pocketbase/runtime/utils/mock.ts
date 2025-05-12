import type { RecordModel } from 'pocketbase'
import { useLocalStorage } from '@vueuse/core'
import type { BaseClient } from '../../types/types'
import tasks from '~/assets/mock/tasks.json'
import debts from '~/assets/mock/debts.json'
import tracks from '~/assets/mock/tracks.json'

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
    const items = [...tasks, ...debts, ...tracks].map((item) => {
      const defaults = createDefaultReturn()
      const { content: _, ...rest } = item
      return {
        ...defaults,
        content: item.content,
        path: `${item.type}/${defaults.id}.md`,
        frontmatter: rest,
      }
    })
    localItems.value = items
  }

  return {
    getItem: async (id: string) => {
      return localItems.value.find(item => item.id === id) || defaultReturn
    },
    toggleItem: async (id: string) => {
      return localItems.value.find(item => item.id === id) || defaultReturn
    },
    addDebtTransaction: async (id: string, _amount: number, _comment: string) => {
      return localItems.value.find(item => item.id === id) || defaultReturn
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
