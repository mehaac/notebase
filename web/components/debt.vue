<script setup lang="ts">
import type { Item, DebtFrontmatter } from "#imports";

const props = defineProps<{ item: Item; isList: boolean }>();

props.item.frontmatter as DebtFrontmatter;
</script>

<template>
  <div>
    {{ props.item.frontmatter.currency }}
    <div v-if="!props.isList">graph and trnsactions</div>
  </div>
</template>
