<script setup lang="ts">
import { watchEffect } from "vue";
import { toast } from "vue-sonner";

import {
  FormControl,
  FormField,
  FormItem,
  FormLabel,
} from "@/components/ui/form";
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { useGetRolesOptions } from "@/services/api/roles/useGetRolesOptions.ts";

const { label, name } = defineProps({
  label: String,
  name: {
    type: String,
    required: true,
  },
});

const { data: roleData, error: roleError } = useGetRolesOptions({
  page: 1,
  itemsPerPage: 25,
});

watchEffect(() => {
  if (roleError.value?.message) {
    toast.error(roleError.value.message);
  }
});
</script>

<template>
  <FormField v-slot="{ componentField }" :name="name">
    {{ roleData?.data?.data }}
    <FormItem>
      <FormLabel v-if="label">{{ label }}</FormLabel>
      <Select v-bind="componentField">
        <FormControl>
          <SelectTrigger class="w-full">
            <SelectValue />
          </SelectTrigger>
        </FormControl>
        <SelectContent>
          <SelectGroup>
            <SelectItem
              v-for="option of roleData?.data"
              :key="option.id"
              :value="option.id"
            >
              {{ option.name }}
            </SelectItem>
          </SelectGroup>
        </SelectContent>
      </Select>
    </FormItem>
  </FormField>
</template>

<style scoped></style>
