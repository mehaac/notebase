<script lang="ts" setup>
import { UBadge, UButtonGroup, UCheckbox, UInput } from '#components'
import { ref, useFiltersStore } from '#imports'

const filtersStore = useFiltersStore()
const deleteModal = ref(false)
function handleDeleteFilter() {
  const success = filtersStore.deleteFilter(filtersStore.appliedFilterId!)
  if (!success) {
    return
  }
  deleteModal.value = false
}
</script>

<template>
  <div class="flex flex-col gap-2 w-full">
    <UButtonGroup>
      <UBadge
        color="neutral"
        variant="outline"
        size="lg"
        label="path"
      >
        <UCheckbox v-model="filtersStore.pathFilterEnabled" />
        path
      </UBadge>
      <UInput
        v-model="filtersStore.pathFilter"
        color="neutral"
        variant="outline"
        placeholder="inbox/activities/%"
        autocapitalize="none"
        autocorrect="off"
        class="w-full"
      />
    </UButtonGroup>

    <UButtonGroup>
      <UBadge
        color="neutral"
        variant="outline"
        size="lg"
        label="type"
      >
        <UCheckbox v-model="filtersStore.typeFilterEnabled" />
        type
      </UBadge>
      <UInput
        v-model="filtersStore.typeFilter"
        color="neutral"
        variant="outline"
        placeholder="debt"
        autocapitalize="none"
        autocorrect="off"
        class="w-full"
      />
    </UButtonGroup>
    <UButtonGroup>
      <UInput
        v-model="filtersStore.saveFilterLabel"
        placeholder="Filter label"
        type="text"
        class="w-full"
      />
      <UButton
        label="Save"
        variant="soft"
        icon="i-lucide-save"
        :disabled="!filtersStore.saveFilterLabel"
        @click="filtersStore.saveFilter()"
      />
      <UButton
        label="Delete"
        variant="soft"
        color="error"
        :disabled="!filtersStore.appliedFilterId"
        icon="i-lucide-trash"
        @click="deleteModal = true"
      />
    </UButtonGroup>
    <UModal
      v-model:open="deleteModal"
      :dismissible="false"
    >
      <template #title>
        Delete filter
      </template>
      <template #description>
        Are you sure you want to delete this filter?
      </template>
      <template #footer>
        <div class="flex gap-2">
          <UButton
            color="neutral"
            variant="soft"
            @click="deleteModal = false"
          >
            Cancel
          </UButton>
          <UButton
            color="error"
            @click="handleDeleteFilter"
          >
            Delete
          </UButton>
        </div>
      </template>
    </UModal>
  </div>
</template>
