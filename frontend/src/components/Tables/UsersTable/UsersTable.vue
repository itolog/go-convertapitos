<script setup lang="ts">
import {
  FlexRender,
  getCoreRowModel,
  getExpandedRowModel,
  getFilteredRowModel,
  getPaginationRowModel,
  getSortedRowModel,
  useVueTable,
} from "@tanstack/vue-table";
import { ChevronDown, FunnelPlus, FunnelX } from "lucide-vue-next";
import { ref, watch } from "vue";

import { useColumns } from "@/components/Tables/UsersTable/hooks/useColumns.ts";
import { useTableConfig } from "@/components/Tables/UsersTable/hooks/useTableConfig.ts";
import { Button } from "@/components/ui/button";
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
import { HIDDEN_COLUMNS, TABLES_CONFIG, USERS_TABLE } from "@/constants";
import type { TablesState } from "@/types/tables.ts";
import type { User } from "@/types/user.ts";

const showFilters = ref(false);

const { data } = defineProps({
  data: {
    type: Array<User>,
    required: true,
  },
  loading: Boolean,
});

const columns = useColumns();
const {
  columnPinning,
  sorting,
  columnFilters,
  columnVisibility,
  rowSelection,
  expanded,
} = useTableConfig();

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

watch(
  () => table.getState().columnVisibility,
  () => {
    const data: TablesState = {
      [USERS_TABLE]: {
        [HIDDEN_COLUMNS]: table.getState().columnVisibility,
      },
    };
    localStorage.setItem(TABLES_CONFIG, JSON.stringify(data));
  },
);
</script>

<template>
  <div class="flex flex-col h-full">
    <div class="flex items-center gap-4 py-4">
      <!--      <div-->
      <!--        v-for="column in table-->
      <!--          .getAllColumns()-->
      <!--          .filter((col) => col.getCanFilter())"-->
      <!--        :key="column.id"-->
      <!--        class="space-y-1"-->
      <!--      >-->
      <!--        <div class="text-sm font-medium">-->
      <!--          {{ column.id.charAt(0).toUpperCase() + column.id.slice(1) }}-->
      <!--        </div>-->
      <!--        <template v-if="column.id === 'email'">-->
      <!--          <Input-->
      <!--            class="max-w-sm"-->
      <!--            placeholder="Filter emails..."-->
      <!--            :model-value="column.getFilterValue() as string"-->
      <!--            @update:model-value="column.setFilterValue($event)"-->
      <!--          />-->
      <!--        </template>-->
      <!--        <template v-else-if="column.id === 'verifiedEmail'">-->
      <!--          <select-->
      <!--            class="w-full rounded-md border border-input bg-background px-3 py-2 text-sm"-->
      <!--            :value="column.getFilterValue() as string"-->
      <!--            @change="-->
      <!--              (e: Event) => {-->
      <!--                column.setFilterValue(e.target.value);-->
      <!--              }-->
      <!--            "-->
      <!--          >-->
      <!--            <option value="">Все</option>-->
      <!--            <option value="true">Да</option>-->
      <!--            <option value="false">Нет</option>-->
      <!--          </select>-->
      <!--        </template>-->
      <!--        <template v-else>-->
      <!--          <Input-->
      <!--            class="max-w-sm"-->
      <!--            :placeholder="`Поиск по ${column.id}...`"-->
      <!--            :model-value="column.getFilterValue() as string"-->
      <!--            @update:model-value="column.setFilterValue($event)"-->
      <!--          />-->
      <!--        </template>-->
      <!--      </div>-->

      <!--      <Button-->
      <!--        variant="outline"-->
      <!--        size="sm"-->
      <!--        @click="-->
      <!--          () => {-->
      <!--            table.resetColumnFilters();-->
      <!--          }-->
      <!--        "-->
      <!--        v-if="table.getState().columnFilters.length > 0"-->
      <!--      >-->
      <!--        Сбросить фильтры-->
      <!--      </Button>-->
      <!--      <Input-->
      <!--        class="max-w-sm"-->
      <!--        placeholder="Filter emails..."-->
      <!--        :model-value="table.getColumn('email')?.getFilterValue() as string"-->
      <!--        @update:model-value="table.getColumn('email')?.setFilterValue($event)"-->
      <!--      />-->
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
      <Button
        variant="outline"
        class="h-9"
        @click="
          () => {
            showFilters = !showFilters;
          }
        "
      >
        <FunnelPlus v-if="!showFilters" />
        <FunnelX v-else />
      </Button>
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
              <div class="flex h-full flex-col items-start">
                <FlexRender
                  v-if="!header.isPlaceholder"
                  :render="header.column.columnDef.header"
                  :props="header.getContext()"
                />
                <Input
                  v-if="
                    table.getColumn(header.id)?.getCanFilter() && showFilters
                  "
                  style="margin-bottom: 8px;}"
                  :placeholder="`Filter ${table.getColumn(header.id)?.id}`"
                  :model-value="
                    table
                      .getColumn(table.getColumn(header.id)?.id ?? '')
                      ?.getFilterValue() as string
                  "
                  @update:model-value="
                    table
                      .getColumn(table.getColumn(header.id)?.id ?? '')
                      ?.setFilterValue($event)
                  "
                />
              </div>
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
