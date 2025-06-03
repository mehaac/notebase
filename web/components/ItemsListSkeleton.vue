<template>
  <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
    <UCard
      v-for="i in 6"
      :key="i"
    >
      <div class="flex flex-col gap-2">
        <div class="flex gap-2">
          <USkeleton class="h-4 w-4" />
          <USkeleton class="h-4 w-full" />
          <USkeleton class="h-4 w-1/3" />
        </div>
        <div class="py-2">
          <USkeleton class="h-3 w-1/2" />
        </div>
      </div>
    </UCard>
  </div>
</template>
