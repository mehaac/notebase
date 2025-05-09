import { defineStore } from 'pinia'
import { useStorage, watchDebounced } from '@vueuse/core'
import { useActivitiesStore } from '#imports'

export const useFiltersStore = defineStore('filters', () => {
  const query = useStorage('query', '', localStorage)
  const pathFilter = useStorage('path-filter', '', localStorage)
  const typeFilter = useStorage('type-filter', '', localStorage)
  const pathFilterEnabled = useStorage('path-filter-enabled', false, localStorage)
  const typeFilterEnabled = useStorage('type-filter-enabled', false, localStorage)

  const activitiesStore = useActivitiesStore()

  const buildQuery = () => {
    const filterParts: string[] = []
    if (typeFilterEnabled.value && typeFilter.value.length > 0) {
      filterParts.push(`frontmatter.type = '${typeFilter.value}'`)
    }
    if (pathFilterEnabled.value && pathFilter.value.length > 0) {
      filterParts.push(`path ~ '${pathFilter.value}'`)
    }
    if (query.value.length > 0) {
      filterParts.push(query.value)
    }
    return filterParts.join(' && ')
  }

  watchDebounced(
    [
      query,
      pathFilter,
      typeFilter,
      pathFilterEnabled,
      typeFilterEnabled,
    ],
    async () => {
      await activitiesStore.load()
    },
    { debounce: 300 },
  )

  return {
    query,
    pathFilter,
    typeFilter,
    pathFilterEnabled,
    typeFilterEnabled,
    buildQuery,
  }
})
