import type {
  ColumnFiltersState,
  ColumnPinningState,
  ExpandedState,
  SortingState,
  VisibilityState,
} from "@tanstack/vue-table";
import { ref } from "vue";

import { HIDDEN_COLUMNS, TABLES_CONFIG } from "@/constants";
import type { TablesState } from "@/types/tables.ts";

export function useTableConfig() {
  const sorting = ref<SortingState>([
    {
      id: "updatedAt",
      desc: true,
    },
  ]);

  const tablesConfig: TablesState = JSON.parse(
    localStorage.getItem(TABLES_CONFIG) ?? "{}",
  );

  const columnFilters = ref<ColumnFiltersState>([]);
  const columnVisibility = ref<VisibilityState>(
    tablesConfig?.users?.[HIDDEN_COLUMNS] ?? {},
  );
  const rowSelection = ref({});
  const expanded = ref<ExpandedState>({});
  const columnPinning = ref<ColumnPinningState>({
    left: ["select"],
    right: ["actions"],
  });

  return {
    sorting,
    columnFilters,
    columnVisibility,
    rowSelection,
    expanded,
    columnPinning,
  };
}
