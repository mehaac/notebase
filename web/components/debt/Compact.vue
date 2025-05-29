<script lang="ts" setup>
import type { BaseItemEmits } from '../BaseItem.vue'
import type { DebtData, DebtProps } from '../ItemDebt.vue'
import { ItemCard, UProgress } from '#components'

const {
  item,
  debtData,
  loading,
} = defineProps<DebtProps & { debtData: DebtData, loading?: boolean }>()

const emits = defineEmits<BaseItemEmits>()
</script>

<template>
  <ItemCard
    :item="item"
    :icon="'i-lucide-credit-card'"
    compact
    :loading="loading"
    @toggle-completed="(payload) => emits('updateFrontmatter', payload)"
  >
    <div class="pt-2">
      <UProgress
        :model-value="debtData.progress"
      />
    </div>
  </ItemCard>
</template>
