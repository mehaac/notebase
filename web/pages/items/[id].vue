<script setup lang="ts">
import { computed } from 'vue'
import { definePageMeta, useActivitiesStore, useRoute } from '#imports'
import { BaseItemComponent } from '#components'

definePageMeta({
  middleware: ['auth'],
})

const route = useRoute()
const activitiesStore = useActivitiesStore()
const item = computed(() =>
  activitiesStore.items.find(item => item.id === route.params.id),
)
</script>

<template>
  <div>
    <template v-if="item">
      <h1>{{ item.title }}</h1>
      <BaseItemComponent
        :item="item"
        :is-list="false"
      />
    </template>
    <template v-else>
      <h1>Item not found</h1>
    </template>
  </div>
</template>
