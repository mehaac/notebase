import { useNuxtApp } from '#app'
import PocketBase from 'pocketbase'

export function usePocketBase() {
  const { $pb } = useNuxtApp()
  return $pb as PocketBase
}

export function usePocketBaseClient() {
  const { $pbClient } = useNuxtApp()
  return $pbClient
} 
