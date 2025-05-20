<script setup lang="ts">
import { ref } from 'vue'

const schema = ref({
  email: { type: 'text', label: 'Email', rules: ['required', 'email'] },
  password: { type: 'text', label: 'Password', rules: ['required'] },
  button: { type: 'button', buttonLabel: 'Login', submits: true },
})
</script>

<template>
  <div
    class="flex flex-col justify-center px-6 py-12 gap-4 lg:px-8 shadow-2xl rounded-2xl dark:shadow-orange-500"
  >
    <div class="sm:mx-auto sm:w-full sm:max-w-sm">
      <h2 class="mt-10 text-center text-2xl/9 font-bold tracking-tight">Sign in to your account</h2>
    </div>

    <div class="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
      <Vueform endpoint="http://localhost:3000/api/v1/auth/login" method="post" :schema="schema" />
    </div>
  </div>
</template>

<style scoped></style>
