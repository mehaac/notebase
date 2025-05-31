<script setup lang="ts">
import { useSelectedFilters } from '~/stores/filters'

const selectedFilters = useSelectedFilters()
</script>

<template>
  <UContainer class="flex flex-col relative">
    <div class="min-h-[72px]">
      <Transition name="slide-right">
        <template v-if="selectedFilters.size > 0">
          <QueryFiltersPanel />
        </template>
        <template v-else>
          <AppFilters />
        </template>
      </Transition>
    </div>
    <AppSettings />
    <slot />
  </UContainer>
</template>

<style>
.slide-right-enter-active,
.slide-right-leave-active {
  transition: all 0.2s ease-out;
  position: absolute;
  width: 100%;
}

.slide-right-enter-from {
  transform: translateX(100%);
  opacity: 0;
}

.slide-right-leave-to {
  transform: translateX(-100%);
  opacity: 0;
}

.slide-right-enter-to,
.slide-right-leave-from {
  transform: translateX(0);
  opacity: 1;
}
</style>
