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

const fetcher = async () => await axios.get("api/v1/auth/refresh-token");

export function useAuth() {
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

  watchEffect(() => {
    if (error) {
      //console.log(error.value?.response?.data.error);
      // TODO: logout user
    }
    if (token) {
      localStorage.setItem(ACCESS_TOKEN, token);
    }
  });

  return {
    isLoading,
  };
}
