<script lang="ts" setup>
import { computed, ref, useFiltersStore, useNotebaseConfig } from '#imports'
import type { CommandPaletteGroup, CommandPaletteItem } from '@nuxt/ui'

const notebaseConfig = useNotebaseConfig()

const open = ref(false)
const filtersStore = useFiltersStore()

const filtersGroup = computed(() => {
  const items: CommandPaletteItem[] = []
  filtersStore.localFilters.forEach((filter) => {
    items.push({
      id: filter.id,
      label: filter.label,
      onSelect: () => {
        filtersStore.applyFilter(filter.id)
      },
    })
  })
  return items
})

const tempFilters = ref<CommandPaletteGroup<CommandPaletteItem>[]>([
  {
    id: 'Quick filter',
    // @ts-expect-error - filtersGroup is a computed ref
    items: filtersGroup,
  },
  {
    id: 'Actions',
    label: 'Actions',
    slot: 'actions' as const,
    ignoreFilter: true,
    items: [
      {
        // label: 'Save',
        // suffix: 'Save active filter',
        // icon: 'i-lucide-save',
        slot: 'save' as const,
        // onSelect: () => {
        //   filtersStore.saveFilter()
        // },
      },
      {
        // label: 'Delete',
        // suffix: 'Delete active filter',
        // icon: 'i-lucide-trash',
        slot: 'delete' as const,

        // onSelect: () => {
        //   filtersStore.deleteFilter()
        // },
      },
      {
        // label: 'Clear',
        // suffix: 'Clear all filters',
        // icon: 'i-lucide-x',
        slot: 'clear' as const,
        // onSelect: () => {
        //   filtersStore.clearFilters()
        // },
      },

    ],
  },
])
</script>

<template>
  <UDrawer
    v-model:open="open"
    direction="left"
  >
    <UButton
      icon="i-lucide-menu"
      color="neutral"
      variant="subtle"
      block
      class="w-16"
    />

    <template #title>
      <div class="flex items-center gap-2">
        <ULink to="/">
          Notebase
        </ULink>
        <UButton
          icon="i-lucide-x"
          variant="ghost"
          class="ml-auto"
          color="neutral"
          @click="open = false"
        />
      </div>
    </template>
    <template #description>
      Your notes, your way!
    </template>
    <template #body>
      <div class="flex flex-col gap-2 h-full">
        <div class="flex flex-col gap-2">
          <FilterForm />
          <UCommandPalette
            v-model:search-term="filtersStore.createdFilterLabel"
            :groups="tempFilters"
            autofocus

            placeholder="Search filters"
            class="h-72"
          >
            <template #save>
              <UButton
                icon="i-lucide-save"
                color="success"
                variant="ghost"
                class="w-full"
                :disabled="!filtersStore.createdFilterLabel"
                @click="filtersStore.saveFilter()"
              >
                Save  <span class="text-xs text-dimmed">Save active filter</span>
              </UButton>
            </template>
            <template #delete>
              <UButton
                icon="i-lucide-trash"
                variant="ghost"
                color="error"
                class="w-full"
                :disabled="!filtersStore.appliedFilterId"
                @click="filtersStore.deleteFilter()"
              >
                Delete filter
                <span class="text-xs text-dimmed">Delete active filter</span>
              </UButton>
            </template>
            <template #clear>
              <UButton
                icon="i-lucide-x"
                color="warning"
                variant="ghost"
                class="w-full"
                @click="filtersStore.clearFilters()"
              >
                Clear filters
                <span class="text-xs text-dimmed">Clear filter</span>
              </UButton>
            </template>
          </UCommandPalette>
        </div>
        <USeparator class="my-2" />

        <div class="flex flex-col gap-2 mt-auto">
          <UCheckbox
            v-model="notebaseConfig.config.value.showFilters"
            label="Show filters"
          />
          <UCheckbox
            v-model="notebaseConfig.config.value.showExtra"
            label="Show extra"
          />
        </div>
      </div>
    </template>

    <template #footer>
      <div class="flex flex-col gap-2 pb-8">
        <USeparator class="my-2" />
        <UButton
          to="/profile"
          class="flex items-center gap-2 text-text"
          variant="ghost"
        >
          <UIcon
            name="i-lucide-user"
          />
          <span>
            user@user.com
          </span>
        </UButton>
        <UButton
          color="error"
          variant="ghost"
          label="Logout"
          icon="i-lucide-log-out"
        />
      </div>
    </template>
  </UDrawer>
</template>
