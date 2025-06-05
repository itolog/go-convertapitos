<script setup lang="ts">
import { useRoute } from "vue-router";

import EditUserForm from "@/components/forms/EditUser/EditUserForm.vue";

const route = useRoute();
</script>

<template>
  <div class="flex items-center justify-center h-full">
    <EditUserForm :id="route.params.id.toString()" />
  </div>
</template>

<style scoped></style>
