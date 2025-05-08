<script setup lang="ts">
import { navigateTo } from "#app";
import { pb } from "#imports";

const onLogout = async () => {
  pb.authStore.clear();
  await navigateTo("/login");
};
</script>
<template>
  <main class="container">
    <nav>
      <ul>
        <li>
          <strong><NuxtLink to="/">Notebase</NuxtLink></strong>
        </li>
      </ul>
      <ul>
        <li v-if="pb.authStore.isValid">
          <a href="#" :onclick="onLogout">Logout</a>
        </li>
        <li v-else><NuxtLink to="/login">Login</NuxtLink></li>
      </ul>
    </nav>
    <slot />
  </main>
</template>
