<script setup lang="ts">
import {
  definePageMeta,
  ItemType,
  useActivitiesStore,
  useRoute,
  type Item,
} from "#imports";
import { onMounted } from "vue";

definePageMeta({
  middleware: ["auth"],
});

const route = useRoute();
const activitiesStore = useActivitiesStore();
let item: Item | undefined;

onMounted(async () => {
  item = activitiesStore.items.find((item) => item.id === route.params.id);
});
</script>

<template>
  <div>
    <h1>{{ item?.title }}</h1>
    <Debt v-if="item?.type == ItemType.Debt" :item="item" :is-list="false" />
  </div>
</template>
