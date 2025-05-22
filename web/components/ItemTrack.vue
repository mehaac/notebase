<script setup lang="ts" generic="T extends ItemRecord & { frontmatter: TrackFrontmatter }">
import type { ItemRecord, TrackFrontmatter } from '#pocketbase-imports'

const { item } = defineProps<{ item: T, isList?: boolean }>()
</script>

<template>
  <div>
    <p>
      <ULink
        as="button"
        :to="item.frontmatter.url"
        target="_blank"
        external
      >{{ item.frontmatter.url }}</ULink>
    </p>
    <TrackIncrement
      :item="item"
      incr-key="season"
    />
    <TrackIncrement
      :item="item"
      incr-key="episode"
    />
  </div>
</template>
