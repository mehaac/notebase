<script lang="ts" setup>
import { UButton, UContainer } from '#components'
import { navigateTo, useClient, useUser, useNow, useDateFormatter, useNotebaseConfig, definePageMeta } from '#imports'

definePageMeta({
  middleware: 'auth',
})

const pb = useClient()
const user = useUser()

const LOCALES = [
  'en-CA',
  'ru-RU',
  'de-DE',
]

const notebaseConfig = useNotebaseConfig()
const { formatShortDate } = useDateFormatter()

const onLogout = async () => {
  await pb.clearAuth()
  user.value.isAuthenticated = false
  await navigateTo({ name: 'login' })
}

const now = useNow()
</script>

<template>
  <UContainer class="flex flex-col">
    <UCard class="ring-0">
      <div class="flex flex-col">
        <div class="flex justify-between items-center pb-5">
          <div class="text-xl font-semibold">
            Profile
          </div>
          <div class="text-right">
            <div class="text-sm text-dimmed">
              Date
            </div>
            <div
              id="timeDisplay"
              class="font-medium"
            >
              {{ formatShortDate(new Date(now)) }}
            </div>
          </div>
        </div>
        <div class="flex flex-col gap-2">
          <label class="text-sm text-dimmed flex flex-col gap-0.5">
            Change date format
            <USelect
              :items="LOCALES"
              :model-value="notebaseConfig.config.value.dateLocales"
              class="w-24"
              @update:model-value="notebaseConfig.setDateLocale"
            />
          </label>
        </div>
        <div class="flex flex-col sm:flex-row gap-3 py-4">
          <UButton
            class="flex-1"
            color="error"
            variant="soft"
            block
            @click="onLogout"
          >
            Logout
          </UButton>
        </div>
      </div>
    </UCard>
  </UContainer>
</template>
