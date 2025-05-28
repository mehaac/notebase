<template>
  <UCard>
    <div class="flex items-center gap-2">
      <UIcon
        name="i-lucide-alert-circle"
        class="text-warning/80"
      />
      <ULink
        :to="`/items/${item.id}`"
      >
        <h2 class="text-md font-bold">
          {{ item.id }}

        </h2>
        <p class="text-sm text-dimmed">
          This Item have data but no component
        </p>
      </ULink>
    </div>
  </UCard>
</template>

<script lang="ts" setup>
import type { ItemRecord } from '#pocketbase-imports'

defineProps<{
  item: ItemRecord
}>()
</script>
