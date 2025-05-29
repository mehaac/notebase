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

    <ItemsList :items="state.data?.items ?? []" />
  </div>
</template>
