<script setup lang="ts">
import { EllipsisVerticalIcon, UserX, UserRoundPen } from "lucide-vue-next";
import type { PropType } from "vue";
import { useRouter } from "vue-router";

import { Button } from "@/components/ui/button";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import type { User } from "@/types/user";

const router = useRouter();

const { user } = defineProps({
  user: Object as PropType<User>,
});

const handleEdit = () => {
  router.push({ name: "editUser", params: { id: user?.id } });
};
</script>

<template>
  <DropdownMenu>
    <DropdownMenuTrigger as-child>
      <Button variant="ghost" class="w-8 h-8 p-0">
        <span class="sr-only">Open menu</span>
        <EllipsisVerticalIcon class="w-4 h-4" />
      </Button>
    </DropdownMenuTrigger>
    <DropdownMenuContent align="end">
      <DropdownMenuLabel class="truncate max-w-[220px]">
        {{ user?.email }}
      </DropdownMenuLabel>

      <DropdownMenuSeparator />
      <DropdownMenuItem @click="handleEdit" class="cursor-pointer">
        <UserRoundPen />
        Edit User
      </DropdownMenuItem>
      <DropdownMenuItem class="cursor-pointer">
        <UserX />
        Delete User
      </DropdownMenuItem>
    </DropdownMenuContent>
  </DropdownMenu>
</template>

<style scoped></style>
