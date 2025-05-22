<script setup lang="ts">
import { RouterView } from "vue-router";
import "vue-sonner/style.css";

import AppBar from "@/components/AppBar/AppBar.vue";
import AppSidebar from "@/components/SideBar/AppSidebar.vue";
import AppSpinner from "@/components/common/AppSpinner/AppSpinner.vue";
import { SidebarProvider, SidebarInset } from "@/components/ui/sidebar";
import { Toaster } from "@/components/ui/sonner";
import { useAuth } from "@/helpers/auth/useAuth";

const { isLoading } = useAuth();
</script>

<template>
  <Toaster richColors position="top-center" />
  <SidebarProvider>
    <AppSidebar />
    <SidebarInset>
      <AppBar />
      <main class="h-full p-2">
        <RouterView v-slot="{ Component }">
          <transition>
            <AppSpinner v-if="isLoading" />
            <component v-else :is="Component" />
          </transition>
        </RouterView>
      </main>
    </SidebarInset>
  </SidebarProvider>
</template>
