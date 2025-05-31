<script lang="ts" setup>
import { ref, useNotebaseConfig } from '#imports'

const notebaseConfig = useNotebaseConfig()

const open = ref(false)

const navigationItems = ref([
  {
    label: 'Home',
    icon: 'i-lucide-house',
    to: '/',
    onSelect: () => {
      open.value = false
    },
  },
  {
    label: 'Profile',
    icon: 'i-lucide-user',
    to: '/profile',
    onSelect: () => {
      open.value = false
    },
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
      </div>
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
        />
      </div>
    </template>
  </USlideover>
</template>
