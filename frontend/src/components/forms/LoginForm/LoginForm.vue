<script setup lang="ts">
import { ref } from 'vue'

const email = ref('')
const password = ref('')

const submit = () => {
  console.log(email.value, password.value)
}
</script>

<template>
  <div
    class="flex flex-col justify-center px-6 py-12 gap-4 lg:px-8 shadow-2xl rounded-2xl dark:shadow-orange-500"
  >
    <div class="sm:mx-auto sm:w-full sm:max-w-sm">
      <h2 class="mt-10 text-center text-2xl/9 font-bold tracking-tight">Sign in to your account</h2>
    </div>

    <div class="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
      <form novalidate class="flex flex-col gap-4" @submit.prevent="submit">
        <div>
          <label for="email" class="block text-sm/6 font-medium">Email address</label>
          <div class="mt-2">
            <input
              v-model="email"
              type="email"
              name="email"
              class="block w-full rounded-md bg-white px-3 py-1.5 text-base outline-1 -outline-offset-1 outline-gray-300 placeholder:text-gray-400 focus:outline-2 focus:-outline-offset-2 focus:outline-indigo-600 sm:text-sm/6"
            />
          </div>
        </div>

        <div>
          <div class="flex items-center justify-between">
            <label for="password" class="block text-sm/6 font-medium">Password</label>
          </div>
          <div class="mt-2">
            <input
              v-model="password"
              type="password"
              name="password"
              class="block w-full rounded-md bg-white px-3 py-1.5 text-base outline-1 -outline-offset-1 outline-gray-300 placeholder:text-gray-400 focus:outline-2 focus:-outline-offset-2 focus:outline-indigo-600 sm:text-sm/6"
            />
          </div>
        </div>

        <div>
          <button
            type="submit"
            class="flex w-full justify-center rounded-md bg-indigo-600 px-3 py-1.5 text-sm/6 font-semibold text-white shadow-xs hover:bg-indigo-500 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
          >
            Sign in
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<style scoped></style>
