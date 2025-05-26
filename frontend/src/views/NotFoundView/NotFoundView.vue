<script setup lang="ts">
import { AlertCircle } from "lucide-vue-next";

import { Alert, AlertDescription, AlertTitle } from "@/components/ui/alert";
</script>

<template>
  <Alert variant="destructive">
    <AlertCircle class="w-4 h-4" />
    <AlertTitle>404</AlertTitle>
    <AlertDescription> Page Not Found </AlertDescription>
  </Alert>
</template>

<style scoped></style>
