<script setup lang="ts">
import {
  definePageMeta,
  useActivitiesStore,
  useRoute,
} from "#imports";
import { computed } from "vue";

definePageMeta({
  middleware: ["auth"],
});

const route = useRoute();
const activitiesStore = useActivitiesStore();
const item = computed(() => 
  activitiesStore.items.find((item) => item.id === route.params.id)
);


</script>

<template>
  <div>
    <h1>{{ item?.title }}</h1>
    <ItemComponent
      :item="item"
      :is-list="false"
    />
  </div>
</template>
