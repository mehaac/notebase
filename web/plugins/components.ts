import { defineNuxtPlugin } from '#app'
import { AppDebt } from '#components'

export default defineNuxtPlugin((nuxtApp) => {
  nuxtApp.vueApp.component('debt', AppDebt)
})
