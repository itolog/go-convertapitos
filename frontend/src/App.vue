<script setup lang="ts">
import { RouterView } from "vue-router";
import { SidebarProvider, SidebarInset } from "@/components/ui/sidebar";
import AppSidebar from "@/components/SideBar/AppSidebar.vue";
import AppBar from "@/components/AppBar/AppBar.vue";
import { Toaster } from "@/components/ui/sonner";
import "vue-sonner/style.css";
</script>

<template>
  <Toaster richColors position="top-center" />

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
