import { defineNuxtRouteMiddleware, navigateTo } from '#app'
import { useClient, useUser } from '#imports'

export default defineNuxtRouteMiddleware(async () => {
  const pb = useClient()
  const user = useUser()
  user.value.isAuthenticated = await pb.isAuthenticated()

  if (!user.value.isAuthenticated) {
    return navigateTo({ name: 'login' })
  }
})
