<script lang="ts" setup>

</script>

<template>
  <UCard class="ring-0">
    <div class="flex flex-col gap-2">
      <div class="flex gap-2">
        <USkeleton class="h-4 w-4" />
        <USkeleton class="h-4 w-56" />
      </div>
      <USkeleton class="h-2 w-28" />

      <div class="py-2 flex flex-col gap-4">
        <USkeleton
          v-for="i in 3"
          :key="i"
          :class="`h-3 w-${56 - (i-1)*14}`"
        />
      </div>
    </div>
  </UCard>
</template>
