import type { SelectOption } from "@/components/forms/EditUser/components/FormSelect/types.ts";

export const selectOptions: SelectOption[] = [
  {
    label: "True",
    value: true,
  },
  {
    label: "False",
    value: false,
  },
];

export const rolesOptions: SelectOption[] = [
  {
    label: "Regular",
    value: "regular",
  },
  {
    label: "Admin",
    value: "admin",
  },
  {
    label: "SuperUser",
    value: "superUser",
  },
];
