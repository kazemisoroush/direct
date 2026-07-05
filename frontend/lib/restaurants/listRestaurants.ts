import type { ApiClient } from "../api/client";
import type { Restaurant } from "./restaurant";

// listRestaurants returns the restaurants that deliver to the given address.
export async function listRestaurants(api: ApiClient, address: string): Promise<Restaurant[]> {
  const { data } = await api.GET("/restaurants", { params: { query: { address } } });
  return data?.restaurants ?? [];
}
