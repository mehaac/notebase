<script setup lang="ts">
import { useActivitiesStore, definePageMeta } from "#imports";
import { onMounted } from "vue";
import { BaseItemComponent } from "#components";

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
    <UInput
      icon="i-lucide-search"
      size="xl"
      variant="outline"
      placeholder="Query"
      class="w-full mb-2"
      v-model="activitiesStore.query"
    />

    <UButtonGroup class="mr-2">
      <UBadge color="neutral" variant="outline" size="lg" label="path">
        <UCheckbox />
        path
      </UBadge>
      <UInput
        color="neutral"
        variant="outline"
        placeholder="inbox/activities/%"
      />
    </UButtonGroup>

    <UButtonGroup>
      <UBadge color="neutral" variant="outline" size="lg" label="type">
        <UCheckbox />
        type
      </UBadge>
      <UInput color="neutral" variant="outline" placeholder="debt" />
    </UButtonGroup>

    <UCard v-for="item in activitiesStore.items" :key="item.id" class="mt-4">
      <UCheckbox>
        <template #label>
          <ULink :to="{ name: 'items-id', params: { id: item.id } }">
            {{ item.title }}
          </ULink>
          <BaseItemComponent :item="item" is-list />
        </template>
      </UCheckbox>
    </UCard>
  </div>
</template>
