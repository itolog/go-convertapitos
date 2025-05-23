<script setup lang="ts">
import { useQuery } from "@tanstack/vue-query";
import type { AxiosError, AxiosResponse } from "axios";
import { watchEffect } from "vue";

import AppTable from "@/components/AppTable/AppTable.vue";
import { axios } from "@/configs/axiosConfig";
import type { ApiResponseData, ApiResponseError } from "@/generated/apiClient/data-contracts";
import type { User } from "@/types/user";

// const { isLoading, data } = useQuery<
//   AxiosResponse<ApiResponseData<User[]>>,
//   AxiosError<ApiResponseError>
// >({
//   queryKey: ["users"],
//   queryFn: async () => await axios.get("api/v1/user"),
//   retry: false,
// });
//
// watchEffect(() => {
//   console.log(data.value?.data.data);
// });
</script>

<template>
  <div>
    <!--    <span v-if="isLoading">LOADING</span>-->
    <!--    <AppTable :data="data?.data?.data ?? []" />-->
  </div>
</template>

<style scoped></style>
