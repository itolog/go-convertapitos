<script setup lang="ts">
import { useForm } from "vee-validate";
import { useMutation } from "@tanstack/vue-query";
import { RectangleEllipsis, LucideMail } from "lucide-vue-next";

import FormInput from "@/components/Inputs/FormInput/FormInput.vue";

import { toTypedSchema } from "@vee-validate/zod";
import * as z from "zod";

import { Button } from "@/components/ui/button";

const formSchema = toTypedSchema(
  z.object({
    email: z.string().email().min(1),
    password: z.string().min(6).max(128),
  }),
);

import { toast } from "vue-sonner";
import { axios } from "@/configs/axiosConfig.ts";
import type {
  AuthLoginRequest,
  CommonAuthResponse,
  ApiResponseError,
  ApiResponseData,
} from "@/generated/apiClient/data-contracts.ts";
import type { AxiosError, AxiosResponse } from "axios";

const { isFieldDirty, handleSubmit } = useForm({
  validationSchema: formSchema,
});

const { isPending, mutate } = useMutation<
  AxiosResponse<ApiResponseData<CommonAuthResponse>>,
  AxiosError<ApiResponseError>,
  AuthLoginRequest
>({
  mutationFn: (payload) => axios.post("/api/v1/auth/login", payload),
  onSuccess: () => {
    toast.success("User logged in successfully");
  },
  onError: (error) => {
    toast.error(error.response?.data.error?.message ?? "Something went wrong");
  },
});

const onSubmit = handleSubmit(({ email, password }) => {
  mutate({
    email,
    password,
  });
});
</script>

<template>
  <div
    class="flex flex-col w-full sm:w-sm md:w-md justify-center p-6 gap-4 lg:p-8 shadow-2xl/50 rounded-2xl dark:shadow-orange-500"
  >
    <span v-if="isPending">Loading...</span>
    <h2 class="text-center text-2xl/9 font-bold tracking-tight">Sign in to your account</h2>
    <form class="flex flex-col gap-8" @submit="onSubmit">
      <div class="flex flex-col gap-6">
        <FormInput name="email" :is-field-dirty="isFieldDirty">
          <template v-slot:icon>
            <LucideMail />
          </template>
        </FormInput>
        <FormInput name="password" :is-field-dirty="isFieldDirty">
          <template v-slot:icon>
            <RectangleEllipsis />
          </template>
        </FormInput>
      </div>

      <Button :disabled="isPending" type="submit"> Submit </Button>
    </form>
  </div>
</template>

<style scoped></style>
