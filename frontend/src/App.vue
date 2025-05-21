<script setup lang="ts">
import { RouterView } from 'vue-router'
import { SidebarProvider, SidebarInset } from '@/components/ui/sidebar'
import AppSidebar from '@/components/SideBar/AppSidebar.vue'
import AppBar from '@/components/AppBar/AppBar.vue'
</script>

<template>
  <SidebarProvider>
    <AppSidebar />

    <SidebarInset>
      <AppBar />
      <main class="h-full p-2">
        <RouterView />
      </main>
    </SidebarInset>
  </SidebarProvider>
</template>
