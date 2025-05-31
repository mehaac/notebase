<script lang="ts" setup>
import { useSelectedFilters } from '~/stores/filters'

const selectedFilters = useSelectedFilters()
</script>

<template>
  <div class="flex min-h-[72px] items-center gap-2 py-4">
    <div
      v-if="selectedFilters.size > 0"
      class="flex items-center gap-2 w-full"
    >
      <template v-if="selectedFilters.size === 1">
        <UButton
          color="primary"
          variant="soft"
          size="lg"
          label="Edit"
          icon="i-lucide-pencil"
        />
      </template>
      <UButton
        color="error"
        variant="soft"
        size="lg"
        label="Remove"
        icon="i-lucide-trash"
      />
      <UButton
        color="neutral"
        variant="soft"
        size="lg"
        label="Clear"
        icon="i-lucide-circle-x"
        class="ml-auto"
        @click="selectedFilters.clear()"
      />
    </div>
  </div>
</template>
