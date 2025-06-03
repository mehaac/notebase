import { defineNuxtPlugin, useBreadcrumbs, useRouter } from '#imports'
import type { RouteLocationNormalizedGeneric } from 'vue-router'

const ROUTES_CONFIG = {
  'index': {
    label: 'Home',
    icon: 'i-lucide-house',
    to: '/',
  },
  'items-id': {
    label: 'Activity',
    icon: 'i-lucide-activity',
  },
  'profile': {
    label: 'Profile',
    icon: 'i-lucide-user',
  },
  'login': {
    label: 'Login',
    icon: 'i-lucide-log-in',
  },
} as const

const ROUTES_WITHOUT_BASE = new Set(['index', 'login'])

export default defineNuxtPlugin(() => {
  const router = useRouter()
  const { setBreadcrumbs } = useBreadcrumbs()

  const baseBreadcrumb = ROUTES_CONFIG.index

  const updateBreadcrumbs = (route: RouteLocationNormalizedGeneric) => {
    const routeConfig = ROUTES_CONFIG[route.name as keyof typeof ROUTES_CONFIG]
    if (!routeConfig) return

    const breadcrumbs = ROUTES_WITHOUT_BASE.has(route.name as string)
      ? [routeConfig]
      : [baseBreadcrumb, routeConfig]

    setBreadcrumbs(breadcrumbs)
  }

  router.afterEach((to) => {
    updateBreadcrumbs(to)
  })

  updateBreadcrumbs(router.currentRoute.value)
})
