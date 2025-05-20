<script setup lang="ts">
import { RouterView } from 'vue-router'
import { SidebarProvider, SidebarTrigger } from '@/components/ui/sidebar'
import AppSidebar from '@/components/SideBar/AppSidebar.vue'
</script>

<template>
  <SidebarProvider>
    <AppSidebar />
    <SidebarTrigger />
    <main class="w-full p-4">
      <RouterView />
    </main>
  </SidebarProvider>
</template>
