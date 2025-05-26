<script lang="ts" setup>
import { resolveDynamicComponent } from 'vue'
import type { ItemRecord } from '#pocketbase-imports'

const { item, isList } = defineProps<{ item: ItemRecord, isList: boolean }>()

const itemComponent = resolveDynamicComponent('item-' + (item?.frontmatter?.type ?? 'none'))
</script>

<template>
  <component
    :is="itemComponent"
    :item="item"
    :is-list="isList"
  />
</template>
