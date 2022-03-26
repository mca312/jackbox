<script setup>
import { RouterLink, RouterView } from "vue-router";
import HelloWorld from "@/components/HelloWorld.vue";
import { store } from "./store";
</script>

<script>
export default {
  data() {
    return {
      store,
    };
  },
  methods: {
    logout() {
      store.clearUserId();
      this.$router.push("/");
    },
  },
};
</script>

<template>
  <header>
    <img
      alt="Vue logo"
      class="logo"
      src="@/assets/jackbox.png"
      width="196"
      height="125"
    />

    <div class="wrapper">
      <HelloWorld msg="Michael Angellotti" />
      <nav>
        <RouterLink to="/">Home</RouterLink>
        <span v-if="!store.loggedIn()">
          <RouterLink to="/login">Login</RouterLink>
          <RouterLink to="/register">Register</RouterLink>
        </span>
        <span v-else>
          <RouterLink to="/dashboard">Dashboard</RouterLink>
          <a href="#" v-on:click.prevent="logout()">Logout</a>
        </span>
      </nav>
    </div>
  </header>

  <RouterView />
</template>

<style>
@import "@/assets/base.css";

#app {
  max-width: 1280px;
  margin: 0 auto;
  padding: 2rem;

  font-weight: normal;
}

header {
  line-height: 1.5;
  max-height: 100vh;
}

.logo {
  display: block;
  margin: 0 auto 2rem;
}

a,
.green {
  text-decoration: none;
  color: hsla(160, 100%, 37%, 1);
  transition: 0.4s;
}

@media (hover: hover) {
  a:hover {
    background-color: hsla(160, 100%, 37%, 0.2);
  }
}

nav {
  width: 100%;
  font-size: 12px;
  text-align: center;
  margin-top: 2rem;
}

nav a.router-link-exact-active {
  color: var(--color-text);
}

nav a.router-link-exact-active:hover {
  background-color: transparent;
}

nav a {
  display: inline-block;
  padding: 0 1rem;
  border-left: 1px solid var(--color-border);
}

nav a:first-of-type {
  border: 0;
}

label {
  display: block;
  margin-top: 20px;
  font-size: 14px;
  font-weight: 600;
  padding: 0;
}

input,
textarea,
select,
.select {
  display: block;
  width: 100%;
  padding: 10px;
  font: inherit;
  border: 1px solid #c4c1c1;
  height: 34px;
  -webkit-appearance: none;
  border-radius: 5px;
}

.btn {
  display: inline-block;
  min-width: 60px;
  padding: 6px 10px;
  line-height: 1.3em;
  font-family: inherit;
  font-size: 14px;
  font-weight: 400;
  text-decoration: none;
  text-align: center;
  text-transform: uppercase;
  color: #fff;
  background: #555;
  border-radius: 5px;
  cursor: pointer;
  border: 0;
}

@media (min-width: 1024px) {
  body {
    display: flex;
    place-items: center;
  }

  #app {
    display: grid;
    grid-template-columns: 1fr 1fr;
    padding: 0 2rem;
  }

  header {
    display: flex;
    place-items: center;
    padding-right: calc(var(--section-gap) / 2);
  }

  header .wrapper {
    display: flex;
    place-items: flex-start;
    flex-wrap: wrap;
  }

  .logo {
    margin: 0 2rem 0 0;
  }

  nav {
    text-align: left;
    margin-left: -1rem;
    font-size: 1rem;

    padding: 1rem 0;
    margin-top: 1rem;
  }
}
</style>
