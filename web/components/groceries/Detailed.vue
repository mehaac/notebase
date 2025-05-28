<script lang="ts" setup>
import { computed, ref } from 'vue'
import type { ItemRecord } from '#pocketbase-imports'
import type { GroceryFrontmatter } from '~/modules/pocketbase/types/schema'
import { useActivitiesUpdateItemMutation } from '~/composables/queries'

const { item, loading = false } = defineProps<{
  item: ItemRecord & { frontmatter: GroceryFrontmatter }
  loading?: boolean
}>()

const newItem = ref('')

const { mutateAsync: updateItem } = useActivitiesUpdateItemMutation()

const formattedCreatedDate = computed(() => {
  return item.frontmatter.created && typeof item.frontmatter.created === 'string'
    ? new Date(item.frontmatter.created).toLocaleDateString()
    : ''
})

const doneItems = computed(() => {
  return item.frontmatter.items.filter(item => item.done)
})

const activeItems = computed(() => {
  return item.frontmatter.items.filter(item => !item.done)
})

async function toggleItem(nameToToggle: string) {
  await updateItem({
    ...item,
    frontmatter: {
      ...item.frontmatter,
      items: item.frontmatter.items.map(item => item.name.localeCompare(nameToToggle) === 0
        ? { ...item, done: !item.done }
        : item),
    },
  })
}

async function addItem() {
  if (!newItem.value.trim()) return

  await updateItem({
    ...item,
    frontmatter: {
      ...item.frontmatter,
      items: [...item.frontmatter.items, { name: newItem.value.trim(), done: false }],
    },
  })
  newItem.value = ''
}

async function removeItem(id: string) {
  await updateItem(
    {
      ...item,
      frontmatter: {
        ...item.frontmatter,
        items: item.frontmatter.items.filter(item => item.name !== id),
      },
    },
  )
}
</script>

<template>
  <div class="flex flex-col gap-6">
    <div class="flex-1">
      <div class="flex items-center mb-1">
        <UIcon
          name="i-lucide-shopping-cart"
          class="mr-2 text-primary"
        />
        <h1 class="text-2xl font-bold">
          {{ item.frontmatter.title }}
        </h1>
      </div>

      <div class="mb-4 flex flex-wrap items-center gap-2">
        <span
          v-if="formattedCreatedDate"
          class="text-xs"
        >Created: {{ formattedCreatedDate }}</span>
      </div>

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
      v-if="item.frontmatter.items.length === 0"
      class="text-center py-8 text-gray-500"
    >
      List is empty. Add your first item!
    </div>
  </div>
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
