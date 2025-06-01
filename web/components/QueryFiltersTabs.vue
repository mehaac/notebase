<script lang="ts" setup>
import { useNotebaseConfig } from '#imports'
import type { DropdownMenuItem } from '@nuxt/ui'
import { computed, ref, watch } from 'vue'
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

const items = computed<DropdownMenuItem[]>(() => {
  return [
    {
      slot: 'edit',
    },
    {
      slot: 'remove',
    },
  ]
})

function toggleFilters() {
  notebaseConfig.setShowFilters(!notebaseConfig.config.value.showFilters)
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
            label="All"
            block

            @click="filtersStore.clearFilters()"
          />
          <UButton
            color="neutral"
            variant="outline"
            class="h-10 w-12"
            block
            :icon="notebaseConfig.config.value.showFilters ? 'i-lucide-minus' : 'i-lucide-plus'"
            @click="toggleFilters"
          />
        </div>
        <div class="flex gap-2 pl-2">
          <template v-if="filtersStore.filteredQueryFilters.length > 0">
            <UButtonGroup
              v-for="filter in filtersStore.filteredQueryFilters"
              :key="filter.id"
              class="flex items-center gap-1 ring-1 transition-all duration-200 ring-(--ui-border) rounded-lg"
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
              <UDropdownMenu
                :items="items"
                arrow
                external-icon="i-lucide-chevron-down"
              >
                <UButton
                  color="neutral"
                  variant="ghost"
                  icon="i-lucide-chevron-down"
                />
                <template
                  v-if="!notebaseConfig.config.value.showFilters"
                  #edit
                >
                  <UButton
                    color="primary"
                    variant="ghost"
                    size="xs"
                    class="w-full"
                    label="Edit"
                    icon="i-lucide-pen"
                    @click="notebaseConfig.setShowFilters(true)"
                  />
                </template>
                <template #remove>
                  <UButton
                    color="error"
                    variant="ghost"
                    size="xs"
                    class="w-full"
                    label="Remove"
                    icon="i-lucide-trash"
                    @click="filtersStore.deleteFilter(filter.id)"
                  />
                </template>
              </UDropdownMenu>
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
