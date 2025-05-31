import { useMutation, useQueryClient } from "@tanstack/vue-query";
import type { AxiosError } from "axios";
import { toast } from "vue-sonner";

import { axios } from "@/configs/axiosConfig.ts";
import type { ApiResponseError } from "@/generated/apiClient/data-contracts.ts";

export function useDeleteUser() {
  const queryClient = useQueryClient();

  const { isPending, mutateAsync } = useMutation({
    mutationFn: async (id: string) => await axios.delete(`/api/v1/user/${id}`),
    onSuccess: async () => {
      await queryClient.invalidateQueries({ queryKey: ["users"] });

      toast.success("User deleted successfully");
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
