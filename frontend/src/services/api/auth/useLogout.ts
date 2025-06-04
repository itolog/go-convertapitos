import { type MutationOptions, useMutation } from "@tanstack/vue-query";
import type { AxiosError, AxiosResponse } from "axios";
import { toast } from "vue-sonner";

import router from "@/router";

import { axios } from "@/configs/axiosConfig.ts";
import type {
  ApiResponseData,
  ApiResponseError,
} from "@/generated/apiClient/data-contracts.ts";
import { useUserStore } from "@/stores/user/user.ts";

type UserLogout = Partial<
  MutationOptions<
    AxiosResponse<ApiResponseData<string>>,
    AxiosError<ApiResponseError>
  >
>;
export function useLogout(props?: UserLogout) {
  const userStore = useUserStore();

  const { isPending, mutate } = useMutation({
    mutationFn: async (payload) =>
      await axios.post("/api/v1/auth/logout", payload),
    onSuccess: async () => {
      userStore.$reset();

      toast.success("User logged out successfully. Redirecting to home page.");
      await router.push({ name: "home" });
    },
    onError: (error: AxiosError<ApiResponseError>) => {
      toast.error(
        error.response?.data.error?.message ?? "Something went wrong",
      );
    },
    ...props,
  });

  return {
    isPending,
    mutate,
  };
}
