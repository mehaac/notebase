<script setup lang="ts">
import type { NavigationMenuItem } from '@nuxt/ui'
import { navigateTo } from '#app'
import {
  ref,
  usePocketBaseClient,
} from '#imports'

const pb = usePocketBaseClient()

const onLogout = async () => {
  pb.client.authStore.clear()
  await navigateTo({ name: 'login' })
}

const items = ref<NavigationMenuItem[]>([
  [
    {
      label: 'Notebase',
      to: { name: 'index' },
    },
  ],
  [
    {
      label: pb.client.authStore.isValid ? 'Logout' : 'Login',
      onClick: pb.client.authStore.isValid
        ? onLogout
        : () => navigateTo({ name: 'login' }),
    },
  ],
])
</script>

<template>
  <UNavigationMenu
    :items="items"
    class="w-full justify-center"
  />
</template>
