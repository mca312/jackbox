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
      console.log(resp);
      this.user = resp;
    });
  },
};
</script>

<template>
  <div class="frame">
    <h2>Dashboard</h2>
    <div>{{ this.user.name }}</div>
    <div>{{ this.user.email }}</div>
  </div>
</template>

<style scoped>
.frame {
  border: 1px solid gray;
}

.error-input {
  border: 1px solid red;
}
</style>
