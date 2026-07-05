import type { Restaurant } from "../lib/restaurants/restaurant";
import { EmptyState } from "./EmptyState";

// RestaurantList renders the restaurants delivering to the address, or an empty state.
export function RestaurantList({ restaurants }: { restaurants: Restaurant[] }) {
  if (restaurants.length === 0) {
    return <EmptyState />;
  }
  return (
    <ul className="restaurant-list">
      {restaurants.map((r) => (
        <li key={r.id}>
          <button className="restaurant-card" type="button">
            <div>
              <p className="name">{r.name}</p>
              <p className="meta">{r.suburb}</p>
            </div>
          </button>
        </li>
      ))}
    </ul>
  );
}
