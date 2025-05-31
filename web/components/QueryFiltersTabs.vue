<script lang="ts" setup>
import { ref, watch } from 'vue'
import { useFiltersStore, useSelectedFilters } from '~/stores/filters'

// const noFilter = {
//   label: 'All',
//   class: 'sticky left-0',
// }

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

const selectedFilters = useSelectedFilters()

const handleChange = (id: string) => {
  if (selectedFilters.value.has(id)) {
    selectedFilters.value.delete(id)
  }
  else {
    selectedFilters.value.add(id)
  }
}
</script>

<template>
  <div class="flex flex-col gap-2">
    <div class="flex items-center">
      <div>
        <UButton
          color="neutral"
          :variant="
            !filtersStore.appliedFilterId || !filtersStore.builtQuery
              ? 'soft'
              : 'ghost'
          "
          size="lg"
          label="All"
          class="w-full"
          @click="filtersStore.clearFilters()"
        />
      </div>
      <div class="scrollable w-full overflow-x-auto py-2 relative">
        <div class="flex gap-2">
          <UButtonGroup
            v-for="filter in filtersStore.filteredQueryFilters"
            :key="filter.id"
            class="flex items-center gap-0.5"
          >
            <UButton
              :label="filter.label"
              color="neutral"
              :variant="
                filtersStore.appliedFilterId === filter.id ? 'soft' : 'ghost'
              "
              size="lg"
              class="w-full"
              @click="filtersStore.applyFilter(filter.id)"
            >
              {{ filter.label }}
            </UButton>
            <UCheckbox
              color="neutral"
              size="lg"
              class="w-full"
              @change="handleChange(filter.id)"
            />
          </UButtonGroup>
        </div>
      </div>
      <div>
        <UButton
          color="neutral"
          variant="ghost"
          size="lg"
          label="Add"
          icon="i-lucide-plus"
          class="w-full"
        />
      </div>
      <!-- <UTabs
        v-model="selectedFilterIndx"
        :items="[noFilter, ...filtersStore.filteredQueryFilters]"
        variant="link"
        class="w-max"
        activation-mode="manual"
        :ui="{ trigger: 'grow' }"
      >
        <template
          v-for="filter in filtersStore.filteredQueryFilters"
          :key="filter.id"
          #[`${filter.label}-content`]
        >
          <UIcon name="i-lucide-trash" />
        </template>
      </UTabs> -->
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
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.5s ease;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
