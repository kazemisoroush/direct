import type { ApiClient } from "../api/client";
import type { Restaurant } from "./restaurant";

// listRestaurants returns the restaurants that deliver to the given address.
export async function listRestaurants(api: ApiClient, address: string): Promise<Restaurant[]> {
  // openapi-fetch does not throw on a non-2xx: it returns { error }. Surface it so a backend
  // failure shows an error, not an empty "no restaurants here" state.
  const { data, error } = await api.GET("/restaurants", { params: { query: { address } } });
  if (error) {
    throw new Error("could not load restaurants");
  }
  return data?.restaurants ?? [];
}
