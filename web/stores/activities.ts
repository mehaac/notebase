import { defineStore } from 'pinia'
import {
  type Item,
  shallowRef,
  transformItem,
  useFiltersStore,
  useClient,
} from '#imports'
import type { ItemType } from '#pocketbase-imports'

export const useActivitiesStore = defineStore('activities', () => {
  const pb = useClient()

  const items = shallowRef<Item[]>([])
  const itemTypes = shallowRef<Set<ItemType>>(new Set())
  const item = shallowRef<Item | undefined>(undefined)
  const filtersStore = useFiltersStore()

  const load = async () => {
    const resultList = await pb.getList(1, 20, filtersStore.buildQuery())
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    items.value = resultList.items.map((item: any) => {
      itemTypes.value.add(item.frontmatter.type)
      return transformItem(item)
    })
  }

  return { items, item, itemTypes, load }
})
