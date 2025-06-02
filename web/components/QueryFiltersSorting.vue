<script lang="ts" setup>
import { useSortable } from '@vueuse/integrations/useSortable'
import { useTemplateRef } from 'vue'
import { useFiltersStore } from '~/stores/filters'

const filtersStore = useFiltersStore()
const sortableRef = useTemplateRef<HTMLElement>('sortableRef')
useSortable(sortableRef, filtersStore.localFilters, {
  animation: 150,
})
</script>

<template>
  <div class="flex flex-col gap-2">
    <h3 class="text-sm font-medium">
      Sorting
    </h3>
    <div
      ref="sortableRef"
      class="flex flex-wrap gap-2"
    >
      <UBadge
        v-for="filter in filtersStore.localFilters"
        :key="filter.id"
        color="neutral"
        variant="soft"
        class="cursor-grab select-none"
      >
        {{ filter.label }}
      </UBadge>
    </div>
  </div>
</template>
