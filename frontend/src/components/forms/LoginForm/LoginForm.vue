<script setup lang="ts">
import { ref } from "vue";
import router from "@/router";
import { type Vueform } from "@vueform/vueform";
import type { AxiosResponse } from "axios";

import type { ApiResponseData } from "@/generated/apiClient/data-contracts.ts";
import { clearFormErrors, handleFormError } from "@/helpers/formHelpers.ts";

const schema = ref({
  email: { type: "text", label: "Email", rules: ["required", "email"] },
  password: { type: "text", label: "Password", rules: ["required"] },
  button: { type: "button", buttonLabel: "Login", submits: true },
  submit: { type: "submit", url: "" },
});

const handleSuccess = (response: AxiosResponse<ApiResponseData>, form$: Vueform) => {
  clearFormErrors(form$);
  // console.log(form$.formErrors);
  // console.log(response.data.data.user);

  form$.reset();
  router.push({ name: "home" });
};
</script>

<template>
  <div
    class="flex flex-col w-full sm:w-sm md:w-md justify-center p-6 gap-4 lg:p-8 shadow-2xl/50 rounded-2xl dark:shadow-orange-500"
  >
    <h2 class="text-center text-2xl/9 font-bold tracking-tight">Sign in to your account</h2>

    <Vueform
      endpoint="/api/v1/auth/login"
      method="post"
      @error="handleFormError"
      @success="handleSuccess"
      :schema="schema"
    />
  </div>
</template>

<style scoped></style>
