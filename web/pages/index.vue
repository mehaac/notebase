<script setup lang="ts">
import { definePageMeta, useNotebaseConfig } from '#imports'
import { useActivitiesListQuery } from '~/composables/queries/'

definePageMeta({
  middleware: ['auth'],
})
const notebaseConfig = useNotebaseConfig()
const { state } = useActivitiesListQuery()
</script>

<template>
  <div class="flex flex-col gap-4">
    <QueryFilterForm v-if="notebaseConfig.config.value.showFilters" />
    <QueryFiltersTabs />
    <div v-if="notebaseConfig.config.value.showExtra">
      extra
    </div>
    <div v-if="state.status !== 'success'">
      <!-- TODO: this can be a skeleton loader -->
      <UProgress indeterminate />
    </div>

    <ItemsList :items="state.data?.items ?? []" />
    <QueryFilterSlideover />
  </div>
</template>
