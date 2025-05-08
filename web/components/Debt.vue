<script setup lang="ts">
import type { Item, DebtFrontmatter } from "#imports";

const { item, isList } = defineProps<{ item: Item; isList?: boolean }>();

const frontmatter = item.frontmatter as DebtFrontmatter;
</script>

<template>
  <div>
    {{ item.frontmatter.currency }}
    <div v-if="!isList">graph and trnsactions</div>
  </div>
</template>
