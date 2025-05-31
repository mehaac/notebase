<script lang="ts" setup>
import { UBadge, UButtonGroup, UCheckbox, UInput } from '#components'
import { useFiltersStore } from '#imports'

const filtersStore = useFiltersStore()
</script>

<template>
  <div class="flex flex-col gap-2">
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
        autocapitalize="none"
        autocorrect="off"
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
        autocapitalize="none"
        autocorrect="off"
      />
    </UButtonGroup>
  </div>
</template>
