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
    window.localStorage.removeItem("userId");
  },
  loggedIn() {
    return this.getUserId() !== 0;
  },
});
