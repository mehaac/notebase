<script setup lang="ts">
import { onMounted, shallowRef, ref, watch, onUnmounted } from 'vue'
import type { MDCParserResult } from '@nuxtjs/mdc'
import {
  definePageMeta,
  useRoute,
  useMarkdownParser,
} from '#imports'
import { BaseItem } from '#components'
import { useActivitiesItemQuery } from '~/composables/queries/'

definePageMeta({
  middleware: ['auth'],
  layout: 'activity',
})

const route = useRoute()
const parseMd = useMarkdownParser()
const contentAst = ref<MDCParserResult | null>(null)
const frontmatterAst = ref<MDCParserResult | null>(null)

const error = shallowRef<string>()

const { state, id } = useActivitiesItemQuery()

watch(() => state.value.status, async (status) => {
  if (status === 'success') {
    contentAst.value = await parseMd(state.value.data?.content ?? '\n')
    frontmatterAst.value = await parseMd(
      '```json\n' + JSON.stringify(state.value.data?.frontmatter ?? {}, null, 2) + '\n```',
    )
  }
}, { immediate: true })

onMounted(async () => {
  if (typeof route.params.id !== 'string') {
    error.value = 'Invalid item id'
    return
  }
  id.value = route.params.id
})
onUnmounted(() => {
  id.value = undefined
})
</script>

<template>
  <UContainer class="flex flex-col relative overflow-x-hidden">
    <template v-if="state.status === 'pending'">
      <h1>Loading...</h1>
    </template>
    <template v-else-if="state.status === 'success'">
      <BaseItem
        :item="state.data"
      />
      <hr class="my-4">
      <UCollapsible v-if="frontmatterAst?.body">
        <UButton
          class="group"
          label="Frontmatter"
          color="neutral"
          variant="subtle"
          trailing-icon="i-lucide-chevron-down"
          :ui="{
            trailingIcon: 'group-data-[state=open]:rotate-180 transition-transform duration-200',
          }"
          block
        />

        <template #content>
          <Suspense>
            <MDCRenderer
              v-if="frontmatterAst?.body"
              :body="frontmatterAst.body"
              :data="frontmatterAst.data"
            />
          </Suspense>
        </template>
      </UCollapsible>
      <UCollapsible
        v-if="contentAst?.body"
        class="mt-4"
      >
        <UButton
          class="group"
          label="Content"
          color="neutral"
          variant="subtle"
          trailing-icon="i-lucide-chevron-down"
          :ui="{
            trailingIcon: 'group-data-[state=open]:rotate-180 transition-transform duration-200',
          }"
          block
        />

        <template #content>
          <Suspense>
            <MDCRenderer
              v-if="contentAst?.body"
              :body="contentAst.body"
              :data="contentAst.data"
              class="prose dark:prose-invert"
            />
          </Suspense>
        </template>
      </UCollapsible>
    </template>
    <template v-else-if="error || state.status === 'error'">
      <h1>{{ error || state.error }}</h1>
    </template>
  </UContainer>
</template>
