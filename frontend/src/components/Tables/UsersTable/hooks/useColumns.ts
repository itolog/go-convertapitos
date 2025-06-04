import {
  type CellContext,
  type Column,
  createColumnHelper,
} from "@tanstack/vue-table";
import { useDateFormat } from "@vueuse/core";
import { ArrowUpDown } from "lucide-vue-next";
import { h } from "vue";

import TableCelAction from "@/components/Tables/UsersTable/components/TableCelAction/TableCelAction.vue";
import UserAvatar from "@/components/common/ui/UserAvatar/UserAvatar.vue";
import { Button } from "@/components/ui/button";
import { Checkbox } from "@/components/ui/checkbox";
import type { User } from "@/types/user";

const headerSortColumn = (column: Column<User>, name?: string) => {
  return h(
    Button,
    {
      variant: "ghost",
      class: "w-full justify-between px-2 items-center",
      size: "icon",
      style: {
        width: `${column.getSize()}px`,
      },
      onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
    },
    () => [
      h("span", { class: "capitalize" }, name ?? column.id),
      h(ArrowUpDown, { class: "ml-2 h-4 w-4" }),
    ],
  );
};

const headerColumn = (column: Column<User>, name?: string) => {
  return h(
    "div",
    {
      class: "flex items-center capitalize h-9",
      style: {
        width: `${column.getSize()}px`,
      },
    },
    name ?? column.id,
  );
};

const textCell = ({
  getValue,
  column,
}: CellContext<User, string | undefined>) => {
  const id = getValue();

  return h(
    "div",
    {
      class: `text-left break-all whitespace-normal line-clamp-2`,
      style: {
        width: `${column.getSize()}px`,
      },
    },
    id,
  );
};

const dataCell = ({ getValue }: CellContext<User, string | undefined>) => {
  return h("div", { class: "lowercase" }, [
    useDateFormat(getValue(), "YYYY-MM-DD HH:mm").value,
  ]);
};

export const useColumns = () => {
  const columnHelper = createColumnHelper<User>();
  return [
    columnHelper.display({
      id: "select",
      header: ({ table, column }) => {
        return h(
          "div",
          {
            class: "flex items-center justify-center h-9",
            style: {
              width: `${column.getSize()}px`,
            },
          },
          h(Checkbox, {
            modelValue:
              table.getIsAllPageRowsSelected() ||
              (table.getIsSomePageRowsSelected() && "indeterminate"),
            "onUpdate:modelValue": (value) =>
              table.toggleAllPageRowsSelected(!!value),
            ariaLabel: "Select all",
          }),
        );
      },
      cell: ({ row, column }) => {
        return h(
          "div",
          {
            class: "flex items-center justify-center",
            style: {
              width: `${column.getSize()}px`,
            },
          },
          h(Checkbox, {
            disabled: !row.getCanSelect(),
            modelValue: row.getIsSelected(),
            "onUpdate:modelValue": (value) => row.toggleSelected(!!value),
            ariaLabel: "Select row",
          }),
        );
      },
      enableSorting: false,
      enableHiding: false,
      enableResizing: false,
      size: 30,
      minSize: 25,
      maxSize: 35,
    }),
    columnHelper.accessor("picture", {
      header: ({ column }) => headerColumn(column),
      cell: ({ row }) => {
        return h("div", { class: "flex justify-center items-center" }, [
          h(UserAvatar, { user: row.original, class: "size-7" }),
        ]);
      },
      size: 35,
      enableColumnFilter: false,
    }),
    columnHelper.accessor("id", {
      header: ({ column }) => headerSortColumn(column, "ID"),
      cell: textCell,
      size: 300,
    }),
    columnHelper.accessor("name", {
      header: ({ column }) => headerSortColumn(column),
      cell: textCell,
    }),
    columnHelper.accessor("email", {
      header: ({ column }) => headerSortColumn(column),
      cell: textCell,
    }),
    columnHelper.accessor("createdAt", {
      header: ({ column }) => headerSortColumn(column),
      cell: dataCell,
    }),
    columnHelper.accessor("updatedAt", {
      header: ({ column }) => headerSortColumn(column),
      cell: dataCell,
    }),
    columnHelper.accessor("verifiedEmail", {
      header: ({ column }) => headerColumn(column, "Verified Email"),
    }),
    columnHelper.display({
      id: "actions",
      cell: ({ row }) => {
        const user = row.original;

        return h(TableCelAction, {
          user,
          onExpand: row.toggleExpanded,
        });
      },
      size: 40,
      minSize: 40,
      maxSize: 60,
      enableResizing: false,
      enableHiding: false,
    }),
  ];
};
