import { h } from "vue";

import GitIcon from "@/assets/icons/git.svg";
import GoogleIcon from "@/assets/icons/google.svg";
import type { SocialLink } from "@/views/Auth/types.ts";

export const socialLinks: SocialLink[] = [
  {
    name: "google",
    icon: h(GoogleIcon),
  },
  {
    name: "github",
    icon: h(GitIcon),
  },
];
