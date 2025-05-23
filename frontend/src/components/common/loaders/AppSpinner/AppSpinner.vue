<template>
  <div class="flex h-full items-center justify-center">
    <div
      class="mx-auto w-[300px] h-[200px] bg-gray-950 dark:bg-gray-800 rounded-xl overflow-hidden drop-shadow-xl"
    >
      <div class="bg-[#333] flex items-center p-[5px] text-white relative">
        <div class="flex absolute left-3 gap-1">
          <span class="h-3.5 w-3.5 bg-[#ff605c] rounded-xl mr-2" />
          <span class="h-3.5 w-3.5 bg-[#ffbd44] rounded-xl mr-2" />
          <span class="h-3.5 w-3.5 bg-[#00ca4e] rounded-xl" />
        </div>
        <div class="flex-1 text-center text-2xl text-white">Status</div>
      </div>
      <div class="p-2.5 text-[#0f0]">
        <div>
          <span class="mr-2 text-2xl font-bold">Loading</span>
          <span class="animate-[ping_1.5s_0.5s_ease-in-out_infinite] text-3xl">.</span>
          <span class="animate-[ping_1.5s_0.7s_ease-in-out_infinite] text-3xl">.</span>
          <span class="animate-[ping_1.5s_0.9s_ease-in-out_infinite] text-3xl">.</span>
        </div>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts"></script>
