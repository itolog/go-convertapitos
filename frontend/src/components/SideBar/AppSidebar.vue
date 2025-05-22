<script setup lang="ts">
import NavMain from "@/components/SideBar/NavMain.vue";
import NavUser from "@/components/SideBar/NavUser.vue";
import ThemeSwitcher from "@/components/ThemeSwitcher/ThemeSwitcher.vue";
import {
  Sidebar,
  SidebarContent,
  SidebarHeader,
  SidebarRail,
  type SidebarProps,
} from "@/components/ui/sidebar";

const props = withDefaults(defineProps<SidebarProps>(), {
  collapsible: "icon",
  variant: "floating",
});
</script>

<template>
  <Sidebar v-bind="props">
    <SidebarHeader>
      <NavUser />
    </SidebarHeader>
    <SidebarContent>
      <ThemeSwitcher />
      <NavMain />
    </SidebarContent>
    <SidebarRail />
  </Sidebar>
</template>
