<script setup lang="ts">
import { watchDebounced, useActivitiesStore, definePageMeta } from "#imports";
import { ref, onMounted } from "vue";
import { ItemType } from "~/utils/types";

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
  { debounce: 300 },
);

onMounted(async () => {
  await activitiesStore.load(filter.value);
});
</script>

<template>
  <div>
    <input v-model="filter" type="search" placeholder="Query" />

    <article v-for="item in activitiesStore.items" :key="item.id">
      <label>
        <input type="checkbox" name="done" checked />
        <NuxtLink :to="`/${item.id}`">
          {{ item.title }}
        </NuxtLink>
      </label>
      <Debt v-if="item.type == ItemType.Debt" :item="item" is-list />
    </article>
  </div>
</template>
