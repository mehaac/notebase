import { defineStore } from 'pinia'
import { useLocalStorage, useStorage, watchDebounced } from '@vueuse/core'
import { ref } from 'vue'

interface Filter {
  id: string
  label: string
  query: string
  queryType: string
  pathFilter: string
  typeFilter: string
  pathFilterEnabled: boolean
  typeFilterEnabled: boolean
}

const useFiltersLocalStorage = () => useLocalStorage<Filter[]>('notebase-filters', () => [])

export const useFiltersStore = defineStore('filters', () => {
  const localFilters = useFiltersLocalStorage()

  const query = useStorage('query', '', localStorage)
  const queryType = useStorage('query-type', 'FTS', localStorage)
  const pathFilter = useStorage('path-filter', '', localStorage)
  const typeFilter = useStorage('type-filter', '', localStorage)
  const pathFilterEnabled = useStorage('path-filter-enabled', false, localStorage)
  const typeFilterEnabled = useStorage('type-filter-enabled', false, localStorage)

  const appliedFilterId = ref<string>()

  const builtQuery = ref('')

  const createdFilterLabel = ref<string>()

  const enabled = ref(false)

  function buildQuery() {
    const filterParts: string[] = []
    if (typeFilterEnabled.value && typeFilter.value.length > 0) {
      filterParts.push(`frontmatter.type = '${typeFilter.value}'`)
    }
    if (pathFilterEnabled.value && pathFilter.value.length > 0) {
      filterParts.push(`path ~ '${pathFilter.value}'`)
    }
    if (query.value.length > 0) {
      if (queryType.value === 'FTS') {
        filterParts.push(`(content~'${query.value}'||frontmatter.summary~'${query.value}'||frontmatter.title~'${query.value}')`)
      }
      else if (queryType.value === 'QL') {
        filterParts.push(query.value)
      }
    }
    return filterParts.join(' && ')
  }

  function saveFilter() {
    const id = crypto.randomUUID()
    if (!createdFilterLabel.value) {
      return
    }
    const filter: Filter = {
      id,
      label: createdFilterLabel.value,
      query: query.value,
      queryType: queryType.value,
      pathFilter: pathFilter.value,
      typeFilter: typeFilter.value,
      pathFilterEnabled: pathFilterEnabled.value,
      typeFilterEnabled: typeFilterEnabled.value,
    }
    localFilters.value.push(filter)
    return filter
  }

  function applyFilter(id: string) {
    const filter = localFilters.value.find(f => f.id === id)
    if (!filter) return false
    if (appliedFilterId.value === id) {
      clearFilters()
      return true
    }

    appliedFilterId.value = id
    createdFilterLabel.value = filter.label
    query.value = filter.query
    queryType.value = filter.queryType
    pathFilter.value = filter.pathFilter
    typeFilter.value = filter.typeFilter
    pathFilterEnabled.value = filter.pathFilterEnabled
    typeFilterEnabled.value = filter.typeFilterEnabled

    return true
  }

  function clearFilters() {
    createdFilterLabel.value = ''
    query.value = ''
    pathFilter.value = ''
    typeFilter.value = ''
    pathFilterEnabled.value = false
    typeFilterEnabled.value = false
    appliedFilterId.value = undefined
  }

  function deleteFilter() {
    const index = localFilters.value.findIndex(f => f.id === appliedFilterId.value)
    if (index !== -1) {
      localFilters.value.splice(index, 1)
      clearFilters()
    }
  }

  function searchFilters(labelQuery: string, limit: number = 5): Filter[] {
    if (!labelQuery.trim()) return []
    const q = labelQuery.toLowerCase()
    return localFilters.value
      .filter(f => f.label.toLowerCase().includes(q))
      .slice(0, limit)
  }

  watchDebounced(
    [query, pathFilter, typeFilter, pathFilterEnabled, typeFilterEnabled],
    () => {
      if (!enabled.value) {
        enabled.value = true
      }
      builtQuery.value = buildQuery()
    },
    { debounce: 300, immediate: true },
  )

  return {
    query,
    queryType,
    pathFilter,
    typeFilter,
    pathFilterEnabled,
    typeFilterEnabled,
    builtQuery,
    enabled,
    createdFilterLabel,
    appliedFilterId,
    buildQuery,
    saveFilter,
    applyFilter,
    clearFilters,
    deleteFilter,
    searchFilters,

    localFilters,
  }
})
