<script lang="ts" setup>
import { ref, watch } from 'vue'
import { useFiltersStore } from '~/stores/filters'

const noFilter = {
  label: 'All',
}

const filtersStore = useFiltersStore()
const selectedFilterIndx = ref(0)
const isLoading = ref(false)
watch(selectedFilterIndx, (newVal) => {
  newVal = typeof newVal === 'string' ? parseInt(newVal) : newVal
  isLoading.value = true
  if (newVal === 0) {
    filtersStore.clearFilters()
  }
  else {
    filtersStore.applyFilter(filtersStore.filteredQueryFilters[newVal - 1]!.id)
  }
})
</script>

<template>
  <div class="flex flex-col gap-2">
    <UInput
      v-model="filtersStore.searchFiltersLabel"
      placeholder="Search filters"
      icon="i-lucide-search"
      type="text"
      size="lg"
      class="w-full"
    />
    <div class="scrollable w-full overflow-x-auto py-2">
      <UTabs
        v-model="selectedFilterIndx"
        :items="[noFilter, ...filtersStore.filteredQueryFilters]"
        variant="link"
        class="w-max"
        activation-mode="manual"
      />
    </div>
  </div>
</template>

<style scoped>
.scrollable {
  scrollbar-width: none;
}

.scrollable::-webkit-scrollbar {
  display: none;
}
</style>
