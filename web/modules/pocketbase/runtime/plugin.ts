import type { BaseClient } from '../types/types'
import { useMockClient } from './utils/mock'
import { createPocketBaseClient } from './utils/pb'
import { defineNuxtPlugin, useRuntimeConfig } from '#app'

export default defineNuxtPlugin((_nuxtApp) => {
  const config = useRuntimeConfig()
  const pocketbaseConfig = config.public.pocketbase

  let client: BaseClient

  switch (pocketbaseConfig.type) {
    case 'pb':
      client = createPocketBaseClient(config.public.apiBase)
      break
    case 'mock':
      client = useMockClient()
      break
    default:
      console.error('Invalid pocketbase type, using pb client')
      client = createPocketBaseClient(config.public.apiBase)
      break
  }

  return {
    provide: {
      client,
    },
  }
})
