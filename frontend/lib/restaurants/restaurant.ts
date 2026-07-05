import type { components } from "../api/schema";

// Restaurant is a restaurant record as defined by openapi.yaml.
export type Restaurant = components["schemas"]["Restaurant"];

// MenuItem is one orderable menu item as defined by openapi.yaml.
export type MenuItem = components["schemas"]["MenuItem"];
