<script setup lang="ts">
import { LucideMail, RectangleEllipsis, UserPen } from "lucide-vue-next";
import { useForm } from "vee-validate";
import { toast } from "vue-sonner";

import router from "@/router";

import FormInput from "@/components/Inputs/FormInput/FormInput.vue";
import formSchema from "@/components/forms/RegisterUserForm/schema";
import { Button } from "@/components/ui/button";
import type { ValidationErrorFields } from "@/generated/apiClient/data-contracts.ts";
import { userUserRegister } from "@/services/api/userRegister.ts";

const { isFieldDirty, handleSubmit, isSubmitting, setErrors } = useForm({
  validationSchema: formSchema,
});

const { mutateAsync, isPending } = userUserRegister({
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

const onSubmit = handleSubmit(async ({ email, password, name }) => {
  await mutateAsync({
    email,
    password,
    name,
  });

  await router.push({ name: "users" });
});
</script>

<template>
  <div
    class="flex flex-col w-full sm:w-sm md:w-md justify-center p-6 gap-4 lg:p-8 shadow-xl/50 rounded-2xl dark:shadow-orange-500"
  >
    <h2 class="text-center text-2xl/9 font-bold tracking-tight">
      User Registration
    </h2>
    <form class="flex flex-col gap-8" @submit="onSubmit">
      <div class="grid grid-cols-1 gap-6">
        <FormInput class="w-full" name="name" :is-field-dirty="isFieldDirty">
          <template v-slot:icon>
            <UserPen />
          </template>
        </FormInput>
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
  </div>
</template>

<style scoped></style>
