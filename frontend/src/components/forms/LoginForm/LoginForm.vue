<script setup lang="ts">
import { useMutation } from "@tanstack/vue-query";
import { toTypedSchema } from "@vee-validate/zod";
import type { AxiosError, AxiosResponse } from "axios";
import { LucideMail, RectangleEllipsis } from "lucide-vue-next";
import { useForm } from "vee-validate";
import { toast } from "vue-sonner";
import * as z from "zod";

import router from "@/router";

import FormInput from "@/components/Inputs/FormInput/FormInput.vue";
import { Button } from "@/components/ui/button";
import { axios } from "@/configs/axiosConfig";
import { ACCESS_TOKEN } from "@/constants";
import type {
  ApiResponseData,
  ApiResponseError,
  AuthLoginRequest,
  CommonAuthResponse,
  ValidationErrorFields,
} from "@/generated/apiClient/data-contracts";

const formSchema = toTypedSchema(
  z.object({
    email: z.string().min(1),
    password: z.string().min(2).max(128),
  }),
);

const { isFieldDirty, handleSubmit, isSubmitting, setErrors } = useForm({
  validationSchema: formSchema,
});

const { isPending, mutate } = useMutation<
  AxiosResponse<ApiResponseData<CommonAuthResponse>>,
  AxiosError<ApiResponseError>,
  AuthLoginRequest
>({
  mutationFn: (payload) => axios.post("/api/v1/auth/login", payload),
  onSuccess: ({ data }) => {
    const token = data.data?.accessToken;
    if (token) {
      localStorage.setItem(ACCESS_TOKEN, token);
    }

    toast.success("User logged in successfully");
    router.push({ name: "home" });
  },
  onError: (error) => {
    toast.error(error.response?.data.error?.message ?? "Something went wrong");
    const fieldsErrors = error.response?.data.error?.fields;
    if (fieldsErrors && fieldsErrors.length > 0) {
      const errors: Record<string, string> = {};

      fieldsErrors.forEach((fieldError: ValidationErrorFields) => {
        if (fieldError.field) {
          errors[fieldError.field.toLowerCase()] = `${fieldError.tag} ${fieldError.param}`;
        }
      });
      setErrors(errors);
    }
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

      <Button :disabled="isPending || isSubmitting" type="submit">Submit</Button>
    </form>
  </div>
</template>

<style scoped></style>
