<script setup lang="ts">
import { useCookies } from "@vueuse/integrations/useCookies";
import { RouterView } from "vue-router";

import AppBar from "@/components/AppBar/AppBar.vue";
import AppSidebar from "@/components/SideBar/AppSidebar.vue";
import { SidebarInset, SidebarProvider } from "@/components/ui/sidebar";
import { SIDEBAR_COOKIE_NAME } from "@/components/ui/sidebar/utils.ts";
import { useAuth } from "@/helpers/auth/useAuth.ts";

useAuth();
const cookies = useCookies();
</script>

<template>
  <SidebarProvider :default-open="cookies.get(SIDEBAR_COOKIE_NAME)">
    <AppSidebar />
    <SidebarInset class="flex overflow-auto">
      <AppBar />

      <div class="h-full p-4">
        <RouterView />
      </div>
    </SidebarInset>
  </SidebarProvider>
</template>

<style scoped></style>
