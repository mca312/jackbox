<template>
  <div>
    <h2>Login</h2>
    <form @submit.prevent="onSubmit">
      <label>Email:</label>
      <input v-model="email" required placeholder="your@email.com" />
      <button class="btn btn-success btn-block my-2" type="submit">
        Login
      </button>
    </form>
  </div>
</template>
