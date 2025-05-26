<script setup lang="ts">
import type { Table } from "@tanstack/vue-table";
import {
  ChevronDown,
  Delete,
  FunnelPlus,
  FunnelX,
  UserPlus,
} from "lucide-vue-next";
import { type PropType, ref } from "vue";
import { useRouter } from "vue-router";

import { Button } from "@/components/ui/button";
import {
  DropdownMenu,
  DropdownMenuCheckboxItem,
  DropdownMenuContent,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import type { User } from "@/types/user.ts";

const showFilters = ref(false);
const emits = defineEmits(["update:showFilters"]);
const { table } = defineProps({
  table: {
    type: Object as PropType<Table<User>>,
    required: true,
  },
});

const router = useRouter();

const handleFilters = () => {
  showFilters.value = !showFilters.value;
  emits("update:showFilters", !showFilters.value);
};

const handleAddUser = () => {
  router.push({ name: "addUser" });
};
</script>

<template>
  <div class="flex items-center justify-between gap-4 py-4">
    <!--   FILTERS   -->
    <div class="flex items-center gap-4">
      <DropdownMenu>
        <DropdownMenuTrigger as-child>
          <Button variant="outline" class="ml-auto">
            Columns <ChevronDown class="ml-2 h-4 w-4" />
          </Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent align="end">
          <DropdownMenuCheckboxItem
            v-for="column in table
              .getAllColumns()
              .filter((item) => item.getCanHide())"
            :key="column.id"
            class="capitalize cursor-pointer"
            :model-value="column.getIsVisible()"
            @update:model-value="
              (value) => {
                column.toggleVisibility(!!value);
              }
            "
          >
            {{ column.id }}
          </DropdownMenuCheckboxItem>
        </DropdownMenuContent>
      </DropdownMenu>
      <Button variant="outline" size="icon" @click="handleFilters">
        <FunnelPlus v-if="!showFilters" />
        <FunnelX v-else />
      </Button>

      <Button
        variant="outline"
        size="icon"
        @click="table.resetColumnFilters"
        v-if="table.getState().columnFilters.length > 0"
      >
        <Delete />
      </Button>
    </div>
    <!--   ACTIONS   -->
    <div>
      <Button size="icon" @click="handleAddUser">
        <UserPlus />
      </Button>
    </div>
  </div>
</template>

<style scoped></style>
