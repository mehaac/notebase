<script lang="ts" setup>
import { useNotebaseConfig } from '#imports'
import { ref, watch, nextTick, useTemplateRef } from 'vue'
import { useFiltersStore } from '~/stores/filters'

const filtersStore = useFiltersStore()
const notebaseConfig = useNotebaseConfig()
const selectedFilterIndx = ref(0)
const isLoading = ref(false)
const sortableContainer = useTemplateRef<HTMLElement>('sortableContainer')
watch(selectedFilterIndx, (newVal) => {
  newVal = typeof newVal === 'string' ? parseInt(newVal) : newVal
  isLoading.value = true
  if (newVal === 0) {
    filtersStore.clearFilters()
  }
  else {
    filtersStore.applyFilter(filtersStore.localFilters[newVal - 1]!.id)
  }
})

function handleFilterMenu(filterId: string) {
  const isFilterSelected = filtersStore.appliedFilterId === filterId
  const isMenuOpen = notebaseConfig.config.value.showFilters

  if (isFilterSelected && isMenuOpen) {
    notebaseConfig.setShowFilters(false)
  }
  else {
    filtersStore.applyFilter(filterId)
    notebaseConfig.setShowFilters(true)
  }
}

async function handleAddFilter() {
  if (filtersStore.appliedFilterId) {
    filtersStore.clearFilters()
  }
  filtersStore.saveFilterLabel = `New ${filtersStore.localFilters.length + 1}`
  const newFilter = filtersStore.saveFilter()
  if (!newFilter) {
    return
  }
  filtersStore.applyFilter(newFilter.id)
  notebaseConfig.setShowFilters(true)

  await nextTick()
  if (sortableContainer.value) {
    const filterElement = sortableContainer.value.querySelector(`[data-filter-id="${newFilter.id}"]`)
    if (filterElement) {
      const containerRect = sortableContainer.value.getBoundingClientRect()
      const elementRect = filterElement.getBoundingClientRect()
      const scrollLeft = sortableContainer.value.scrollLeft + (elementRect.left - containerRect.left) - (containerRect.width / 2) + (elementRect.width / 2)

      sortableContainer.value.scrollTo({
        left: scrollLeft,
        behavior: 'smooth',
      })
    }
  }
}
function handleClearFilters() {
  filtersStore.clearFilters()
  notebaseConfig.setShowFilters(false)
}
</script>

<template>
  <div class="flex flex-col">
    <div class="flex items-center relative min-h-12">
      <div class="flex gap-1 items-center bg-(--ui-bg)">
        <UButton
          color="neutral"
          variant="link"
          size="xs"
          icon="i-lucide-arrow-down-up"
          @click="notebaseConfig.setShowTabsSorting(!notebaseConfig.config.value.showTabsSorting)"
        />
        <UButton
          color="neutral"
          :variant="
            !filtersStore.appliedFilterId || !filtersStore.builtQuery
              ? 'soft'
              : 'outline'
          "
          icon="i-lucide-filter-x"
          block
          @click="handleClearFilters"
        />
      </div>
      <div
        ref="sortableContainer"
        class="scrollable flex items-center w-full overflow-x-auto"
      >
        <div class="flex gap-2 pl-2 py-1">
          <template v-if="filtersStore.localFilters.length > 0">
            <UButtonGroup
              v-for="filter in filtersStore.localFilters"
              :key="filter.id"
              :data-filter-id="filter.id"
              class="rounded-md"
              :class="filtersStore.appliedFilterId === filter.id ? 'ring-1 ring-primary' : 'ring-inset ring-(--ui-border)'"
            >
              <UButton
                :label="filter.label"
                color="neutral"
                :variant="
                  filtersStore.appliedFilterId === filter.id ? 'outline' : 'ghost'
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
      <div class="flex items-center gap-0.5">
        <UButton
          color="neutral"
          variant="outline"
          block
          icon="i-lucide-plus"
          @click="handleAddFilter"
        />
      </div>
    </div>
  </div>
</template>

<style scoped>
.scrollable {
  scroll-behavior: smooth;
  scroll-snap-type: x mandatory;
  scrollbar-width: thin;
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
