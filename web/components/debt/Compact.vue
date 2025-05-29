<script lang="ts" setup>
import type { BaseItemEmits } from '../BaseItem.vue'
import type { DebtData, DebtProps } from '../ItemDebt.vue'
import { ItemsListCard, UProgress } from '#components'

const {
  item,
  debtData,
  loading,
} = defineProps<DebtProps & { debtData: DebtData, loading?: boolean }>()

const emits = defineEmits<BaseItemEmits>()
</script>

<template>
  <ItemsListCard
    :item="item"
    :icon="'i-lucide-credit-card'"
    :loading="loading"
    @toggle-completed="emits('updateFrontmatter', item)"
  >
    <div class="pt-2">
      <UProgress
        :model-value="debtData.progress"
      />
    </div>
  </ItemsListCard>
</template>
