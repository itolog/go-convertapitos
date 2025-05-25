<script setup lang="ts">
import {
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { cn } from "@/lib/utils";

defineProps({
  isFieldDirty: Function,
  name: {
    type: String,
    required: true,
  },
  type: {
    type: String,
    default: "text",
  },
  placeholder: {
    type: String,
    default: "",
  },
  label: {
    type: String,
    default: "",
  },
});
</script>

<template>
  <FormField
    v-slot="{ componentField }"
    :name="name"
    :validate-on-blur="!isFieldDirty"
  >
    <FormItem class="relative">
      <FormLabel class="capitalize">{{ label ? label : name }}</FormLabel>
      <div class="relative w-full items-center">
        <FormControl>
          <Input
            :class="cn({ 'pl-10': $slots.icon })"
            :type="type"
            :placeholder="placeholder"
            v-bind="componentField"
          />
        </FormControl>
        <span
          v-if="$slots.icon"
          class="absolute start-0 inset-y-0 flex items-center justify-center px-2"
        >
          <slot class="size-6 text-muted-foreground" name="icon" />
        </span>
      </div>

      <FormMessage class="absolute bottom-[-20px]" />
    </FormItem>
  </FormField>
</template>

<style scoped></style>
