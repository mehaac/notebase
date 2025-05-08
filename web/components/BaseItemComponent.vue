<script lang="ts" setup>
import { resolveComponent } from "vue";
import type { Item } from "~/utils/types";

const { item, isList } = defineProps<{ item: Item; isList: boolean }>();

const itemComponent = resolveComponent(item.frontmatter.type);
</script>

<template>
  <component :is="itemComponent" :item="item" :is-list="isList" />
</template>
