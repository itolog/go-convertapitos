import type { UserUser } from "@/generated/apiClient/data-contracts.ts";

export type User = UserUser | null;

export type UserAuth = {
  accessToken: string;
  user: User;
};
