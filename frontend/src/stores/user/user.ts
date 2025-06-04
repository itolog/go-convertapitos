import { useCookies } from "@vueuse/integrations/useCookies";
import { defineStore } from "pinia";
import { ref, watchEffect } from "vue";

import {
  ACCESS_TOKEN,
  IS_LOGGED_STORAGE_KEY,
  USER_STORAGE_KEY,
} from "@/constants";
import type { User, UserAuth } from "@/types/user";

export const useUserStore = defineStore("user", () => {
  const user = ref<User>(null);
  const isLoggedIn = ref(false);
  const cookies = useCookies();

  function setUser(newUser: User) {
    user.value = newUser;
  }

  function setIsLogged(logged: boolean) {
    isLoggedIn.value = logged;
  }

  const storedIsLoggedIn = localStorage.getItem(IS_LOGGED_STORAGE_KEY);
  if (storedIsLoggedIn) {
    isLoggedIn.value = JSON.parse(storedIsLoggedIn) as boolean;
  }

  const $reset = () => {
    user.value = null;
    isLoggedIn.value = false;
    localStorage.removeItem(IS_LOGGED_STORAGE_KEY);
  };

  watchEffect(() => {
    const storedUser = cookies.get(USER_STORAGE_KEY) as UserAuth;
    if (storedUser) {
      user.value = storedUser.user;
      isLoggedIn.value = true;
      localStorage.setItem(ACCESS_TOKEN, storedUser.accessToken);
    }

    if (isLoggedIn.value) {
      localStorage.setItem(
        IS_LOGGED_STORAGE_KEY,
        JSON.stringify(isLoggedIn.value),
      );
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
