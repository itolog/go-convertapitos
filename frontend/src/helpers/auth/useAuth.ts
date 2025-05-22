import { useQuery } from "@tanstack/vue-query";
import type { AxiosError, AxiosResponse } from "axios";
import { watch } from "vue";

import { axios } from "@/configs/axiosConfig.ts";
import { ACCESS_TOKEN } from "@/constants";
import type {
  ApiResponseData,
  ApiResponseError,
  CommonRefreshResponse,
} from "@/generated/apiClient/data-contracts.ts";

const fetcher = async () => await axios.get("api/v1/auth/refresh-token");

export function useAuth() {
  const { isLoading, data } = useQuery<
    AxiosResponse<ApiResponseData<CommonRefreshResponse>>,
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

      if (token) {
        localStorage.setItem(ACCESS_TOKEN, token);
      }
    },
    { deep: true },
  );

  return {
    isLoading,
  };
}
