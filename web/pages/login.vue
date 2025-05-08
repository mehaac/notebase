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

const onSubmit = async (e: Event) => {
  const formData = new FormData(e.target as HTMLFormElement);
  const email = formData.get("email")
  const password = formData.get("password")

  if (!email || !password) {
    console.error("Invalid email or password");
    return;
  }

  try {
    await pb
      .collection("_superusers")
      .authWithPassword(email as string, password as string);
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
