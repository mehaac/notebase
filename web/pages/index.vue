<script setup lang="ts">
import { useActivitiesStore, definePageMeta } from "#imports";
import { onMounted } from "vue";
import { ItemType } from "~/utils/types";

definePageMeta({
  middleware: ["auth"],
});

const activitiesStore = useActivitiesStore();

onMounted(async () => {
  await activitiesStore.load();
});
</script>

<template>
  <div>
    <input v-model="activitiesStore.query" type="search" placeholder="Query" />

    <article v-for="item in activitiesStore.items" :key="item.id">
      <label>
        <input type="checkbox" name="done" checked />
        <NuxtLink :to="`/${item.id}`">
          {{ item.title }}
        </NuxtLink>
      </label>
      <LazyDebt v-if="item.type == ItemType.Debt" :item="item" is-list />
    </article>
  </div>
</template>
