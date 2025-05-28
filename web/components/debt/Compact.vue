<script lang="ts" setup>
import { computed } from 'vue'
import type { BaseItemEmits } from '../BaseItem.vue'
import type { DebtData, DebtProps } from '../ItemDebt.vue'
import { ItemsListCard, UProgress } from '#components'

const {
  item,
  debtData,
  loading,
} = defineProps<DebtProps & { debtData: DebtData, loading?: boolean }>()

const emits = defineEmits<BaseItemEmits>()
const isChecked = computed(() => Boolean(item.frontmatter?.completed))
const title = computed(() => item.frontmatter.title || item.frontmatter.summary || 'None')
</script>

<template>
  <ItemsListCard
    :id="item.id"
    :title="title"
    :icon="'i-lucide-credit-card'"
    :checked="isChecked"
    :loading="loading"
    @done="emits('done', item.id)"
  >
    <div class="pt-2">
      <UProgress
        :model-value="debtData.progress"
      />
    </div>
  </ItemsListCard>
</template>
