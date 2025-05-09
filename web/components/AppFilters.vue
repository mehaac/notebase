<script setup lang="ts">
import { useFiltersStore } from '#imports'

const filtersStore = useFiltersStore()
</script>

<template>
  <UInput
    v-model="filtersStore.query"
    icon="i-lucide-search"
    size="xl"
    variant="outline"
    placeholder="Query"
    class="w-full mb-2"
  />

  <UButtonGroup class="mr-2">
    <UBadge
      color="neutral"
      variant="outline"
      size="lg"
      label="path"
    >
      <UCheckbox v-model="filtersStore.pathFilterEnabled" />
      path
    </UBadge>
    <UInput
      v-model="filtersStore.pathFilter"
      color="neutral"
      variant="outline"
      placeholder="inbox/activities/%"
    />
  </UButtonGroup>

  <UButtonGroup>
    <UBadge
      color="neutral"
      variant="outline"
      size="lg"
      label="type"
    >
      <UCheckbox v-model="filtersStore.typeFilterEnabled" />
      type
    </UBadge>
    <UInput
      v-model="filtersStore.typeFilter"
      color="neutral"
      variant="outline"
      placeholder="debt"
    />
  </UButtonGroup>
</template>
