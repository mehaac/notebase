<script setup lang="ts" generic="T extends Item">
import { useClient, ref, useActivitiesStore, computed } from '#imports'
import { LazyBaseItem } from '#components'
import type { Item } from '#pocketbase-imports'

const { item } = defineProps<{ item: T }>()

const pb = useClient()
const activitiesStore = useActivitiesStore()
const isChecked = computed(() => item.done)
const loading = ref(false)
// TODO: fix this must be handled by state machine that awaits for responce
async function toggleItem() {
  loading.value = true
  try {
    const result = await pb.toggleItem(item.id)
    if (result) {
      activitiesStore.items = activitiesStore
        .items.map(item => item.id === result.id ? result : item)
    }
  }
  catch (err) {
    console.log('error', err)
  }
  loading.value = false
}
</script>

<template>
  <UCheckbox
    :loading="loading"
    :disabled="loading"
    :model-value="isChecked"
    @change="() => toggleItem()"
  >
    <template #label>
      <ULink :to="{ name: 'items-id', params: { id: item.id } }">
        {{ item.title }}
      </ULink>
    </template>
  </UCheckbox>
  <LazyBaseItem
    :item="item"
    is-list
    class="ml-6"
  />
</template>
