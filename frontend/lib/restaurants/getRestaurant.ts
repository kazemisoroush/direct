import type { ApiClient } from "../api/client";
import type { Restaurant } from "./restaurant";

// getRestaurant returns one restaurant with its menu, or null when the id is unknown (404).
export async function getRestaurant(api: ApiClient, id: string): Promise<Restaurant | null> {
  const { data, error, response } = await api.GET("/restaurants/{id}", {
    params: { path: { id } },
  });
  if (response.status === 404) {
    return null;
  }
  if (error) {
    throw new Error("could not load restaurant");
  }
  return data ?? null;
}
