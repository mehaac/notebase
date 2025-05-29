<script lang="ts" setup>
const { itemId } = defineProps<{
  itemId: string
}>()
</script>

<template>
  <ItemsListCard
    :id="itemId"
    :title="itemId"
    :icon="'i-lucide-alert-circle'"
    :icon-variant="{
      color: 'warning',
    }"
  >
    <div class="flex gap-2 items-center py-2">
      <p class="text-sm text-dimmed">
        Empty item
      </p>
    </div>
  </ItemsListCard>
</template>
