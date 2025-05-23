import { useQuery } from "@tanstack/vue-query";
import type { AxiosError, AxiosResponse } from "axios";
import { watchEffect } from "vue";

import { axios } from "@/configs/axiosConfig";
import { ACCESS_TOKEN } from "@/constants";
import type {
  ApiResponseData,
  ApiResponseError,
  CommonRefreshResponse,
} from "@/generated/apiClient/data-contracts";
import { isLogged } from "@/helpers";
import { useLogout } from "@/services/api/useLogout";

const fetcher = async () => await axios.get("api/v1/auth/refresh-token");

export function useAuth() {
  const { mutate, isPending } = useLogout();

  const { isLoading, data, error } = useQuery<
    AxiosResponse<ApiResponseData<CommonRefreshResponse>>,
    AxiosError<ApiResponseError>
  >({
    queryKey: ["auth"],
    queryFn: fetcher,
    retry: false,
    enabled: isLogged(),
  });

  const token = data.value?.data?.data?.accessToken;
  const errorData = error.value?.response?.data.error;

  watchEffect(() => {
    if (errorData?.code === 401) {
      mutate();
    }

    if (token) {
      localStorage.setItem(ACCESS_TOKEN, token);
    }
  });

  return {
    isLoading: isLoading || isPending,
  };
}
