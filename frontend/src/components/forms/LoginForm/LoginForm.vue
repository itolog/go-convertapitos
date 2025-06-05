<script setup lang="ts">
import { toTypedSchema } from "@vee-validate/zod";
import { LucideMail, RectangleEllipsis } from "lucide-vue-next";
import { useForm } from "vee-validate";
import { useRouter } from "vue-router";
import { toast } from "vue-sonner";
import * as z from "zod";

import FormInput from "@/components/Inputs/FormInput/FormInput.vue";
import { Button } from "@/components/ui/button";
import type { ValidationErrorFields } from "@/generated/apiClient/data-contracts";
import { useLogin } from "@/services/api/auth/useLogin.ts";
import { socialLinks } from "@/views/Auth/data";

const baseUrl = import.meta.env.VITE_API_URL;
const originUrl = window.location.origin;
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
  <div class="form-container">
    <h2 class="form-title">Sign in to your account</h2>
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

      <Button :disabled="isPending || isSubmitting" type="submit">
        Submit
      </Button>
    </form>

    <div class="flex flex-col gap-2">
      <h2 class="text-center font-bold">or login with</h2>
      <div class="flex gap-2">
        <a
          v-for="link of socialLinks"
          class="flex size-8"
          :key="link.name"
          :href="`${baseUrl}/api/v1/auth/${link.name}/?redirect_to=${originUrl}/users`"
        >
          <component :is="link.icon" />
        </a>
      </div>
    </div>
  </div>
</template>

<style scoped></style>
