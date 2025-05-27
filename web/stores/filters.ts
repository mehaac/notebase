import { defineStore } from 'pinia'
import { useStorage, watchDebounced } from '@vueuse/core'
import { ref } from 'vue'

export const useFiltersStore = defineStore('filters', () => {
  const query = useStorage('query', '', localStorage)
  const queryType = useStorage('query-type', 'FTS', localStorage)
  const pathFilter = useStorage('path-filter', '', localStorage)
  const typeFilter = useStorage('type-filter', '', localStorage)
  const pathFilterEnabled = useStorage('path-filter-enabled', false, localStorage)
  const typeFilterEnabled = useStorage('type-filter-enabled', false, localStorage)

  const buildedQuery = ref('')

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

  watchDebounced(
    [
      query,
      pathFilter,
      typeFilter,
      pathFilterEnabled,
      typeFilterEnabled,
    ],
    () => {
      buildedQuery.value = buildQuery()
    },
    { debounce: 300 },
  )

  return {
    query,
    queryType,
    pathFilter,
    typeFilter,
    pathFilterEnabled,
    typeFilterEnabled,
    buildQuery,
    buildedQuery,
  }
})
