<script setup lang="ts">
import { VueQueryDevtools } from "@tanstack/vue-query-devtools";
import { useCookies } from "@vueuse/integrations/useCookies";
import { RouterView } from "vue-router";
import "vue-sonner/style.css";

import AppBar from "@/components/AppBar/AppBar.vue";
import AppSidebar from "@/components/SideBar/AppSidebar.vue";
import AppSpinner from "@/components/common/AppSpinner/AppSpinner.vue";
import { SidebarProvider, SidebarInset } from "@/components/ui/sidebar";
import { SIDEBAR_COOKIE_NAME } from "@/components/ui/sidebar/utils.ts";
import { Toaster } from "@/components/ui/sonner";
import { useAuth } from "@/helpers/auth/useAuth";

const { isLoading } = useAuth();
const cookies = useCookies();
</script>

<template>
  <Toaster richColors position="top-center" />
  <SidebarProvider :default-open="cookies.get(SIDEBAR_COOKIE_NAME)">
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

  <VueQueryDevtools />
</template>
