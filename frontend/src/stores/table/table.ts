import { defineStore } from "pinia";
import { ref } from "vue";

export const useTableStore = defineStore("table", () => {
  const page = ref(1);
  const itemsPerPage = ref(10);

  const setPage = (newPage: number) => {
    page.value = newPage;
  };

  const setItemsPerPage = (newItemsPerPage: number) => {
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
    $reset,
  };
});
