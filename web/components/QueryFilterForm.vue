<script lang="ts" setup>
import { UBadge, UButtonGroup, UCheckbox, UInput } from '#components'
import { useFiltersStore } from '#imports'

const filtersStore = useFiltersStore()
</script>

<template>
  <div class="flex flex-col gap-2 w-full">
    <UButtonGroup>
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
        class="w-full"
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
        class="w-full"
      />
    </UButtonGroup>
    <UButtonGroup>
      <UInput
        v-model="filtersStore.saveFilterLabel"
        placeholder="Filter label"
        type="text"
        class="w-full"
      />
      <UButton
        label="Save"
        variant="soft"
        icon="i-lucide-save"
        :disabled="!filtersStore.saveFilterLabel"
        @click="filtersStore.saveFilter()"
      />
    </UButtonGroup>
  </div>
</template>
