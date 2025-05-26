import { APP_TABLES, HIDDEN_COLUMNS, TABLE_PER_PAGE } from "@/constants";

export type TableKey = (typeof APP_TABLES)[number];

export interface TableConfig {
  [HIDDEN_COLUMNS]?: {
    [key: string]: boolean;
  };
  [TABLE_PER_PAGE]?: number;
}

export type TablesState = {
  [key in TableKey]: TableConfig;
};
