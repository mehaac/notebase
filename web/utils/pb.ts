import PocketBase from 'pocketbase'
import { useRuntimeConfig } from '#app'

const config = useRuntimeConfig()

export const pb = new PocketBase(config.public.apiBase)
