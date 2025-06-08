import { useQuery } from "@tanstack/vue-query";

import { axios } from "@/configs/axiosConfig.ts";

interface Payload {
  page: number;
  itemsPerPage: number;
}

export function useGetRolesOptions(payload: Payload) {
  const { isLoading, error, data } = useQuery({
    queryKey: ["users", payload.page, payload.itemsPerPage],
    queryFn: async () =>
      await axios
        .get(
          `api/v1/role/options?page=${payload.page}&limit=${payload.itemsPerPage}`,
        )
        .then((response) => response.data),
  });

  return {
    isLoading,
    error,
    data,
  };
}
