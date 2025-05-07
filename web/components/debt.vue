<script setup lang="ts">
import type { Item, DebtFrontmatter } from "#imports";

const props = defineProps<{ item: Item }>();
props.item.frontmatter as DebtFrontmatter;
</script>

<template>
  <div>
    <h2>debt, {{ props.item.frontmatter.currency }}</h2>
  </div>
</template>
