import { APP_TABLES, HIDDEN_COLUMNS } from "@/constants";

export type TableKey = (typeof APP_TABLES)[number];

export interface TableConfig {
  [HIDDEN_COLUMNS]: {
    [key: string]: boolean;
  };
}

export type TablesState = {
  [key in TableKey]: TableConfig;
};
