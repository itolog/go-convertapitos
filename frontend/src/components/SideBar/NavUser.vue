<script setup lang="ts">
import { BadgeCheck, Bell, ChevronsUpDown, CreditCard, LogIn, LogOut } from "lucide-vue-next";

import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { Button } from "@/components/ui/button";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuGroup,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import {
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
  useSidebar,
} from "@/components/ui/sidebar";
import { ACCESS_TOKEN } from "@/constants";
import { useUserStore } from "@/types/user.ts";

defineProps<{
  user: {
    name: string;
    email: string;
    avatar: string;
  };
}>();

const { isMobile } = useSidebar();
const userStore = useUserStore();

const handleLogOut = () => {
  //TODO: call logout api
  localStorage.removeItem(ACCESS_TOKEN);
  userStore.setUser(null);
  userStore.setIsLogged(false);
};
</script>

<template>
  <SidebarMenu>
    <SidebarMenuItem>
      <DropdownMenu>
        <DropdownMenuTrigger as-child>
          <SidebarMenuButton
            size="lg"
            class="cursor-pointer data-[state=open]:bg-sidebar-accent data-[state=open]:text-sidebar-accent-foreground"
          >
            <Avatar class="h-8 w-8 rounded-lg">
              <AvatarImage :src="userStore.user?.picture ?? ''" :alt="userStore.user?.name ?? ''" />
              <AvatarFallback class="rounded-lg"> AN </AvatarFallback>
            </Avatar>
            <div class="grid flex-1 text-left text-sm leading-tight">
              <span class="truncate font-semibold">{{ userStore.user?.name }}</span>
              <span class="truncate text-xs">{{ userStore.user?.email }}</span>
            </div>
            <ChevronsUpDown class="ml-auto size-4" />
          </SidebarMenuButton>
        </DropdownMenuTrigger>
        <DropdownMenuContent
          class="w-[--reka-dropdown-menu-trigger-width] min-w-56 rounded-lg"
          :side="isMobile ? 'bottom' : 'right'"
          align="end"
          :side-offset="4"
        >
          <DropdownMenuLabel class="p-0 font-normal">
            <div class="flex items-center gap-2 px-1 py-1.5 text-left text-sm">
              <Avatar class="h-8 w-8 rounded-lg">
                <AvatarImage :src="userStore.user?.picture ?? ''" :alt="userStore.user?.name" />
                <AvatarFallback class="rounded-lg">AN</AvatarFallback>
              </Avatar>
              <div class="grid flex-1 text-left text-sm leading-tight">
                <span class="truncate font-semibold">{{ userStore.user?.name }}</span>
                <span class="truncate text-xs">{{ userStore.user?.email }}</span>
              </div>
            </div>
          </DropdownMenuLabel>
          <DropdownMenuSeparator />
          <DropdownMenuGroup>
            <DropdownMenuItem>
              <BadgeCheck />
              Account
            </DropdownMenuItem>
            <DropdownMenuItem>
              <CreditCard />
              Billing
            </DropdownMenuItem>
            <DropdownMenuItem>
              <Bell />
              Notifications
            </DropdownMenuItem>
          </DropdownMenuGroup>
          <DropdownMenuSeparator />
          <DropdownMenuItem v-if="!userStore.isLoggedIn">
            <RouterLink class="flex gap-2 items-center w-full" to="/login">
              <LogIn />
              Log In
            </RouterLink>
          </DropdownMenuItem>
          <DropdownMenuItem class="pl-0" v-else>
            <Button
              @click="handleLogOut"
              variant="ghost"
              size="sm"
              class="w-full flex justify-start cursor-pointer"
            >
              <LogOut />
              Log Out
            </Button>
          </DropdownMenuItem>
        </DropdownMenuContent>
      </DropdownMenu>
    </SidebarMenuItem>
  </SidebarMenu>
</template>
