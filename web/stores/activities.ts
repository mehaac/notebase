import { defineStore } from 'pinia'
import {
  shallowRef,
  useFiltersStore,
  useClient,
} from '#imports'
import type { Item, ItemType } from '#pocketbase-imports'

export const useActivitiesStore = defineStore('activities', () => {
  const pb = useClient()

  const items = shallowRef<Item[]>([])
  const itemTypes = shallowRef<Set<ItemType>>(new Set())
  const item = shallowRef<Item | undefined>(undefined)
  const filtersStore = useFiltersStore()

  const load = async () => {
    const resultList = await pb.getList(1, 20, filtersStore.buildQuery())
    items.value = resultList.items
  }

  return { items, item, itemTypes, load }
})
