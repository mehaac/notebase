<script setup lang="ts">
import { computed } from 'vue'
import { definePageMeta, useActivitiesStore, useRoute } from '#imports'
import { LazyBaseItem } from '#components'

definePageMeta({
  middleware: ['auth'],
})

const route = useRoute()
const activitiesStore = useActivitiesStore()
const item = computed(() => {
  if (activitiesStore.items.length === 0) {
    // TODO: find out how to handle async here
    activitiesStore.load()
  }
  return activitiesStore.items.find(item => item.id === route.params.id)
})
</script>

<template>
  <div>
    <ULink to="/">back</ULink>

    <template v-if="item">
      <h1 class="text-2xl font-bold mb-5">
        {{ item.title }}
      </h1>
      <LazyBaseItem
        :item="item"
        :is-list="false"
      />
    </template>
    <template v-else>
      <h1>Item not found</h1>
    </template>
  </div>
</template>
