<script lang="ts" setup>
import { useFiltersStore, useNewFilterSlideover } from '#imports'

const newFilterSlideover = useNewFilterSlideover()

const filtersStore = useFiltersStore()

function onSave() {
  filtersStore.saveFilter()
  newFilterSlideover.value = false
}
</script>

<template>
  <USlideover
    v-model:open="newFilterSlideover"
    side="right"
    close-icon="i-lucide-arrow-right"
  >
    <template #title>
      <div>
        <h1>
          Query Filter
        </h1>
      </div>
    </template>

    <template #description>
      <div>
        <p>
          Create a new query filter
        </p>
      </div>
    </template>
    <template #body>
      <div class="flex flex-col gap-2">
        <QueryFilterForm />
        <UInput
          v-model="filtersStore.saveFilterLabel"
          placeholder="Filter label"
          type="text"
          size="lg"
        />
        <UButton
          label="Save"
          block
          variant="soft"
          size="lg"
          class="w-24 self-end"
          @click="onSave"
        />
      </div>
    </template>
  </USlideover>
</template>
