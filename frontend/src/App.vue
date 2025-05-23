<script setup lang="ts">
import { VueQueryDevtools } from "@tanstack/vue-query-devtools";
import "vue-sonner/style.css";

import { Toaster } from "@/components/ui/sonner";
import MainLayout from "@/layouts/MainLayout.vue";
</script>

<template>
  <Toaster richColors position="top-center" />
  <MainLayout />
  <VueQueryDevtools />
</template>
