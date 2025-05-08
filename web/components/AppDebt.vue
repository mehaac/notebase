<script setup lang="ts">
import type { Item, DebtFrontmatter } from '#imports'

const { item, isList } = defineProps<{ item: Item, isList?: boolean }>()

const frontmatter = item.frontmatter as DebtFrontmatter
</script>

<template>
  <div>
    {{ frontmatter.currency }}
    <div v-if="!isList">
      <p>graph and trnsactions</p>
    </div>
  </div>
</template>
