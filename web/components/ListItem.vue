<script setup lang="ts" generic="T extends ItemRecord">
import { computed } from '#imports'
import { LazyBaseItem, UCard } from '#components'
import type { ItemRecord } from '#pocketbase-imports'
import { useActivitiesToggleItemMutation } from '~/composables/queries'

const { item } = defineProps<{ item: T }>()

const isChecked = computed(() => Boolean(item.frontmatter?.completed))

const { mutateAsync, asyncStatus } = useActivitiesToggleItemMutation()

async function toggleItem() {
  await mutateAsync(item)
}
</script>

<template>
  <UCard
    as="li"
    class="relative"
  >
    <template #header>
      <UCheckbox
        :loading="asyncStatus === 'loading'"
        :disabled="asyncStatus === 'loading'"
        :model-value="isChecked"
        @change="() => toggleItem()"
      >
        <template #label>
          <ULink :to="{ name: 'items-id', params: { id: item.id } }">
            {{ item.frontmatter?.title ?? item.frontmatter?.summary ?? 'None' }}
          </ULink>
        </template>
      </UCheckbox>
    </template>
    <LazyBaseItem
      :item="item"
      is-list
      class="ml-6"
    />
  </UCard>
</template>
