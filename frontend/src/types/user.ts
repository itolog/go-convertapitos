import { defineStore } from "pinia";
import { ref, watchEffect } from "vue";

import type { User } from "@/stores/user/types.ts";

export const useUserStore = defineStore("user", () => {
  const user = ref<User>(null);
  const isLoggedIn = ref(false);

  function setUser(newUser: User) {
    user.value = newUser;
  }

  function setIsLogged(logged: boolean) {
    isLoggedIn.value = logged;
  }

  const storedUser = localStorage.getItem("user");
  if (storedUser) {
    user.value = JSON.parse(storedUser) as User;
  }
  const storedIsLoggedIn = localStorage.getItem("isLoggedIn");
  if (storedIsLoggedIn) {
    isLoggedIn.value = JSON.parse(storedIsLoggedIn) as boolean;
  }

  watchEffect(() => {
    if (user.value) {
      localStorage.setItem("user", JSON.stringify(user.value));
    }

    if (isLoggedIn.value) {
      localStorage.setItem("isLoggedIn", JSON.stringify(isLoggedIn.value));
    }
  });

  return {
    user,
    isLoggedIn,
    setUser,
    setIsLogged,
  };
});
