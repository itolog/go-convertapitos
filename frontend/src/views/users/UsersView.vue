<script setup lang="ts">
import { useQuery } from "@tanstack/vue-query";
import type { AxiosError, AxiosResponse } from "axios";
import { storeToRefs } from "pinia";
import { onUnmounted } from "vue";

import UsersTable from "@/components/Tables/UsersTable/UsersTable.vue";
import { axios } from "@/configs/axiosConfig.ts";
import type {
  ApiResponseData,
  ApiResponseError,
} from "@/generated/apiClient/data-contracts";
import { useTableStore } from "@/stores/table/table";
import type { User } from "@/types/user";

const tableStore = useTableStore();

const { page, itemsPerPage } = storeToRefs(tableStore);

const { isLoading, isFetching, data } = useQuery<
  AxiosResponse<ApiResponseData<User[]>>,
  AxiosError<ApiResponseError>
>({
  queryKey: ["users", page, itemsPerPage],
  queryFn: async () =>
    await axios.get(
      `api/v1/user?page=${page.value}&limit=${itemsPerPage.value}`,
    ),
});

onUnmounted(() => {
  tableStore.setPage(1);
});
</script>

<template>
  <div class="h-full">
    <UsersTable
      :isLoading="isLoading"
      :isFetching="isFetching"
      :meta="data?.data?.meta ?? {}"
      :data="data?.data?.data ?? []"
    />
  </div>
</template>

<style scoped></style>
