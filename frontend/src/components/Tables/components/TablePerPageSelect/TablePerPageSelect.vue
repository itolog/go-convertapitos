<script setup lang="ts">
import type { AcceptableValue } from "reka-ui";

import { options } from "@/components/Tables/components/TablePerPageSelect/data";
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { useTableStore } from "@/stores/table/table";

const tableStore = useTableStore();

const handleUpdate = (value: AcceptableValue) => {
  tableStore.setItemsPerPage(Number(value));
};
</script>

<template>
  <Select
    @update:modelValue="handleUpdate"
    :model-value="tableStore.itemsPerPage"
  >
    <SelectTrigger class="w-[80px] cursor-pointer">
      <SelectValue />
    </SelectTrigger>
    <SelectContent>
      <SelectGroup>
        <SelectItem
          class="cursor-pointer"
          :value="item"
          v-for="item in options"
          :key="item"
        >
          {{ item }}
        </SelectItem>
      </SelectGroup>
    </SelectContent>
  </Select>
</template>

<style scoped></style>
