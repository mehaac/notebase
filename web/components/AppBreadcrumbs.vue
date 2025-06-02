<script setup lang="ts">
import { useBreadcrumbs } from '#imports'

const { breadcrumbItems } = useBreadcrumbs()
</script>

<template>
  <UBreadcrumb
    v-if="breadcrumbItems.length > 0"
    separator-icon="i-lucide-chevron-right"
    :items="breadcrumbItems"
    class="py-4"
  />
</template>
