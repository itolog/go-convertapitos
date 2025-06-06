<script setup lang="ts">
import { EllipsisVerticalIcon, UserRoundPen } from "lucide-vue-next";
import { type PropType, ref } from "vue";
import { useRouter } from "vue-router";

import DeleteUserModal from "@/components/Tables/UsersTable/components/DeleteUserModal/DeleteUserModal.vue";
import { Button } from "@/components/ui/button";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { useUserStore } from "@/stores/user/user";
import type { User } from "@/types/user";

const { user: loggedUser } = useUserStore();
const router = useRouter();
const isDropdownOpen = ref(false);

const { user } = defineProps({
  user: {
    type: Object as PropType<User>,
    required: true,
  },
});

const handleEdit = () => {
  router.push({ name: "editUser", params: { id: user?.id } });
};

const closeDropdown = () => {
  isDropdownOpen.value = false;
};
</script>

<template>
  <DropdownMenu v-model:open="isDropdownOpen">
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
        <Button
          variant="ghost"
          size="icon"
          class="flex justify-start items-center w-full gap-2 py-2"
          @click="handleEdit"
        >
          <UserRoundPen />
          Edit User
        </Button>
      </DropdownMenuItem>
      <DropdownMenuItem
        :disabled="loggedUser?.email === user?.email"
        class="cursor-pointer"
      >
        <DeleteUserModal :user="user" @closeDropdown="closeDropdown" />
      </DropdownMenuItem>
    </DropdownMenuContent>
  </DropdownMenu>
</template>

<style scoped></style>
