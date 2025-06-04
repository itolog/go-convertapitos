import { h, type VNode } from "vue";

import GitIcon from "@/assets/icons/git.svg";
import GoogleIcon from "@/assets/icons/google.svg";

export const providers = ["google", "github"] as const;
export type IconsKey = (typeof providers)[number];

export const icons: Record<IconsKey, VNode> = {
  google: h(GoogleIcon),
  github: h(GitIcon),
};
