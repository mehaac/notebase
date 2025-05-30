<script lang="ts" setup>
import { computed, ref } from 'vue'
import type { ItemRecord } from '#pocketbase-imports'
import type { GroceriesFrontmatter } from '~/modules/pocketbase/types/schema'
import type { BaseItemEmits } from '../BaseItem.vue'

const { item, loading = false } = defineProps<{
  item: ItemRecord & { frontmatter: GroceriesFrontmatter }
  loading?: boolean
}>()

const emits = defineEmits<BaseItemEmits>()

const newItem = ref('')

const doneItems = computed(() => {
  return item.frontmatter.checklist.filter(item => item.done)
})

const activeItems = computed(() => {
  return item.frontmatter.checklist.filter(item => !item.done)
})

async function toggleItem(nameToToggle: string) {
  const payload = { ...item,
    frontmatter: {
      ...item.frontmatter,
      checklist: item.frontmatter.checklist.map(item => item.name.localeCompare(nameToToggle) === 0
        ? { ...item, done: !item.done }
        : item),
    },
  }
  emits('updateFrontmatter', payload)
}

async function addItem() {
  if (!newItem.value.trim()) return
  const payload = {
    ...item,
    frontmatter: {
      ...item.frontmatter,
      checklist: [...item.frontmatter.checklist, { name: newItem.value.trim(), done: false }],
    },
  }
  emits('updateFrontmatter', payload)
  newItem.value = ''
}

async function removeItem(id: string) {
  const payload = {
    ...item,
    frontmatter: {
      ...item.frontmatter,
      checklist: item.frontmatter.checklist.filter(item => item.name !== id),
    },
  }
  emits('updateFrontmatter', payload)
}
</script>

<template>
  <ItemCard
    :item="item"
    :icon="'i-lucide-shopping-cart'"
    :loading="loading"
    @toggle-completed="(payload) => emits('updateFrontmatter', payload)"
  >
    <div class="flex-1">
      <!-- Add new item -->
      <div class="mb-6">
        <div class="flex gap-2">
          <UInput
            v-model="newItem"
            placeholder="Add item..."
            :disabled="loading"
            @keyup.enter="addItem"
          />
          <UButton
            color="primary"
            icon="i-lucide-plus"
            :loading="loading"
            :disabled="!newItem.trim()"
            @click="addItem"
          >
            Add
          </UButton>
        </div>
      </div>

      <div class="space-y-2">
        <h2 class="text-lg font-bold">
          Active
        </h2>
        <TransitionGroup
          name="list"
          tag="div"
          class="relative space-y-2"
        >
          <div
            v-for="groceryItem in activeItems"
            :key="groceryItem.name"
            class="transform transition-all duration-300 ease-out will-change-transform backface-hidden"
          >
            <GroceriesDetailedItem
              :grocery-item="groceryItem"
              :loading="loading"
              @toggle="toggleItem"
              @remove="removeItem"
            />
          </div>
        </TransitionGroup>
      </div>

      <USeparator class="my-4" />

      <div class="space-y-2">
        <h2 class="text-lg font-bold">
          Done
        </h2>
        <TransitionGroup
          name="list"
          tag="div"
          class="relative space-y-2"
        >
          <div
            v-for="groceryItem in doneItems"
            :key="groceryItem.name"
            class="transform transition-all duration-300 ease-out will-change-transform backface-hidden"
          >
            <GroceriesDetailedItem
              :grocery-item="groceryItem"
              :loading="loading"
              @toggle="toggleItem"
              @remove="removeItem"
            />
          </div>
        </TransitionGroup>
      </div>
    </div>

    <div
      v-if="item.frontmatter.checklist.length === 0"
      class="text-center py-8 text-gray-500"
    >
      List is empty. Add your first item!
    </div>
  </ItemCard>
</template>

<style scoped>
.list-enter-active,
.list-leave-active,
.list-move {
  transition: all 0.3s ease;
}

.list-enter-from,
.list-leave-to {
  opacity: 0;
  transform: translateY(10px);
}

.list-enter-to,
.list-leave-from {
  opacity: 1;
  transform: translateY(0);
}

.list-leave-active {
  position: absolute;
}
</style>
