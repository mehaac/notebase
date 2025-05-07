<script setup lang="ts">
import { definePageMeta, useActivitiesStore, useRoute } from "#imports";
import { onMounted } from "vue";

definePageMeta({
  middleware: ["auth"],
});

const route = useRoute();
const activitiesStore = useActivitiesStore();

onMounted(async () => {
  activitiesStore.setItem(route.params.id as string);
});
</script>

<template>
  <div>
    <h1>{{ activitiesStore.item?.title }}</h1>
    <Debt
      v-if="activitiesStore.item?.type === 'debt'"
      :item="activitiesStore.item"
    />
  </div>
</template>
