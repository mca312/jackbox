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
    getUser(store.getUserId()).then((resp) => {
      this.user = resp;
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

<style scoped>

</style>
