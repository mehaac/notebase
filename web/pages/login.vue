<script lang="ts" setup>
import { onMounted, reactive } from 'vue'
import { navigateTo } from '#app'
import { definePageMeta, useClient, useToast } from '#imports'

definePageMeta({
  layout: 'empty',
})

const state = reactive({
  email: '',
  password: '',
})

const pb = useClient()

const isAuthorized = defineModel('isAuthorized', {
  type: Boolean,
  default: false,
})

async function setAuthorized(value: boolean) {
  if (!value) {
    await pb.clearAuth()
  }
  isAuthorized.value = value
}

const toast = useToast()

const onSubmit = async () => {
  if (!state.email || !state.password) {
    toast.add({
      title: 'Error',
      description: 'Invalid email or password',
      color: 'error',
    })
    return
  }

  try {
    await pb.authenticatedUser({ email: state.email, password: state.password })
    setAuthorized(true)
    if (await pb.isAuthenticated()) {
      await navigateTo({ name: 'index' })
    }
  }
  catch (error) {
    toast.add({
      title: 'Error',
      description: `${error}`,
      color: 'error',
    })
  }
}

onMounted(async () => {
  if (await pb.isAuthenticated()) {
    setAuthorized(true)
  }
})
</script>

<template>
  <UContainer class="w-full h-full flex">
    <div class="m-auto">
      <h2 class="text-2xl font-bold pb-6">
        Notebase
      </h2>
      <UForm
        :state="state"
        @submit="onSubmit"
      >
        <UCard class="p-4">
          <div class="flex flex-col gap-4">
            <UFormField
              label="Email"
              name="email"
            >
              <UInput
                v-model="state.email"
              />
            </UFormField>

            <UFormField
              label="Password"
              name="password"
            >
              <UInput
                v-model="state.password"
                type="password"
              />
            </UFormField>
          </div>
          <template #footer>
            <UButton
              type="submit"
              variant="solid"
              color="primary"
              block
            >
              Login
            </UButton>
          </template>
        </UCard>
      </UForm>
    </div>
  </UContainer>
</template>
