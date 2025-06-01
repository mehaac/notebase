<script setup lang="ts">
import { UCollapsible } from '#components'
import { definePageMeta, useNotebaseConfig } from '#imports'
import { useActivitiesListQuery } from '~/composables/queries/'

definePageMeta({
  middleware: ['auth'],
})
const notebaseConfig = useNotebaseConfig()
const { state } = useActivitiesListQuery()
</script>

<template>
  <div class="flex flex-col">
    <QueryFiltersTabs />
    <UCollapsible
      class="py-2"
      :open="notebaseConfig.config.value.showFilters"
    >
      <template #content>
        <QueryFilterForm />
      </template>
    </UCollapsible>
    <div v-if="notebaseConfig.config.value.showExtra">
      extra
    </div>
    <div v-if="state.status !== 'success'">
      <!-- TODO: this can be a skeleton loader -->
      <UProgress indeterminate />
    </div>
    <div class="py-2">
      <ItemsList :items="state.data?.items ?? []" />
    </div>
  </div>
</template>
