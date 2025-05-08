<script setup lang="ts">
import { useSettingsStore } from "~/stores/settings";

const settingsStore = useSettingsStore();
</script>
<template>
  <dialog :open="settingsStore.dialog">
    <article>
      <header>
        <button
          aria-label="Close"
          rel="prev"
          @click="settingsStore.toggleDialog"
        ></button>
        <p>
          <strong>Settings</strong>
        </p>
      </header>
      <h5>Filters</h5>
      <form>
        <fieldset role="group">
          <input type="text" name="text" placeholder="Text" aria-label="Text" />
          <label>
            <input name="terms" type="checkbox" />
          </label>
        </fieldset>
      </form>
    </article>
  </dialog>
</template>
