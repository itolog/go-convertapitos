import { defineStore } from "pinia";
import { ref } from "vue";
import type { UserUser } from "@/generated/apiClient/data-contracts";

export const useUserStore = defineStore("user", () => {
  const user = ref<UserUser | null>(null);
  const isLoggedIn = ref(false);

  function setUser(newUser: UserUser | null) {
    user.value = newUser;
  }

  function setIsLogged(logged: boolean) {
    isLoggedIn.value = logged;
  }

  return {
    user,
    isLoggedIn,
    setUser,
    setIsLogged,
  };
});
