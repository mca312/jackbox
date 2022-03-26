<script>
import { postRegistration } from "../api/api";
import { store } from "../store";

export default {
  data() {
    return {
      error: "",
    };
  },
  computed: {
    classObject() {
      return this.error.length > 0 ? "is-invalid" : "";
    },
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
  <div class="col">
    <h2>Register</h2>
    <form @submit.prevent="onSubmit">
      <div class="form-floating mb-3">
        <input
          id="nameInput"
          class="form-control"
          aria-describedby="nameValidation"
          v-model="name"
          required
        />
        <label for="nameInput">Name:</label>
      </div>
      <div class="form-floating mb-3">
        <input
          id="emailInput"
          class="form-control"
          :class="classObject"
          aria-describedby="emailValidation"
          v-model="email"
          required
        />
        <label for="emailInput">Email:</label>
        <div id="emailValidation" class="invalid-feedback">
          {{ error }}
        </div>
      </div>
      <button class="btn btn-primary" type="submit">Regsiter</button>
    </form>
  </div>
</template>

<style scoped></style>
