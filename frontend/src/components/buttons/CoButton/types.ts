import type { PrimitiveProps } from "reka-ui";
import type { HTMLAttributes } from "vue";

import type { ButtonVariants } from "@/components/ui/button";

export interface CoButtonProps extends PrimitiveProps {
  variant?: ButtonVariants["variant"];
  size?: ButtonVariants["size"];
  class?: HTMLAttributes["class"];
}
