<script setup lang="ts">
import { LucideMail, UserPen } from "lucide-vue-next";
import { useForm } from "vee-validate";
import { watchEffect } from "vue";
import { useRouter } from "vue-router";
import { toast } from "vue-sonner";

import FormInput from "@/components/Inputs/FormInput/FormInput.vue";
import FormSelect from "@/components/forms/EditUser/components/FormSelect/FormSelect.vue";
import type { SelectOption } from "@/components/forms/EditUser/components/FormSelect/types.ts";
import { formSchema } from "@/components/forms/EditUser/formSchema";
import { Button } from "@/components/ui/button";
import { editUser } from "@/services/api/users/editUser.ts";
import { useGetUser } from "@/services/api/users/useGetUser";

const router = useRouter();

const selectOptions: SelectOption[] = [
  {
    label: "True",
    value: true,
  },
  {
    label: "False",
    value: false,
  },
];

const { id } = defineProps({
  id: {
    type: String,
    required: true,
  },
});

const { data, error, isLoading: getUserLoading } = useGetUser({ id });

const { mutateAsync, isPending } = editUser();

const { isFieldDirty, handleSubmit, isSubmitting, setValues } = useForm({
  validationSchema: formSchema,
  initialValues: {
    name: "",
    email: "",
    verifiedEmail: false,
  },
});

watchEffect(() => {
  if (error.value?.message) {
    toast.error(error.value.message);
  }

  if (data.value) {
    setValues({
      name: data.value.name,
      email: data.value.email,
      verifiedEmail: data.value.verifiedEmail,
    });
  }
});

const onSubmit = handleSubmit(async (values) => {
  await mutateAsync({
    id,
    data: values,
  });
  await router.push({ name: "users" });
});
</script>

<template>
  <div class="form-container">
    <h2 class="form-title">Edit user</h2>

    <form class="flex flex-col gap-8" @submit="onSubmit">
      <div class="flex flex-col gap-6">
        <FormInput name="name" :is-field-dirty="isFieldDirty">
          <template v-slot:icon>
            <UserPen />
          </template>
        </FormInput>
        <FormInput name="email" :is-field-dirty="isFieldDirty">
          <template v-slot:icon>
            <LucideMail />
          </template>
        </FormInput>

        <FormSelect
          label="Verified Email"
          :options="selectOptions"
          name="verifiedEmail"
        />
      </div>

      <Button
        :disabled="getUserLoading || isPending || isSubmitting"
        type="submit"
      >
        Submit
      </Button>
    </form>
  </div>
</template>

<style scoped></style>
