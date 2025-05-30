<script setup lang="ts">
import { defineShortcuts, useFiltersStore, useTemplateRef } from '#imports'

const filtersStore = useFiltersStore()

const input = useTemplateRef('input')

defineShortcuts({
  '/': () => {
    input.value?.inputRef?.focus()
  },
})
</script>

<template>
  <div class="flex py-4 sticky top-0 z-10 bg-(--ui-bg)">
    <UButtonGroup class="w-full">
      <AppDrawer />
      <UInput
        ref="input"
        v-model="filtersStore.query"
        icon="i-lucide-search"
        size="xl"
        variant="outline"
        placeholder="Query"
        class="w-full"
        autocapitalize="none"
        autocorrect="off"
      >
        <template #trailing>
          <UButton
            v-if="filtersStore.query.length"
            color="neutral"
            variant="link"
            size="sm"
            icon="i-lucide-circle-x"
            aria-label="Clear input"
            @click="filtersStore.query = ''"
          />
          <UKbd value="/" />
        </template>
      </UInput>
      <USelectMenu
        v-model="filtersStore.queryType"
        :items="['FTS', 'QL']"
      />
    </UButtonGroup>
  </div>
</template>
