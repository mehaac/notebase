<script setup lang="ts">
import { navigateTo } from "#app";
import { pb, ref, useActivitiesStore } from "#imports";
import { useSettingsStore } from "~/stores/settings";

const activitiesStore = useActivitiesStore();

const onLogout = async () => {
  pb.authStore.clear();
  await navigateTo({ name: "login" });
};

const queries = {
  debt: "frontmatter.type = 'debt'",
  track: "frontmatter.type = 'track'",
};

const setQuery = (query: string) => {
  activitiesStore.query = query;
};

const settingsStore = useSettingsStore();
</script>

<template>
  <nav>
    <ul>
      <li>
        <strong><NuxtLink to="/">Notebase</NuxtLink></strong>
      </li>
    </ul>
    <ul>
      <li>
        <a @click="settingsStore.toggleDialog">Settings</a>
      </li>
      <!-- <li>
        <a @click="setQuery(queries.debt)">Debt</a>
      </li>
      <li>
        <a @click="setQuery(queries.track)">Track</a>
      </li> -->
      <li v-if="pb.authStore.isValid">
        <a href="#" :onclick="onLogout">Logout</a>
      </li>
      <li v-else><NuxtLink to="/login">Login</NuxtLink></li>
    </ul>
  </nav>
</template>
