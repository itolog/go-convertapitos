<script setup lang="ts">
import { useForm } from "vee-validate";

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
const { isFieldDirty, handleSubmit } = useForm({
  validationSchema: formSchema,
});

const onSubmit = handleSubmit(() => {
  toast("Event has been created", {
    description: "Sunday, December 03, 2023 at 9:00 AM",
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

      <Button type="submit"> Submit </Button>
    </form>
  </div>
</template>

<style scoped></style>
