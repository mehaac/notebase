<script lang="ts" setup>
import { useNotebaseConfig } from '#imports'
import { ref, watch, nextTick } from 'vue'
import { useFiltersStore } from '~/stores/filters'

const filtersStore = useFiltersStore()
const notebaseConfig = useNotebaseConfig()
const selectedFilterIndx = ref(0)
const isLoading = ref(false)
const sortableContainer = ref<HTMLElement>()
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
      <div
        class="scrollable flex items-center w-full overflow-x-auto"
      >
        <div class="absolute right-0 -top-4 z-20">
          <UButton
            color="neutral"
            variant="link"
            size="xs"
            icon="i-lucide-stretch-horizontal"
            @click="notebaseConfig.setShowTabsSorting(!notebaseConfig.config.value.showTabsSorting)"
          />
        </div>
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
        <div
          class="flex gap-2 pl-2"
        >
          <template v-if="filtersStore.localFilters.length > 0">
            <UButtonGroup
              v-for="filter in filtersStore.localFilters"
              :key="filter.id"
              :data-filter-id="filter.id"
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
