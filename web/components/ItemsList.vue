<script lang="ts" setup>
import type { ItemRecord } from '#pocketbase-imports'

const { items } = defineProps<{ items: ItemRecord[] }>()
</script>

<template>
  <ul
    class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4"
  >
    <LazyBaseItem
      v-for="item in items"
      :key="item.id"
      :item="item"
      compact
    />
  </ul>
</template>
