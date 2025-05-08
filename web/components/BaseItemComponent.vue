<script lang="ts" setup>
import { type Item } from "#imports";
import { resolveDynamicComponent } from "vue";

const { item, isList } = defineProps<{ item: Item; isList: boolean }>();

const itemComponent = resolveDynamicComponent(item.frontmatter.type);
</script>

<template>
  <component :is="itemComponent" :item="item" :is-list="isList" />
</template>
