import { defineStore } from "pinia";
import { ref, watchEffect } from "vue";

import type { User } from "@/types/user";

const USER_KEY = "user";
const IS_LOGGED_IN_KEY = "isLoggedIn";

export const useUserStore = defineStore("user", () => {
  const user = ref<User>(null);
  const isLoggedIn = ref(false);

  function setUser(newUser: User) {
    user.value = newUser;
  }

  function setIsLogged(logged: boolean) {
    isLoggedIn.value = logged;
  }

  const storedUser = localStorage.getItem(USER_KEY);
  if (storedUser) {
    user.value = JSON.parse(storedUser) as User;
  }
  const storedIsLoggedIn = localStorage.getItem(IS_LOGGED_IN_KEY);
  if (storedIsLoggedIn) {
    isLoggedIn.value = JSON.parse(storedIsLoggedIn) as boolean;
  }

  const $reset = () => {
    user.value = null;
    isLoggedIn.value = false;
    localStorage.removeItem(USER_KEY);
    localStorage.removeItem(IS_LOGGED_IN_KEY);
  };

  watchEffect(() => {
    if (user.value) {
      localStorage.setItem(USER_KEY, JSON.stringify(user.value));
    }

    if (isLoggedIn.value) {
      localStorage.setItem(IS_LOGGED_IN_KEY, JSON.stringify(isLoggedIn.value));
    }
  });

  return {
    user,
    isLoggedIn,
    setUser,
    setIsLogged,
    $reset,
  };
});
