<script lang="ts" setup>
import { pb } from "~/utils/pb";
import { useTemplateRef, onMounted } from "vue";
import { navigateTo } from "#app";

const isAuthorized = defineModel("isAuthorized", {
  type: Boolean,
  default: false,
});

function setAuthorized(value: boolean) {
  if (!value) {
    pb.authStore.clear();
  }
  isAuthorized.value = value;
}

const email = useTemplateRef("email");
const password = useTemplateRef("password");

const onSubmit = async (e: Event) => {
  const emailValue = email.value?.value;
  const passwordValue = password.value?.value;
  if (!emailValue || !passwordValue) return;
  try {
    await pb
      .collection("_superusers")
      .authWithPassword(emailValue, passwordValue);
    setAuthorized(true);
    if (pb.authStore.isValid) {
      await navigateTo({ name: "index" });
    }
  } catch (error) {
    console.error(error);
  }
};

const onLogout = async () => {
  setAuthorized(false);
  await navigateTo({ name: "login" });
};

onMounted(() => {
  if (pb.authStore.isValid) {
    setAuthorized(true);
  }
});
</script>

<template>
  <div class="grid">
    <div></div>
    <div v-if="!isAuthorized">
      <form @submit.prevent="onSubmit">
        <fieldset>
          <label>
            Email
            <input
              name="email"
              placeholder="Email"
              autocomplete="email"
              ref="email"
            />
          </label>
          <label>
            Password
            <input
              type="password"
              name="password"
              placeholder="Password"
              autocomplete="current-password"
              ref="password"
            />
          </label>
        </fieldset>
        <input type="submit" value="Login" />
      </form>
    </div>
    <div v-else>
      <button @click="onLogout">Logout</button>
    </div>
    <div></div>
  </div>
</template>
