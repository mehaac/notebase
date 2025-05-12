import { useNuxtApp } from '#app'

export function useClient() {
  const { $client } = useNuxtApp()
  return $client
}
