import { defineNuxtPlugin } from '#app'
import { ItemDebt } from '#components'

export default defineNuxtPlugin((nuxtApp) => {
  nuxtApp.vueApp.component('debt', ItemDebt)
})
