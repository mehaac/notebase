import { defineNuxtPlugin } from "#app";
import Debt from "~/components/Debt.vue";


export default defineNuxtPlugin((nuxtApp) => {
  nuxtApp.vueApp.component("debt", Debt);
});
