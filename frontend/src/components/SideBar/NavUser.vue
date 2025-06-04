<script setup lang="ts">
import {
  BadgeCheck,
  Bell,
  ChevronsUpDown,
  CreditCard,
  LogIn,
  LogOut,
} from "lucide-vue-next";

import UserAvatar from "@/components/common/ui/UserAvatar/UserAvatar.vue";
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
import { useLogout } from "@/services/api/auth/useLogout.ts";
import { useUserStore } from "@/stores/user/user";

const userStore = useUserStore();

const { mutate, isPending } = useLogout();
const { isMobile } = useSidebar();
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
            <UserAvatar :user="userStore.user" />

            <div class="grid flex-1 text-left text-sm leading-tight">
              <span class="truncate font-semibold">{{
                userStore.user?.name
              }}</span>
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
              <UserAvatar :user="userStore.user" />
              <div class="grid flex-1 text-left text-sm leading-tight">
                <span class="truncate font-semibold">{{
                  userStore.user?.name
                }}</span>
                <span class="truncate text-xs">{{
                  userStore.user?.email
                }}</span>
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
              @click="mutate"
              variant="ghost"
              size="sm"
              :disabled="isPending"
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
