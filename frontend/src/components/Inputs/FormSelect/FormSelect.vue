<script setup lang="ts">
import type { SelectOption } from "@/components/Inputs/FormSelect/types.ts";
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

const { label, name, options } = defineProps({
  label: String,
  name: {
    type: String,
    required: true,
  },
  options: {
    type: Array<SelectOption>,
    required: true,
  },
});
</script>

<template>
  <FormField v-slot="{ componentField }" :name="name">
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
              v-for="option of options"
              :key="option.value"
              :value="option.value"
            >
              {{ option.label }}
            </SelectItem>
          </SelectGroup>
        </SelectContent>
      </Select>
    </FormItem>
  </FormField>
</template>

<style scoped></style>
