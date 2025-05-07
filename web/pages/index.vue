<script setup lang="ts">
import {
  watchDebounced,
  useActivitiesStore,
  defineAsyncComponent,
  definePageMeta,
  type Component,
} from "#imports";
import { ref, onMounted } from "vue";

definePageMeta({
  middleware: ["auth"],
});

const activitiesStore = useActivitiesStore();

const filter = ref("");

watchDebounced(
  filter,
  async () => {
    await activitiesStore.load(filter.value);
  },
  { debounce: 500 },
);

let components = new Map<string, Component>();

onMounted(async () => {
  await activitiesStore.load(filter.value);
  for (let t in activitiesStore.itemTypes) {
    let componentName = t + "ListItem";
    let component = defineAsyncComponent(async () => {
      try {
        let comp = await import(`../components/${componentName}.vue`);
        return comp;
      } catch (e) {
        console.log("listItem component is not defined for", t);
        return await import(`../components/emptyListItem.vue`);
      }
    });
    components.set(t, component);
  }
});

// TODO: move templates importing out of the loop
const listItemComponent = (itemType: string): Component | undefined => {
  return components.get(itemType);
};
</script>

<template>
  <div>
    <input v-model="filter" type="search" placeholder="Query" />

    <article v-for="item in activitiesStore.items">
      <label>
        <input type="checkbox" name="done" checked />
        <NuxtLink :to="`/${item.id}`">
          {{ item.title }}
        </NuxtLink>
      </label>
      <component :is="listItemComponent(item.type)" :item="item" />
    </article>
  </div>
</template>
