import { useQuery } from "@tanstack/vue-query";

import { axios } from "@/configs/axiosConfig.ts";
import type {
  ApiResponseData,
  UserUser,
} from "@/generated/apiClient/data-contracts.ts";

interface GetUserProps {
  id: string;
}

export function useGetUser({ id }: GetUserProps) {
  const { isLoading, data, error } = useQuery({
    queryKey: ["user", id],
    queryFn: (): Promise<ApiResponseData<UserUser>> =>
      axios.get(`/api/v1/user/${id}`).then((response) => response.data),
    select: (data) => data.data,
    retry: false,
  });

  return {
    isLoading,
    data,
    error,
  };
}
