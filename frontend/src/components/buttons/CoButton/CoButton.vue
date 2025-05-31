<script setup lang="ts">
import type { CoButtonProps } from "@/components/buttons/CoButton/types.ts";
import { Button } from "@/components/ui/button";

withDefaults(defineProps<CoButtonProps>(), {});
</script>

<template>
  <Button v-bind="$props">
    <div v-if="$slots.icon" class="icon">
      <slot name="icon" />
    </div>

    <div v-if="$slots.default" class="card-content">
      <slot />
    </div>
  </Button>
</template>

<style scoped></style>
