import PocketBase from 'pocketbase'
import type { BaseClient } from '../../types/types'

export function createPocketBaseClient(url: string): BaseClient {
  const pb = new PocketBase(url)

  const getItem = (id: string) => {
    return pb.collection('files').getOne(id)
  }

  const toggleItem = async (id: string) => {
    const item = await getItem(id)
    if (item.frontmatter.completed) {
      item.frontmatter.completed = ''
    }
    else {
      item.frontmatter.completed = new Date().toISOString()
    }
    return pb.collection('files').update(item.id, {
      frontmatter: item.frontmatter,
    })
  }

  const addDebtTransaction = async (id: string, amount: number, comment: string) => {
    const item = await getItem(id)
    const transaction = {
      amount,
      date: new Date().toISOString(),
      comment,
    }
    item.frontmatter.transactions.push(transaction)
    return pb.collection('files').update(item.id, {
      frontmatter: item.frontmatter,
    })
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
    return pb.collection('files').getList(page, pageSize, {
      filter,
    })
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
