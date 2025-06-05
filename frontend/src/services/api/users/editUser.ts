import { useMutation, useQueryClient } from "@tanstack/vue-query";
import type { AxiosError, AxiosResponse } from "axios";
import { toast } from "vue-sonner";

import { axios } from "@/configs/axiosConfig.ts";
import type {
  ApiResponseData,
  ApiResponseError,
  UserUpdateRequest,
} from "@/generated/apiClient/data-contracts.ts";

interface Payload {
  id: string;
  data: UserUpdateRequest;
}

export function editUser() {
  const queryClient = useQueryClient();

  const { isPending, mutateAsync } = useMutation<
    AxiosResponse<ApiResponseData<string>>,
    AxiosError<ApiResponseError>,
    Payload
  >({
    mutationFn: async (payload: Payload) =>
      await axios.patch(`/api/v1/user/${payload.id}`, payload.data),
    onSuccess: async () => {
      await queryClient.invalidateQueries({ queryKey: ["users"] });

      toast.success("User updated successfully");
    },
    onError: (error: AxiosError<ApiResponseError>) => {
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
