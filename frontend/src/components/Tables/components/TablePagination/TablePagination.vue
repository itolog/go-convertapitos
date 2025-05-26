<script setup lang="ts">
import { Icon } from "@iconify/vue";
import {
  PaginationFirst,
  PaginationLast,
  PaginationList,
  PaginationListItem,
  PaginationPrev,
} from "reka-ui";

import {
  Pagination,
  PaginationEllipsis,
  PaginationNext,
} from "@/components/ui/pagination";
import { cn } from "@/lib/utils.ts";
import { useTableStore } from "@/stores/table/table.ts";

const tableStore = useTableStore();

const { total } = defineProps({
  total: {
    type: Number,
  },
});

const setPage = (page: number) => {
  tableStore.setPage(page);
};

const btnClass = "size-7 md:size-9";
</script>

<template>
  <Pagination
    :total="total"
    :sibling-count="1"
    :items-per-page="tableStore.itemsPerPage"
    show-edges
    :page="tableStore.page"
    :default-page="1"
    @update:page="setPage"
    class="flex md:justify-end"
  >
    <PaginationList
      v-slot="{ items }"
      class="flex items-center flex-wrap gap-1 text-stone-700 dark:text-white"
    >
      <PaginationFirst
        :class="
          cn(
            'cursor-pointer flex items-center justify-center bg-transparent hover:bg-white dark:hover:bg-stone-700/70 transition disabled:opacity-50 rounded-lg',
            btnClass,
          )
        "
      >
        <Icon icon="radix-icons:double-arrow-left" />
      </PaginationFirst>
      <PaginationPrev
        :class="
          cn(
            'cursor-pointer flex items-center justify-center bg-transparent hover:bg-white dark:hover:bg-stone-700/70 transition mr-4 disabled:opacity-50 rounded-lg',
            btnClass,
          )
        "
      >
        <Icon icon="radix-icons:chevron-left" />
      </PaginationPrev>
      <template v-for="(page, index) in items">
        <PaginationListItem
          v-if="page.type === 'page'"
          :key="index"
          :class="
            cn(
              'border cursor-pointer dark:border-stone-800 rounded-lg   hover:bg-white dark:hover:bg-stone-700/70 transition',
              'data-[selected]:bg-foreground data-[selected]:dark:text-slate-900 data-[selected]:shadow-sm data-[selected]:text-primary-foreground',
              btnClass,
            )
          "
          :value="page.value"
        >
          {{ page.value }}
        </PaginationListItem>
        <PaginationEllipsis
          v-else
          :key="page.type"
          :index="index"
          class="w-9 h-9 hidden md:flex items-center justify-center"
        >
          &#8230;
        </PaginationEllipsis>
      </template>
      <PaginationNext
        :class="
          cn(
            'cursor-pointer flex items-center justify-center bg-transparent hover:bg-white dark:hover:bg-stone-700/70 transition ml-4 disabled:opacity-50 rounded-lg',
            btnClass,
          )
        "
      >
        <Icon icon="radix-icons:chevron-right" />
      </PaginationNext>
      <PaginationLast
        :class="
          cn(
            'cursor-pointer flex items-center justify-center bg-transparent hover:bg-white dark:hover:bg-stone-700/70 transition disabled:opacity-50 rounded-lg',
            btnClass,
          )
        "
      >
        <Icon icon="radix-icons:double-arrow-right" />
      </PaginationLast>
    </PaginationList>
  </Pagination>
</template>

<style scoped></style>
