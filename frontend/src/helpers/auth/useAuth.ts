import { useMutation } from "@tanstack/vue-query";
import type { AxiosError, AxiosResponse } from "axios";
import { onMounted } from "vue";

import { axios } from "@/configs/axiosConfig";
import { ACCESS_TOKEN } from "@/constants";
import type {
  ApiResponseData,
  ApiResponseError,
  CommonRefreshResponse,
} from "@/generated/apiClient/data-contracts";
import { isLogged } from "@/helpers";
import { useLogout } from "@/services/api/auth/useLogout.ts";

const fetcher = async () => await axios.post("api/v1/auth/refresh-token");

export function useAuth() {
  const { mutate, isPending: isLogOutPending } = useLogout();

  const { isPending, mutateAsync } = useMutation<
    AxiosResponse<ApiResponseData<CommonRefreshResponse>>,
    AxiosError<ApiResponseError>
  >({
    mutationFn: fetcher,
    onError: (e) => {
      const errorCode = e.response?.data.error?.code;
      if (errorCode === 401) {
        mutate();
      }
    },
    onSuccess: (response) => {
      const token = response?.data?.data?.accessToken;
      if (token) {
        localStorage.setItem(ACCESS_TOKEN, token);
      }
    },
  });

  onMounted(async () => {
    if (isLogged()) {
      await mutateAsync();
    }
  });

  return {
    isLoading: isPending || isLogOutPending,
  };
}
