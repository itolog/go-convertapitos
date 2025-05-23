import { type MutationOptions, useMutation } from "@tanstack/vue-query";
import type { AxiosError, AxiosResponse } from "axios";
import { toast } from "vue-sonner";

import router from "@/router";

import { axios } from "@/configs/axiosConfig";
import { ACCESS_TOKEN } from "@/constants";
import type {
  ApiResponseData,
  ApiResponseError,
  AuthLoginRequest,
} from "@/generated/apiClient/data-contracts";
import { useUserStore } from "@/stores/user/user";

type UserLogout = Partial<
  MutationOptions<
    AxiosResponse<ApiResponseData<string>>,
    AxiosError<ApiResponseError>,
    AuthLoginRequest
  >
>;
export function useLogout(props?: UserLogout) {
  const userStore = useUserStore(); // Переместить сюда

  const { isPending, mutate } = useMutation({
    mutationFn: async (payload) => await axios.post("/api/v1/auth/logout", payload),
    onSuccess: async () => {
      localStorage.removeItem(ACCESS_TOKEN);
      userStore.$reset();

      toast.success("User logged out successfully. Redirecting to home page.");
      await router.push({ name: "home" });
    },
    onError: (error: AxiosError<ApiResponseError>) => {
      toast.error(error.response?.data.error?.message ?? "Something went wrong");
    },
    ...props,
  });

  return {
    isPending,
    mutate,
  };
}
