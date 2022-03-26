import { reactive } from "vue";

export const store = reactive({
  userId: 0,
  saveUserId(id) {
    window.localStorage.setItem("userId", id);
  },
  getUserId() {
    return window.localStorage.getItem("userId") || this.userId;
  },
  clearUserId() {
    console.log('here');
    window.localStorage.removeItem("userId");
  },
  loggedIn() {
    return this.getUserId() !== 0;
  },
});
