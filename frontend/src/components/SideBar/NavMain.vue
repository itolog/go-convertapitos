<script setup lang="ts">
import { ChevronRight, Home, Users } from "lucide-vue-next";
import { useRoute } from "vue-router";

import { Collapsible, CollapsibleContent, CollapsibleTrigger } from "@/components/ui/collapsible";
import {
  SidebarGroup,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
  SidebarMenuSub,
  SidebarMenuSubButton,
  SidebarMenuSubItem,
} from "@/components/ui/sidebar";
import type { Navigation } from "@/types/navigation.ts";

const navItems: Navigation = [
  {
    title: "Home",
    url: "/",
    icon: Home,
  },
  {
    title: "Users",
    url: "/users",
    icon: Users,
  },
];

const route = useRoute();
</script>

<template>
  <SidebarGroup>
    <SidebarMenu>
      <Collapsible
        v-for="item in navItems"
        :key="item.title"
        as-child
        :default-open="item.isActive"
        class="group/collapsible"
      >
        <SidebarMenuItem>
          <CollapsibleTrigger v-if="item.items?.length" as-child>
            <SidebarMenuButton :tooltip="item.title">
              <component :is="item.icon" v-if="item.icon" />
              <span>{{ item.title }}</span>
              <ChevronRight
                class="ml-auto transition-transform duration-200 group-data-[state=open]/collapsible:rotate-90"
              />
            </SidebarMenuButton>
          </CollapsibleTrigger>
          <SidebarMenuButton
            :is-active="route.path === item.url"
            class="cursor-pointer"
            v-else
            :tooltip="item.title"
          >
            <component :is="item.icon" v-if="item.icon" />
            <RouterLink :to="item.url">{{ item.title }} </RouterLink>
          </SidebarMenuButton>

          <CollapsibleContent v-if="item.items?.length">
            <SidebarMenuSub>
              <SidebarMenuSubItem v-for="subItem in item.items" :key="subItem.title">
                <SidebarMenuSubButton
                  :is-active="route.path === item.url"
                  class="cursor-pointer"
                  as-child
                >
                  <RouterLink :to="subItem.url">
                    <span>{{ subItem.title }}</span>
                  </RouterLink>
                </SidebarMenuSubButton>
              </SidebarMenuSubItem>
            </SidebarMenuSub>
          </CollapsibleContent>
        </SidebarMenuItem>
      </Collapsible>
    </SidebarMenu>
  </SidebarGroup>
</template>
