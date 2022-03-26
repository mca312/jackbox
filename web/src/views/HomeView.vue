<script setup>
import LoginBox from "../components/LoginBox.vue";
import DashboardView from "../views/DashboardView.vue";
import { store } from "../store";
</script>

<script>
export default {
  components: { LoginBox },
  data() {
    return {
      store,
    };
  },
};
</script>

<template>
  <span v-if="!store.loggedIn()">
    <LoginBox />
  </span>
  <span v-else>
    <DashboardView />
  </span>
</template>
