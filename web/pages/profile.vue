<script lang="ts" setup>
import { UButton, UContainer } from '#components'
import { navigateTo, useClient, useUser } from '#imports'

const pb = useClient()
const user = useUser()

const onLogout = async () => {
  await pb.clearAuth()
  user.value.isAuthenticated = false
  await navigateTo({ name: 'login' })
}
</script>

<template>
  <UContainer>
    <UButton @click="onLogout">
      Logout
    </UButton>
  </UContainer>
</template>
