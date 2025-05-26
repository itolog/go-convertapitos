<script setup lang="ts">
import {
  FlexRender,
  getCoreRowModel,
  getExpandedRowModel,
  getFilteredRowModel,
  getSortedRowModel,
  useVueTable,
} from "@tanstack/vue-table";
import { SquareX } from "lucide-vue-next";
import { type PropType, ref, watch } from "vue";

import UserTableHeader from "@/components/Tables/UsersTable/components/UserTableHeader/UserTableHeader.vue";
import { useColumns } from "@/components/Tables/UsersTable/hooks/useColumns.ts";
import { useTableConfig } from "@/components/Tables/UsersTable/hooks/useTableConfig.ts";
import TableItemsInfo from "@/components/Tables/components/TableItemsInfo/TableItemsInfo.vue";
import TablePagination from "@/components/Tables/components/TablePagination/TablePagination.vue";
import TablePerPageSelect from "@/components/Tables/components/TablePerPageSelect/TablePerPageSelect.vue";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Skeleton } from "@/components/ui/skeleton";
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
import type { ApiMeta } from "@/generated/apiClient/data-contracts";
import { cn } from "@/lib/utils.ts";
import { useTableStore } from "@/stores/table/table.ts";
import { useUserStore } from "@/stores/user/user.ts";
import type { TablesState } from "@/types/tables.ts";
import type { User } from "@/types/user.ts";

const showFilters = ref(false);
const { user: loggedUser } = useUserStore();
const tableStore = useTableStore();

const { data, isFetching, meta, isLoading } = defineProps({
  data: {
    type: Array<User>,
    required: true,
  },
  meta: {
    type: Object as PropType<ApiMeta>,
    required: true,
  },
  isLoading: Boolean,
  isFetching: Boolean,
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
  get data() {
    return data;
  },
  columns,
  getCoreRowModel: getCoreRowModel(),
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
  getRowId: (row) => row?.id ?? "",
  enableRowSelection: (row) => row.original?.id !== loggedUser?.id,
  defaultColumn: {
    size: 180,
    minSize: 25,
    maxSize: 300,
  },
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
  () => data,
  () => {
    table.options.data = data;
  },
  { deep: true },
);

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

const handleShowFilters = (value: boolean) => {
  showFilters.value = value;
};
</script>

<template>
  <div class="flex flex-col h-full">
    <UserTableHeader @update:showFilters="handleShowFilters" :table="table" />

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
              :class="
                cn({
                  'table-sticky-column table-sticky-left':
                    header.column.getIsPinned() === 'left',
                  'table-sticky-column table-sticky-right':
                    header.column.getIsPinned() === 'right',
                })
              "
            >
              <div
                class="flex h-full relative flex-col py-1 gap-1 items-center"
              >
                <FlexRender
                  v-if="!header.isPlaceholder"
                  :render="header.column.columnDef.header"
                  :props="header.getContext()"
                />
                <div
                  class="flex absolute top-1 z-1 items-center bg-background"
                  v-if="
                    table.getColumn(header.id)?.getCanFilter() && showFilters
                  "
                >
                  <Input
                    :name="
                      table
                        .getColumn(table.getColumn(header.id)?.id ?? '')
                        ?.getFilterValue()
                    "
                    :placeholder="`Filter ${table.getColumn(header.id)?.id}`"
                    class="pr-8"
                    :style="`width: ${header.getSize()}px`"
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
                  <Button
                    size="icon"
                    variant="link"
                    class="absolute right-0 z-1 size-8"
                    :disabled="
                      !table
                        .getColumn(table.getColumn(header.id)?.id ?? '')
                        ?.getFilterValue()
                    "
                    @click="
                      () => {
                        table
                          .getColumn(table.getColumn(header.id)?.id ?? '')
                          ?.setFilterValue('');
                      }
                    "
                  >
                    <SquareX />
                  </Button>
                </div>
              </div>
            </TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          <template v-if="isLoading">
            <TableRow v-for="item in tableStore.itemsPerPage" :key="item">
              <TableCell class="h-[49px] w-full" :colspan="columns.length">
                <Skeleton class="h-full w-full" />
              </TableCell>
            </TableRow>
          </template>

          <template v-else-if="table.getRowModel().rows?.length && !isLoading">
            <template v-for="row in table.getRowModel().rows" :key="row.id">
              <TableRow
                :class="
                  cn({
                    'animate-pulse text-gray-400': isFetching,
                  })
                "
                :data-state="row.getIsSelected() && 'selected'"
              >
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

    <div
      class="flex flex-col md:flex-row items-center justify-between space-x-2 gap-4 py-4"
    >
      <div class="flex items-center gap-2">
        <TableItemsInfo :meta="meta" />
        <TablePerPageSelect />
      </div>

      <TablePagination :total="meta.items" />
    </div>
  </div>
</template>

<style scoped>
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
