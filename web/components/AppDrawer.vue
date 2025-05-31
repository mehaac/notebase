<script lang="ts" setup>
import { navigateTo, useClient, useFiltersStore, useNewFilterSlideover, useNotebaseConfig, useUser } from '#imports'

import { ref } from 'vue'

const notebaseConfig = useNotebaseConfig()
const filtersStore = useFiltersStore()
const newFilterSlideover = useNewFilterSlideover()

const open = ref(false)

const pb = useClient()
const user = useUser()

const onLogout = async () => {
  await pb.clearAuth()
  user.value.isAuthenticated = false
  await navigateTo({ name: 'login' })
}

const closeSlideover = () => {
  open.value = false
}

const openSlideover = () => {
  open.value = true
}

const navigationItems = ref([
  {
    label: 'Home',
    icon: 'i-lucide-house',
    to: '/',
    onSelect: closeSlideover,
  },
  {
    label: 'Profile',
    icon: 'i-lucide-user',
    to: '/profile',
    onSelect: closeSlideover,
  },
])
</script>

<template>
  <USlideover
    v-model:open="open"
    side="left"
    :ui="{
      content: 'max-w-[250px]',
    }"
    @after:leave="closeSlideover"
  >
    <UButton
      icon="i-lucide-menu"
      color="neutral"
      variant="subtle"
      block
      class="w-16"
      @click="openSlideover"
    />

    <template #title>
      <div class="flex items-center gap-2">
        <ULink to="/">
          Notebase
        </ULink>
      </div>
    </template>
    <template #close>
      <UButton
        icon="i-lucide-x"
        color="neutral"
        variant="ghost"
        class="ml-auto"
        @click="closeSlideover"
      />
    </template>
    <template #description>
      Your notes, your way!
    </template>
    <template #body>
      <div class="flex flex-col gap-2 h-full">
        <UNavigationMenu

          :items="navigationItems"
          orientation="vertical"
        />
        <USeparator class="my-2" />
        <div class="flex flex-col gap-2">
          <h2 class="font-medium">
            Saved filters
          </h2>
          <UInput
            v-model="filtersStore.searchFiltersLabel"
            placeholder="Search filters"
            type="text"
            size="lg"
            class="w-full"
          />
          <div class="flex flex-col gap-2 max-h-[180px] overflow-y-auto">
            <div
              v-for="filter in filtersStore.filteredQueryFilters"
              :key="filter.id"
              class="flex gap-2"
            >
              <UButton
                color="neutral"
                variant="soft"
                size="lg"
                :label="filter.label"
                class="w-full"
                @click="filtersStore.applyFilter(filter.id)"
              />
              <UButton
                icon="i-lucide-trash"
                color="error"
                variant="soft"
                size="lg"
                block
                class="w-12"
                @click="filtersStore.deleteFilter(filter.id)"
              />
            </div>
          </div>
        </div>
        <USeparator class="my-2" />
        <div class="flex flex-col gap-2">
          <UButton
            icon="i-lucide-plus"
            color="neutral"
            variant="soft"
            block
            size="lg"
            @click="newFilterSlideover = true"
          >
            Add new filter
          </UButton>
          <UButton
            icon="i-lucide-x"
            color="warning"
            variant="soft"
            block
            size="lg"
            @click="filtersStore.clearFilters"
          >
            Clear filters
          </UButton>
        </div>
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
      <div class="flex flex-col gap-2 pb-4">
        <UButton
          to="/profile"
          class="flex items-center gap-2 text-text"
          variant="ghost"
        >
          <UIcon
            name="i-lucide-user"
            size="16"
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
          @click="onLogout"
        />
      </div>
    </template>
  </USlideover>
</template>
