<script setup lang="ts">
import type { NavigationMenuItem } from '@nuxt/ui'
import { ref } from '#imports'

const items = ref<NavigationMenuItem[]>([
  [
    {
      label: 'Notebase',
      to: { name: 'index' },
    },
  ],
  [
    {
      label: 'Profile',
      to: { name: 'profile' },
    },
  ],
])
</script>

<template>
  <UNavigationMenu
    :items="items"
    class="w-full justify-start py-2"
  />
</template>
