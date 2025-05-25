<script setup lang="ts">
import { useQuery } from "@tanstack/vue-query";
import type { AxiosError, AxiosResponse } from "axios";

import UsersTable from "@/components/Tables/UsersTable/UsersTable.vue";
import TableSkeleton from "@/components/common/loaders/TableSkeleton/TableSkeleton.vue";
import { axios } from "@/configs/axiosConfig.ts";
import type {
  ApiResponseData,
  ApiResponseError,
} from "@/generated/apiClient/data-contracts.ts";
import type { User } from "@/types/user.ts";

const { isPending, isLoading, data } = useQuery<
  AxiosResponse<ApiResponseData<User[]>>,
  AxiosError<ApiResponseError>
>({
  queryKey: ["users"],
  queryFn: async () => await axios.get("api/v1/user"),
});
</script>

<template>
  <div class="h-full">
    <TableSkeleton v-if="isLoading" />
    <UsersTable v-else :loading="isPending" :data="data?.data?.data ?? []" />
  </div>
</template>

<style scoped></style>
