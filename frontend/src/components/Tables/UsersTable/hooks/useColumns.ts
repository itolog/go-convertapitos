import {
  type CellContext,
  type Column,
  createColumnHelper,
} from "@tanstack/vue-table";
import { useDateFormat } from "@vueuse/core";
import {
  ArrowUpDown,
  CircleUserRound,
  EllipsisVerticalIcon,
} from "lucide-vue-next";
import { h } from "vue";

import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { Button } from "@/components/ui/button";
import { Checkbox } from "@/components/ui/checkbox";
import type { User } from "@/types/user";

const headerSortColumn = (column: Column<User>, name?: string) => {
  return h(
    Button,
    {
      variant: "ghost",
      class: "w-full justify-between items-center",
      size: "icon",
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
    { class: "flex items-center capitalize h-9" },
    name ?? column.id,
  );
};

const textCell = ({ getValue }: CellContext<User, string | undefined>) => {
  const id = getValue();
  return h("div", { class: "text-left" }, id);
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
      header: ({ table }) => {
        return h(
          "div",
          { class: "flex items-center justify-center h-9" },
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
      cell: ({ row }) => {
        return h(
          "div",
          { class: " " },
          h(Checkbox, {
            modelValue: row.getIsSelected(),
            "onUpdate:modelValue": (value) => row.toggleSelected(!!value),
            ariaLabel: "Select row",
          }),
        );
      },
      enableSorting: false,
      enableHiding: false,
      size: 15,
    }),
    columnHelper.accessor("picture", {
      header: ({ column }) => headerColumn(column),
      cell: ({ getValue }) => {
        const pictureUrl = getValue();
        return h("div", { class: "flex justify-center items-center" }, [
          h(Avatar, null, {
            default: () => {
              return [
                // eslint-disable-next-line @typescript-eslint/ban-ts-comment
                // @ts-expect-error
                h(AvatarImage, { src: pictureUrl }),
                h(AvatarFallback, null, () => h(CircleUserRound)),
              ];
            },
          }),
        ]);
      },
      size: 30,
      enableColumnFilter: false,
    }),
    columnHelper.accessor("id", {
      header: ({ column }) => headerSortColumn(column, "ID"),
      cell: textCell,
    }),
    columnHelper.accessor("name", {
      header: ({ column }) => headerSortColumn(column),
      cell: textCell,
    }),
    columnHelper.accessor("email", {
      header: ({ column }) => headerSortColumn(column),
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

        return h(EllipsisVerticalIcon, {
          class: "cursor-pointer",
          user,
          onExpand: row.toggleExpanded,
        });
      },
      size: 25,
    }),
  ];
};
