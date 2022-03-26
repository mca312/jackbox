<script>
import { postLogin } from "../api/api";
import { store } from "../store";

export default {
  data() {
    return {
      error: "",
      email: "",
    };
  },
  computed: {
    classObject() {
      return this.error.length > 0 ? "is-invalid" : "";
    },
  },
  methods: {
    onSubmit() {
      let user = {
        email: this.email,
      };
      postLogin(user)
        .then((resp) => {
          store.saveUserId(resp.id);
          this.$router.push("/dashboard");
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
  <div class="col">
    <h2>Login</h2>
    <form @submit.prevent="onSubmit">
      <div class="form-floating mb-3">
        <input
          id="emailInput"
          class="form-control"
          :class="classObject"
          aria-describedby="emailValidation"
          v-model="email"
          required
          placeholder="your@email.com"
        />
        <label for="emailInput">Email:</label>
        <div id="emailValidation" class="invalid-feedback">
          {{ error }}
        </div>
      </div>
      <button class="btn btn-primary" type="submit">Login</button>
    </form>
  </div>
</template>

<style scoped></style>
