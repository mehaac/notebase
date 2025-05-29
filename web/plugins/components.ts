import { defineNuxtPlugin } from '#app'
import { ItemDebt, ItemTrack, ItemGroceries, ItemNone } from '#components'

export default defineNuxtPlugin((nuxtApp) => {
  nuxtApp.vueApp.component('item-debt', ItemDebt)
  nuxtApp.vueApp.component('item-track', ItemTrack)
  nuxtApp.vueApp.component('item-groceries', ItemGroceries)
  nuxtApp.vueApp.component('item-none', ItemNone)
})
