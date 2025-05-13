import PocketBase from 'pocketbase'
import type { BaseClient } from '../../types/types'
import { frontmatterSchema, transformItem } from '../../types/schema'

export function createPocketBaseClient(url: string): BaseClient {
  const pb = new PocketBase(url)

  const getItem = async (id: string) => {
    const item = await pb.collection('files').getOne(id)
    return transformItem(item)
  }

  const toggleItem = async (id: string) => {
    const item = await getItem(id)
    const frontmatter = item.frontmatter
    if (frontmatter.completed) {
      frontmatter.completed = ''
    }
    else {
      frontmatter.completed = new Date().toISOString()
    }
    const res = await pb.collection('files').update(item.id, {
      frontmatter,
    })
    return transformItem(res)
  }

  const addDebtTransaction = async (id: string, amount: number, comment: string) => {
    const item = await getItem(id)

    const frontmatter = frontmatterSchema.parse(item.frontmatter ? item.frontmatter : {})

    const transaction = {
      amount,
      created: new Date().toISOString(),
      comment,
    }
    frontmatter.transactions?.push(transaction)
    const res = await pb.collection('files').update(item.id, {
      frontmatter,
    })
    return transformItem(res)
  }
  const isAuthenticated = async () => {
    return pb.authStore.isValid
  }
  const clearAuth = async () => {
    pb.authStore.clear()
  }

  const authenticatedUser = async ({ email, password }: { email: string, password: string }) => {
    return await pb.collection('_superusers')
      .authWithPassword(email, password)
  }

  const getList = async (page: number, pageSize: number, filter: string) => {
    const { items, ...rest } = await pb.collection('files').getList(page, pageSize, {
      filter,
    })
    return {
      ...rest,
      items: items.map(transformItem),
    }
  }

  return {
    getItem,
    toggleItem,
    addDebtTransaction,
    isAuthenticated,
    clearAuth,
    authenticatedUser,
    getList,
  }
}
