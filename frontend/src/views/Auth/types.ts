import type { VNode } from "vue";

export const providers = ["google", "github"] as const;
export type IconsKey = (typeof providers)[number];

export interface SocialLink {
  name: IconsKey;
  icon: VNode;
}
