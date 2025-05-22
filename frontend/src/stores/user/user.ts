import { defineStore } from "pinia";
import { ref, watchEffect } from "vue";

import { IS_LOGGED_STORAGE_KEY, USER_STORAGE_KEY } from "@/constants";
import type { User } from "@/types/user";

export const useUserStore = defineStore("user", () => {
  const user = ref<User>(null);
  const isLoggedIn = ref(false);

  function setUser(newUser: User) {
    user.value = newUser;
  }

  function setIsLogged(logged: boolean) {
    isLoggedIn.value = logged;
  }

  const storedUser = localStorage.getItem(USER_STORAGE_KEY);
  if (storedUser) {
    user.value = JSON.parse(storedUser) as User;
  }
  const storedIsLoggedIn = localStorage.getItem(IS_LOGGED_STORAGE_KEY);
  if (storedIsLoggedIn) {
    isLoggedIn.value = JSON.parse(storedIsLoggedIn) as boolean;
  }

  const $reset = () => {
    user.value = null;
    isLoggedIn.value = false;
    localStorage.removeItem(USER_STORAGE_KEY);
    localStorage.removeItem(IS_LOGGED_STORAGE_KEY);
  };

  watchEffect(() => {
    if (user.value) {
      localStorage.setItem(USER_STORAGE_KEY, JSON.stringify(user.value));
    }

    if (isLoggedIn.value) {
      localStorage.setItem(IS_LOGGED_STORAGE_KEY, JSON.stringify(isLoggedIn.value));
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
