<script setup lang="ts">
import type { Item, TrackFrontmatter } from '#imports'

const { item } = defineProps<{ item: Item, isList?: boolean }>()

const fm = item.frontmatter as TrackFrontmatter
</script>

<template>
  <div>
    <p>
      <ULink
        as="button"
        :to="fm.url"
        target="_blank"
        external
      >{{ fm.url }}</ULink>
    </p>
    <UButtonGroup class="mr-2">
      <UButton
        color="warning"
        variant="subtle"
        icon="material-symbols:remove"
      />
      <UInput
        color="neutral"
        variant="outline"
        placeholder="Season"
        type="number"
        class="w-24"
      />

      <UButton
        color="success"
        variant="subtle"
        icon="material-symbols:add"
      />
    </UButtonGroup>
    <UButtonGroup>
      <UButton
        color="warning"
        variant="subtle"
        icon="material-symbols:remove"
      />
      <UInput
        color="neutral"
        placeholder="Episode"
        type="number"
        class="w-24"
      />

      <UButton
        color="success"
        variant="subtle"
        icon="material-symbols:add"
      />
    </UButtonGroup>
  </div>
</template>
