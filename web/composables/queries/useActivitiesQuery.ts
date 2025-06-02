import {
  defineQuery,
  useClient,
  useFiltersStore,
  useMutation,
  useQuery,
  ref,
  useQueryCache,
} from '#imports'
import type { ItemRecord } from '#pocketbase-imports'
import type { ListResult } from 'pocketbase'

export const useActivitiesListQuery = defineQuery(() => {
  const client = useClient()
  const filtersStore = useFiltersStore()
  const page = ref(1)
  const size = ref(20)

  const { state, error } = useQuery<ListResult<ItemRecord>>({
    key: () => ['activities', { filters: filtersStore.builtQuery, page: page.value, size: size.value }],
    query: async () => {
      return await client.getList(page.value, size.value, filtersStore.builtQuery)
    },
    staleTime: 1000 * 60 * 5, // 5 minutes
    gcTime: 1000 * 60 * 10, // 10 minutes
    enabled: () => filtersStore.enabled,
  })

  return { state, page, size, error }
})

export const useActivitiesItemQuery = defineQuery(() => {
  const client = useClient()
  const id = ref<string>()
  const query = useQuery<ItemRecord>({
    key: () => ['activity', { id: id.value }],
    query: async () => (
      await client.getItem(id.value!)
    ),
    staleTime: 1000 * 60 * 5, // 5 minutes
    gcTime: 1000 * 60 * 10, // 10 minutes
    enabled: () => !!id.value,
  })
  return { ...query, id }
})

export const useActivitiesToggleItemMutation = (
  opts: { onSuccess?: (item: ItemRecord) => Promise<void> | void } = {},
) => {
  const queryCache = useQueryCache()
  const pb = useClient()

  const { state, mutate, mutateAsync, asyncStatus } = useMutation({
    mutation: async (item: ItemRecord) => {
      const frontmatter = item.frontmatter || {}
      if (frontmatter?.completed) {
        frontmatter.completed = ''
      }
      else {
        frontmatter.completed = new Date().toISOString().split('.')[0]!
      }

      await pb.updateFrontmatter(item.id, frontmatter)
      return item
    },
    async onSuccess(_data, vars) {
      await opts.onSuccess?.(vars)
    },
    async onSettled() {
      await queryCache.invalidateQueries({ key: ['activities'], exact: false })
    },
  })

  return { state, mutate, mutateAsync, asyncStatus }
}

export const useActivitiesAddItemMutation = (opts: { onSuccess?: (item: ItemRecord) => Promise<void> | void }) => {
  const _pb = useClient()

  const { state, mutate } = useMutation({
    key: () => ['activities', 'addItem'], // optional
    mutation: async (item: ItemRecord) => new Promise((resolve, reject) =>
      setTimeout(() => Math.random() > 0.5 ? resolve(item) : reject(new Error('Error')), 500),
    ),
    async onSuccess(_data, vars) {
      await opts.onSuccess?.(vars)
    },
  })

  return { state, mutate }
}

export const useActivitiesUpdateItemMutation = (
  opts: { onSuccess?: (item: ItemRecord) => Promise<void> | void } = {},
) => {
  const client = useClient()
  const queryCache = useQueryCache()

  const mutation = useMutation({
    mutation: async (item: ItemRecord) => client.updateFrontmatter(item.id, item.frontmatter!),
    async onSuccess(_data, vars) {
      await queryCache.invalidateQueries({ key: ['activity', { id: vars.id }], exact: true })
      await opts.onSuccess?.(vars)
    },
    async onSettled() {
      await queryCache.invalidateQueries({ key: ['activities'], exact: false })
    },
  })

  return mutation
}
