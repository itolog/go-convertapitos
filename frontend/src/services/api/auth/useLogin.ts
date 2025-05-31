import { type MutationOptions, useMutation } from "@tanstack/vue-query";
import type { AxiosError, AxiosResponse } from "axios";
import { toast } from "vue-sonner";

import { axios } from "@/configs/axiosConfig.ts";
import { ACCESS_TOKEN } from "@/constants";
import type {
  ApiResponseData,
  ApiResponseError,
  AuthLoginRequest,
  CommonAuthResponse,
} from "@/generated/apiClient/data-contracts.ts";
import { useUserStore } from "@/stores/user/user.ts";

type UserLogin = Partial<
  MutationOptions<
    AxiosResponse<ApiResponseData<CommonAuthResponse>>,
    AxiosError<ApiResponseError>,
    AuthLoginRequest
  >
>;

export function useLogin(props?: UserLogin) {
  const userStore = useUserStore();

  const { isPending, mutateAsync, data } = useMutation({
    mutationFn: async (payload) =>
      await axios.post("/api/v1/auth/login", payload),
    onSuccess: async ({ data }) => {
      const token = data.data?.accessToken;
      if (token) {
        localStorage.setItem(ACCESS_TOKEN, token);
      }

      userStore.$patch({
        user: data.data?.user,
        isLoggedIn: true,
      });

      toast.success("User logged in successfully");
    },
    ...props,
  });

  return {
    isPending,
    mutateAsync,
    data,
  };
}
