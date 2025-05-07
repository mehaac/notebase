import { defineNuxtRouteMiddleware, navigateTo } from "#app";
import { pb } from "~/utils/pb";

export default defineNuxtRouteMiddleware((to, from) => {
  if (!pb.authStore.isValid) {
    return navigateTo("/login");
  }
});
