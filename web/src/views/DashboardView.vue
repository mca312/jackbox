<script>
import { getUser } from "../api/api";
import { store } from "../store";

export default {
  data() {
    return {
      error: "",
      user: {},
    };
  },

  created() {
    getUser(store.getUserId())
      .then((resp) => {
        this.user = resp;
      })
      .catch((err) => {
        // If proper authentication was implemented. I'd catch 401 unauthorized at the api level (within callAPI in api.js) and handle it there.
        // Since I have no real authentication, check the message and redirect.
        if (err?.error && err.error === "Invalid id.") {
          store.clearUserId();
          this.$router.push("/");
        }
      });
  },
};
</script>

<template>
  <div class="col">
    <h2>Dashboard</h2>
    <div>Name: {{ this.user.name }}</div>
    <div>Email: {{ this.user.email }}</div>
  </div>
</template>

<style scoped></style>
