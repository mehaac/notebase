import { defineStore } from 'pinia'
import { shallowRef, watch } from 'vue'

export const useFiltersStore = defineStore('filters', () => {
  const query = shallowRef('')
  const pathFilter = shallowRef('')
  const typeFilter = shallowRef('')
  const pathFilterEnabled = shallowRef(false)
  const typeFilterEnabled = shallowRef(false)

  const buildQuery = () => {
    let filter = query.value
    if (pathFilterEnabled.value && pathFilter.value.length > 0) {
      filter += ` && path ~ '${pathFilter.value}'`
    }
    if (typeFilterEnabled.value && typeFilter.value.length > 0) {
      filter += ` && frontmatter.type = '${typeFilter.value}'`
    }
    console.log(filter)
    return filter
  }
  watch(query, buildQuery)

  return { query, pathFilter, typeFilter, pathFilterEnabled, typeFilterEnabled, buildQuery }
})
