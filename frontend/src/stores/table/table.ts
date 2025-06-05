import { defineStore } from "pinia";
import { ref } from "vue";

import { TABLE_PER_PAGE, TABLES_CONFIG, USERS_TABLE } from "@/constants";
import type { TablesState } from "@/types/tables.ts";


export const useTableStore = defineStore("table", () => {
  const page = ref(1);
  const itemsPerPage = ref(10);

  const setPage = (newPage: number) => {
    page.value = newPage;
  };

  const tableConfig: TablesState = JSON.parse(
    <string>localStorage.getItem(TABLES_CONFIG)
  ) ?? {};
  if (tableConfig) {
    const { users } = tableConfig;
    itemsPerPage.value = users?.[TABLE_PER_PAGE] ?? 10;
  }

  const setItemsPerPage = (newItemsPerPage: number) => {
    const data = {
      ...tableConfig,
      [USERS_TABLE]: {
        ...tableConfig?.[USERS_TABLE],
        [TABLE_PER_PAGE]: newItemsPerPage
      }
    };
    localStorage.setItem(TABLES_CONFIG, JSON.stringify(data));
    itemsPerPage.value = newItemsPerPage;
  };

  const $reset = () => {
    page.value = 1;
    itemsPerPage.value = 10;
  };

  return {
    page,
    itemsPerPage,
    setPage,
    setItemsPerPage,
    $reset
  };
});
