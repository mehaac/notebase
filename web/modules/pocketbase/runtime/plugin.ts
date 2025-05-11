import { defineNuxtPlugin, useRuntimeConfig } from '#app'
import PocketBase from 'pocketbase'

interface PocketbaseConfig {
  url?: string
  type?: string
}

export default defineNuxtPlugin((_nuxtApp) => {
  const config = useRuntimeConfig()
  
  // Get URL from config with type safety
  const pbConfig = config.pocketbase as PocketbaseConfig | undefined
  const pbPublicConfig = config.public.pocketbase as PocketbaseConfig | undefined
  const url = pbConfig?.url || pbPublicConfig?.url || (config.public as any).apiBase
  
  if (!url) {
    console.error('PocketBase: No URL provided in configuration')
    return
  }
  
  // Create PocketBase client
  const pb = new PocketBase(url)
  
  // Create custom helpers
  const getItem = (id: string) => {
    return pb.collection('files').getOne(id)
  }

  const toggleItem = async (id: string) => {
    const item = await getItem(id)
    if (item.frontmatter.completed) {
      item.frontmatter.completed = ''
    } else {
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
  
  // Create client with helpers
  const pbClient = {
    client: pb,
    getItem,
    toggleItem,
    addDebtTransaction,
  }
  
  return {
    provide: {
      pb,
      pbClient
    }
  }
}) 
