<script setup lang="ts">
import { definePageMeta } from '#imports'
import { useActivitiesListQuery } from '~/composables/queries/'

definePageMeta({
  middleware: ['auth'],
})

const { state } = useActivitiesListQuery()
</script>

<template>
  <div class="flex flex-col gap-4">
    <AppFilters />
    <div v-if="state.status !== 'success'">
      <!-- TODO: this can be a skeleton loader -->
      <UProgress indeterminate />
    </div>

    <ul
      v-if="state.status === 'success'"
      class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4"
    >
      <LazyListItem
        v-for="item in state.data?.items"
        :key="item.id"
        :item="item"
      />
    </ul>
  </div>
</template>
