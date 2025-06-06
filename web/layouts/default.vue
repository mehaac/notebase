<script setup lang="ts">
</script>

<template>
  <div class="flex flex-col h-dvh w-dvw relative overflow-x-hidden">
    <UContainer class="flex py-4 backdrop-blur-xs min-h-[72px] flex-col sticky top-0 z-10">
      <AppFilters />
      <AppSettings />
    </UContainer>
    <UContainer>
      <slot />
    </UContainer>
  </div>
</template>
