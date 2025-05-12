<script setup lang="ts">
import type { NavigationMenuItem } from '@nuxt/ui'
import { navigateTo } from '#app'
import {
  ref,
  useClient,
  useUser,
} from '#imports'

const pb = useClient()
const user = useUser()

const onLogout = async () => {
  await pb.clearAuth()
  user.value.isAuthenticated = false
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
      slot: 'login' as const,
    },
  ],
])
</script>

<template>
  <UNavigationMenu
    :items="items"
    class="w-full justify-center"
  >
    <template #login-label>
      <div @click="onLogout">
        {{ user.isAuthenticated ? 'Logout' : 'Login' }}
      </div>
    </template>
  </UNavigationMenu>
</template>
