<script lang="ts" setup>
import { pb } from "~/utils/pb";
import { onMounted } from "vue";
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

const validateEmail = (email: FormDataEntryValue | null): email is string => {
  if (!email || typeof email !== "string") return false;
  return /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email);
};

const validatePassword = (password: FormDataEntryValue | null): password is string => {
  if (!password || typeof password !== "string") return false;
  return password.length >= 8;
};

const onSubmit = async (e: Event) => {
  const formData = new FormData(e.target as HTMLFormElement);
  const email = formData.get("email")
  const password = formData.get("password")

  if (!validateEmail(email)) {
    console.error("Invalid email");
    return;
  }

  if (!validatePassword(password)) {
    console.error("Invalid password");
    return;
  }
  try {
    await pb
      .collection("_superusers")
      .authWithPassword(email, password);
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
      <form @submit.prevent=" (e) => onSubmit(e)" ref="form">
        <fieldset>
          <label>
            Email
            <input
              name="email"
              placeholder="Email"
              autocomplete="email"
            />
          </label>
          <label>
            Password
            <input
              type="password"
              name="password"
              placeholder="Password"
              autocomplete="current-password"
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
