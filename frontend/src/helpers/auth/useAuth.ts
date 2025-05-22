import { useQuery } from "@tanstack/vue-query";
import type { AxiosError, AxiosResponse } from "axios";
import { watch } from "vue";

import { axios } from "@/configs/axiosConfig.ts";
import { ACCESS_TOKEN, IS_LOGGED_STORAGE_KEY, USER_STORAGE_KEY } from "@/constants";
import type {
  ApiResponseData,
  ApiResponseError,
  CommonAuthResponse,
} from "@/generated/apiClient/data-contracts.ts";

const fetcher = async () => await axios.get("api/v1/auth/refresh-token");

export function useAuth() {
  const { isLoading, data } = useQuery<
    AxiosResponse<ApiResponseData<CommonAuthResponse>>,
    AxiosError<ApiResponseError>
  >({
    queryKey: ["auth"],
    queryFn: fetcher,
    retry: false,
  });

  watch(
    () => data,
    (data) => {
      const token = data.value?.data?.data?.accessToken;
      const user = data.value?.data?.data?.user;

      if (token) {
        localStorage.setItem(ACCESS_TOKEN, token);
      }
      localStorage.setItem(USER_STORAGE_KEY, JSON.stringify(user));
      localStorage.setItem(IS_LOGGED_STORAGE_KEY, "true");
    },
    { deep: true },
  );

  return {
    isLoading,
  };
}
