<script lang="ts" setup>
import { resolveDynamicComponent } from 'vue'
import type { Item } from '#imports'

const { item, isList } = defineProps<{ item: Item, isList: boolean }>()

const itemComponent = resolveDynamicComponent(item.frontmatter.type)
</script>

<template>
  <component
    :is="itemComponent"
    :item="item"
    :is-list="isList"
  />
</template>
