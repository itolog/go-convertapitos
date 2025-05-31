<script setup lang="ts">
import { toTypedSchema } from "@vee-validate/zod";
import { LucideMail, RectangleEllipsis } from "lucide-vue-next";
import { useForm } from "vee-validate";
import { useRouter } from "vue-router";
import { toast } from "vue-sonner";
import * as z from "zod";

import GoogleIcon from "@/assets/icons/google.svg";
import FormInput from "@/components/Inputs/FormInput/FormInput.vue";
import { Button } from "@/components/ui/button";
import type { ValidationErrorFields } from "@/generated/apiClient/data-contracts";
import { useLogin } from "@/services/api/auth/useLogin.ts";

const router = useRouter();

const formSchema = toTypedSchema(
  z.object({
    email: z.string().min(1),
    password: z.string().min(6).max(128),
  }),
);

const { isFieldDirty, handleSubmit, isSubmitting, setErrors } = useForm({
  validationSchema: formSchema,
});

const { mutateAsync, isPending } = useLogin({
  onError: (error) => {
    toast.error(error.response?.data.error?.message ?? "Something went wrong");
    const fieldsErrors = error.response?.data.error?.fields;
    if (fieldsErrors && fieldsErrors.length > 0) {
      const errors: Record<string, string> = {};

      fieldsErrors.forEach((fieldError: ValidationErrorFields) => {
        if (fieldError.field) {
          errors[fieldError.field.toLowerCase()] =
            `${fieldError.tag} ${fieldError.param}`;
        }
      });
      setErrors(errors);
    }
  },
});

const onSubmit = handleSubmit(async ({ email, password }) => {
  await mutateAsync({
    email,
    password,
  });

  await router.push({ name: "home" });
});
</script>

<template>
  <div
    class="flex flex-col w-full sm:w-sm md:w-md justify-center p-6 gap-4 lg:p-8 shadow-xl/50 rounded-2xl dark:shadow-orange-500"
  >
    <h2 class="text-center text-2xl/9 font-bold tracking-tight">
      Sign in to your account
    </h2>
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

      <Button :disabled="isPending || isSubmitting" type="submit"
        >Submit</Button
      >
    </form>

    <div class="flex flex-col gap-2">
      <h2 class="text-center font-bold">or login with</h2>
      <div>
        <a class="flex size-8" href="http://localhost:3000/api/v1/auth/google">
          <GoogleIcon />
        </a>
      </div>
    </div>
  </div>
</template>

<style scoped></style>
