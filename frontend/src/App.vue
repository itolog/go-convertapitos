<script setup lang="ts">
import { RouterView } from 'vue-router'
import { SidebarProvider, SidebarTrigger } from '@/components/ui/sidebar'
import AppSidebar from '@/components/SideBar/AppSidebar.vue'
</script>

<template>
  <SidebarProvider>
    <AppSidebar />
    <SidebarTrigger />
    <main class="w-full">
      <RouterView />
    </main>
  </SidebarProvider>
</template>
