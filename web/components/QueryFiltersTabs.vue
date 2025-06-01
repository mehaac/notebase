<script lang="ts" setup>
import { useNotebaseConfig } from '#imports'
import { ref, watch } from 'vue'
import { useFiltersStore } from '~/stores/filters'

const filtersStore = useFiltersStore()
const notebaseConfig = useNotebaseConfig()
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

function handleFilterMenu(filterId: string) {
  const isFilterSelected = filtersStore.appliedFilterId === filterId
  const isMenuOpen = notebaseConfig.config.value.showFilters

  if (isFilterSelected && isMenuOpen) {
    // Filter selected and menu open -> close menu
    notebaseConfig.setShowFilters(false)
  }
  else {
    // All other cases -> select filter and open menu
    filtersStore.applyFilter(filterId)
    notebaseConfig.setShowFilters(true)
  }
}

function handleAddFilter() {
  if (filtersStore.appliedFilterId) {
    filtersStore.clearFilters()
  }
  filtersStore.saveFilterLabel = `New ${filtersStore.filteredQueryFilters.length + 1}`
  const newFilter = filtersStore.saveFilter()
  if (!newFilter) {
    return
  }
  filtersStore.applyFilter(newFilter.id)
  notebaseConfig.setShowFilters(true)
}
function handleClearFilters() {
  filtersStore.clearFilters()
  notebaseConfig.setShowFilters(false)
}
</script>

<template>
  <div class="flex flex-col">
    <div class="flex items-center relative min-h-12">
      <div class="scrollable flex items-center w-full overflow-x-auto py-2">
        <div class="flex gap-1 sticky left-0 z-10 items-center bg-(--ui-bg)">
          <UButton
            color="neutral"
            :variant="
              !filtersStore.appliedFilterId || !filtersStore.builtQuery
                ? 'soft'
                : 'outline'
            "
            class="h-10 w-12"
            icon="i-lucide-filter-x"
            block
            @click="handleClearFilters"
          />
          <UButton
            color="neutral"
            variant="outline"
            class="h-10 w-12"
            block
            icon="i-lucide-plus"
            @click="handleAddFilter"
          />
        </div>
        <div class="flex gap-2 pl-2">
          <template v-if="filtersStore.filteredQueryFilters.length > 0">
            <UButtonGroup
              v-for="filter in filtersStore.filteredQueryFilters"
              :key="filter.id"
              class="flex items-center gap-1 ring-1 transition-all duration-200 ring-(--ui-border) rounded-none"
              :class="filtersStore.appliedFilterId === filter.id ? 'ring-primary' : ''"
            >
              <UButton
                :label="filter.label"
                color="neutral"
                :variant="
                  filtersStore.appliedFilterId === filter.id ? 'soft' : 'ghost'
                "
                class="w-full truncate"
                @click="filtersStore.applyFilter(filter.id)"
              >
                {{ filter.label }}
              </UButton>
              <UButton
                color="neutral"
                variant="ghost"
                :icon="notebaseConfig.config.value.showFilters && filtersStore.appliedFilterId === filter.id ? 'i-lucide-chevron-up' : 'i-lucide-chevron-down'"
                @click="handleFilterMenu(filter.id)"
              />
            </UButtonGroup>
          </template>
          <template v-else>
            <span class="text-xs text-dimmed">
              No filters
            </span>
          </template>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.scrollable {
  scroll-behavior: smooth;
  scroll-snap-type: x mandatory;
  scroll-snap-align: start;
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.5s ease;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
