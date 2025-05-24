<script setup lang="ts">
import type {
  ColumnFiltersState,
  ColumnPinningState,
  ExpandedState,
  SortingState,
  VisibilityState,
} from "@tanstack/vue-table";
import {
  FlexRender,
  getCoreRowModel,
  getExpandedRowModel,
  getFilteredRowModel,
  getPaginationRowModel,
  getSortedRowModel,
  useVueTable,
  createColumnHelper,
} from "@tanstack/vue-table";
import { useDateFormat } from "@vueuse/core";
import {
  ArrowUpDown,
  ChevronDown,
  CircleUserRound,
  EllipsisVerticalIcon,
} from "lucide-vue-next";
import { h, ref } from "vue";

import { Avatar, AvatarImage, AvatarFallback } from "@/components/ui/avatar";
import { Button } from "@/components/ui/button";
import { Checkbox } from "@/components/ui/checkbox";
import {
  DropdownMenu,
  DropdownMenuCheckboxItem,
  DropdownMenuContent,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { Input } from "@/components/ui/input";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { valueUpdater } from "@/components/ui/table/utils.ts";
import type { User } from "@/types/user.ts";

const { data } = defineProps({
  data: {
    type: Array<User>,
    required: true,
  },
  loading: Boolean,
});

const columnHelper = createColumnHelper<User>();

const columns = [
  columnHelper.display({
    id: "select",
    header: ({ table }) => {
      return h(Checkbox, {
        modelValue:
          table.getIsAllPageRowsSelected() ||
          (table.getIsSomePageRowsSelected() && "indeterminate"),
        "onUpdate:modelValue": (value) =>
          table.toggleAllPageRowsSelected(!!value),
        ariaLabel: "Select all",
      });
    },
    cell: ({ row }) => {
      return h(Checkbox, {
        modelValue: row.getIsSelected(),
        "onUpdate:modelValue": (value) => row.toggleSelected(!!value),
        ariaLabel: "Select row",
      });
    },
    enableSorting: false,
    enableHiding: false,
  }),
  columnHelper.accessor("picture", {
    header: () => h("div", { class: "text-left" }, "Avatar"),
    cell: ({ getValue }) => {
      const pictureUrl = getValue();
      return h("div", { class: "max-w-[60px]" }, [
        h(Avatar, null, {
          default: () => {
            return [
              h(AvatarImage, { src: pictureUrl }),
              h(AvatarFallback, null, () => h(CircleUserRound)),
            ];
          },
        }),
      ]);
    },
  }),
  columnHelper.accessor("id", {
    header: () => h("div", { class: "text-left" }, "ID"),
    cell: ({ getValue }) => {
      const id = getValue();
      return h("div", { class: "text-left" }, id);
    },
  }),
  columnHelper.accessor("name", {
    header: ({ column }) => {
      return h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        () => ["Name", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })],
      );
    },
    cell: ({ row }) => {
      return h("div", { class: "capitalize" }, row.getValue("name"));
    },
  }),
  columnHelper.accessor("email", {
    header: ({ column }) => {
      return h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        () => ["Email", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })],
      );
    },
  }),
  columnHelper.accessor("createdAt", {
    header: ({ column }) => {
      return h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        () => ["Created At", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })],
      );
    },
    cell: ({ row }) => {
      return h("div", { class: "lowercase" }, [
        useDateFormat(row.getValue("createdAt"), "YYYY-MM-DD HH:mm").value,
      ]);
    },
  }),
  columnHelper.accessor("updatedAt", {
    header: ({ column }) => {
      return h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        () => ["Update dAt", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })],
      );
    },
    cell: ({ row }) => {
      return h("div", { class: "lowercase" }, [
        useDateFormat(row.getValue("updatedAt"), "YYYY-MM-DD HH:mm").value,
      ]);
    },
  }),
  columnHelper.accessor("verifiedEmail", {
    header: "Verified Email",
  }),
  columnHelper.display({
    id: "actions",
    // enableHiding: false,
    cell: ({ row }) => {
      const user = row.original;

      return h(EllipsisVerticalIcon, {
        user,
        onExpand: row.toggleExpanded,
      });
    },
  }),
];

const sorting = ref<SortingState>([
  {
    id: "updatedAt",
    desc: true,
  },
]);
const columnFilters = ref<ColumnFiltersState>([]);
const columnVisibility = ref<VisibilityState>({});
const rowSelection = ref({});
const expanded = ref<ExpandedState>({});
const columnPinning = ref<ColumnPinningState>({
  left: ["select"],
  right: ["actions"],
});

const table = useVueTable({
  data,
  columns,
  getCoreRowModel: getCoreRowModel(),
  getPaginationRowModel: getPaginationRowModel(),
  getSortedRowModel: getSortedRowModel(),
  getFilteredRowModel: getFilteredRowModel(),
  getExpandedRowModel: getExpandedRowModel(),
  columnResizeMode: "onChange",
  onSortingChange: (updaterOrValue) => valueUpdater(updaterOrValue, sorting),
  onColumnFiltersChange: (updaterOrValue) =>
    valueUpdater(updaterOrValue, columnFilters),
  onColumnVisibilityChange: (updaterOrValue) =>
    valueUpdater(updaterOrValue, columnVisibility),
  onRowSelectionChange: (updaterOrValue) =>
    valueUpdater(updaterOrValue, rowSelection),
  onExpandedChange: (updaterOrValue) => valueUpdater(updaterOrValue, expanded),
  onColumnPinningChange: (updaterOrValue) =>
    valueUpdater(updaterOrValue, columnPinning),
  state: {
    get columnPinning() {
      return columnPinning.value;
    },
    get sorting() {
      return sorting.value;
    },
    get columnFilters() {
      return columnFilters.value;
    },
    get columnVisibility() {
      return columnVisibility.value;
    },
    get rowSelection() {
      return rowSelection.value;
    },
    get expanded() {
      return expanded.value;
    },
  },
});
</script>

<template>
  <div class="flex flex-col h-full">
    <div class="flex items-center gap-4 py-4">
      <Input
        class="max-w-sm"
        placeholder="Filter emails..."
        :model-value="table.getColumn('email')?.getFilterValue() as string"
        @update:model-value="table.getColumn('email')?.setFilterValue($event)"
      />
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
              .filter((column) => column.getCanHide())"
            :key="column.id"
            class="capitalize"
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
    </div>
    <div class="rounded-md border flex-1">
      <Table style="height: 100%">
        <TableHeader>
          <TableRow
            v-for="headerGroup in table.getHeaderGroups()"
            :key="headerGroup.id"
          >
            <TableHead
              v-for="header in headerGroup.headers"
              :key="header.id"
              :style="{ width: `${header.getSize()}px` }"
              :colSpan="header.colSpan"
              :class="[
                header.column.getIsPinned() === 'left'
                  ? 'table-sticky-column table-sticky-left'
                  : '',
                header.column.getIsPinned() === 'right'
                  ? 'table-sticky-column table-sticky-right'
                  : '',
              ]"
            >
              <FlexRender
                v-if="!header.isPlaceholder"
                :render="header.column.columnDef.header"
                :props="header.getContext()"
              />
            </TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          <template v-if="table.getRowModel().rows?.length">
            <template v-for="row in table.getRowModel().rows" :key="row.id">
              <TableRow :data-state="row.getIsSelected() && 'selected'">
                <TableCell
                  v-for="cell in row.getVisibleCells()"
                  :key="cell.id"
                  :class="[
                    cell.column.getIsPinned() === 'left'
                      ? 'table-sticky-column table-sticky-left'
                      : '',
                    cell.column.getIsPinned() === 'right'
                      ? 'table-sticky-column table-sticky-right'
                      : '',
                  ]"
                >
                  <FlexRender
                    :render="cell.column.columnDef.cell"
                    :props="cell.getContext()"
                  />
                </TableCell>
              </TableRow>
              <TableRow v-if="row.getIsExpanded()">
                <TableCell :colspan="row.getAllCells().length">
                  {{ JSON.stringify(row.original) }}
                </TableCell>
              </TableRow>
            </template>
          </template>

          <TableRow class="hover:bg-background" v-else>
            <TableCell
              :colspan="columns.length"
              class="h-24 text-2xl text-center font-bold"
            >
              No results.
            </TableCell>
          </TableRow>
        </TableBody>
      </Table>
    </div>

    <div class="flex items-center justify-end space-x-2 gap-2 py-4">
      <div class="flex-1 text-sm text-muted-foreground">
        {{ table.getFilteredSelectedRowModel().rows.length }} of
        {{ table.getFilteredRowModel().rows.length }} row(s) selected.
      </div>
      <div class="flex space-x-2 gap-2">
        <Button
          variant="outline"
          size="sm"
          :disabled="!table.getCanPreviousPage()"
          @click="table.previousPage()"
        >
          Previous
        </Button>
        <Button
          variant="outline"
          size="sm"
          :disabled="!table.getCanNextPage()"
          @click="table.nextPage()"
        >
          Next
        </Button>
      </div>
    </div>
  </div>
</template>

<style>
.table-sticky-column {
  position: sticky !important;
  z-index: 1;
  background-color: var(--background, #fff);
}

.table-sticky-left {
  left: 0;
  box-shadow: 2px 0 3px -1px rgba(0, 0, 0, 0.1);
}

.table-sticky-right {
  right: 0;
  box-shadow: -2px 0 3px -1px rgba(0, 0, 0, 0.1);
}
</style>
