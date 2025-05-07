<script setup lang="ts">
import type { Item, DebtFrontmatter } from "#imports";

const props = defineProps<{ item: Item }>();
props.item.frontmatter as DebtFrontmatter;
</script>

<template>
  <div>
    {{ props.item.frontmatter.currency }}
  </div>
</template>
