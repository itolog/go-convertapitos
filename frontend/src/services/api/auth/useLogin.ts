import { type MutationOptions, useMutation } from "@tanstack/vue-query";
import type { AxiosError, AxiosResponse } from "axios";
import { toast } from "vue-sonner";

import { axios } from "@/configs/axiosConfig.ts";
import type {
  ApiResponseData,
  ApiResponseError,
  AuthLoginRequest,
  CommonAuthResponse,
} from "@/generated/apiClient/data-contracts.ts";

type UserLogin = Partial<
  MutationOptions<
    AxiosResponse<ApiResponseData<CommonAuthResponse>>,
    AxiosError<ApiResponseError>,
    AuthLoginRequest
  >
>;

export function useLogin(props?: UserLogin) {
  const { isPending, mutateAsync, data } = useMutation({
    mutationFn: async (payload) =>
      await axios.post("/api/v1/auth/login", payload),
    onSuccess: async () => {
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
