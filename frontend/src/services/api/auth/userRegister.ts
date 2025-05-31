import {
  type MutationOptions,
  useMutation,
  useQueryClient,
} from "@tanstack/vue-query";
import type { AxiosError, AxiosResponse } from "axios";
import { toast } from "vue-sonner";

import { axios } from "@/configs/axiosConfig.ts";
import type {
  ApiResponseData,
  ApiResponseError,
  CommonAuthResponse,
  AuthRegisterRequest,
} from "@/generated/apiClient/data-contracts.ts";

type UserRegister = Partial<
  MutationOptions<
    AxiosResponse<ApiResponseData<CommonAuthResponse>>,
    AxiosError<ApiResponseError>,
    AuthRegisterRequest
  >
>;

export function userUserRegister(props?: UserRegister) {
  const queryClient = useQueryClient();

  const { isPending, mutateAsync, data } = useMutation({
    mutationFn: async (payload) =>
      await axios.post("/api/v1/auth/register", payload),
    onSuccess: async () => {
      await queryClient.invalidateQueries({ queryKey: ["users"] });

      toast.success("User registered successfully");
    },
    ...props,
  });

  return {
    isPending,
    mutateAsync,
    data,
  };
}
