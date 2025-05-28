<template>
  <UCard>
    <div class="flex items-center gap-2">
      <UIcon
        name="i-lucide-alert-circle"
        class="text-warning/80"
      />
      <ULink
        :to="`/items/${itemId}`"
      >
        <h2 class="text-lg font-bold">
          Empty item
        </h2>
      </ULink>
    </div>
  </UCard>
</template>

<script lang="ts" setup>
defineProps<{
  itemId: string
}>()
</script>
