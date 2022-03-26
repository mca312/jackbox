<script>
import { postRegistration } from "../api/api";
import { store } from "../store";

export default {
  data() {
    return {
      error: "",
    };
  },
  methods: {
    onSubmit() {
      this.error = "";
      let user = {
        email: this.email,
        name: this.name,
      };
      postRegistration(user)
        .then((resp) => {
          store.saveUserId(resp.id);
          this.$router.push("/");
        })
        .catch((err) => {
          if (err?.error) {
            this.error = err.error;
          } else {
            this.error = "Error during registration.";
          }
        });
    },
  },
};
</script>

<template>
  <div class="frame">
    <h2>Register</h2>
    <form @submit.prevent="onSubmit">
      <div>
        <label>Name:</label>
        <input v-model="name" required />
      </div>
      <div>
        <label>Email:</label>
        <input v-model="email" required placeholder="your@email.com" />
        <span v-if="error">{{ error }}</span>
      </div>
      <button class="btn" type="submit">Regsiter</button>
    </form>
  </div>
</template>

<style scoped>
.frame {
  border: 1px solid gray;
}
</style>
