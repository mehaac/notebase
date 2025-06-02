import type { BreadcrumbItem } from '@nuxt/ui'
import { useState } from '#imports'

export const useBreadcrumbsState = () => useState<BreadcrumbItem[]>('breadcrumbs', () => [])

export function useBreadcrumbs() {
  const items = useBreadcrumbsState()

  const setBreadcrumbs = (itemsToSet: BreadcrumbItem[]) => {
    items.value = itemsToSet.map(item => ({
      label: item.label,
      icon: item.icon,
      to: item.to,
      disabled: item.disabled,
    }))
  }

  const addBreadcrumb = (itemToAdd: BreadcrumbItem) => {
    items.value.push({
      label: itemToAdd.label,
      icon: itemToAdd.icon,
      to: itemToAdd.to,
      disabled: itemToAdd.disabled,
    })
  }

  const removeBreadcrumb = (index: number) => {
    items.value.splice(index, 1)
  }

  const clearBreadcrumbs = () => {
    items.value = []
  }

  return {
    breadcrumbItems: items,
    setBreadcrumbs,
    addBreadcrumb,
    removeBreadcrumb,
    clearBreadcrumbs,
  }
}
