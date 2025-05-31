import { useMutation, useQueryClient } from "@tanstack/vue-query";
import type { AxiosError, AxiosResponse } from "axios";
import { toast } from "vue-sonner";

import { axios } from "@/configs/axiosConfig.ts";
import type {
  ApiResponseData,
  ApiResponseError,
  UserBatchDeleteRequest,
} from "@/generated/apiClient/data-contracts.ts";

export function useBatchDeleteUsers() {
  const queryClient = useQueryClient();

  const { isPending, mutateAsync } = useMutation<
    AxiosResponse<ApiResponseData<string>>,
    AxiosError<ApiResponseError>,
    UserBatchDeleteRequest
  >({
    mutationFn: async (payload) =>
      await axios.delete(`/api/v1/user/by_ids`, {
        data: payload,
      }),
    onSuccess: async () => {
      await queryClient.invalidateQueries({ queryKey: ["users"] });

      toast.success("Users deleted successfully");
    },
    onError: (error) => {
      toast.error(
        error.response?.data.error?.message ?? "Something went wrong",
      );
    },
  });

  return {
    isPending,
    mutateAsync,
  };
}
