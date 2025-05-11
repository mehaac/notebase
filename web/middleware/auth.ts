import { defineNuxtRouteMiddleware, navigateTo } from '#app'
import { usePocketBaseClient } from '#imports'

export default defineNuxtRouteMiddleware(() => {
  const pb = usePocketBaseClient()
  if (!pb.client.authStore.isValid) {
    return navigateTo({ name: 'login' })
  }
})
