import { IS_LOGGED_STORAGE_KEY } from "@/constants";

export const isLogged = (): boolean => {
  const isLoggedIn = localStorage.getItem(IS_LOGGED_STORAGE_KEY);

  return isLoggedIn === "true" && isLoggedIn !== undefined;
};
