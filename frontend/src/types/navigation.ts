import type { FunctionalComponent } from "vue";

export type NavigationItem = {
  title: string;
  url: string;
};
export type Navigation = {
  title: string;
  url: string;
  icon?: FunctionalComponent;
  isActive?: boolean;
  items?: NavigationItem[];
}[];
