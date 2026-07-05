import Link from "next/link";

import type { Restaurant } from "../lib/restaurants/restaurant";
import { EmptyState } from "./EmptyState";

// RestaurantList renders the restaurants delivering to the address (each a link to its menu),
// or an empty state.
export function RestaurantList({ restaurants }: { restaurants: Restaurant[] }) {
  if (restaurants.length === 0) {
    return <EmptyState />;
  }
  return (
    <ul className="restaurant-list">
      {restaurants.map((r) => (
        <li key={r.id}>
          <Link className="restaurant-card" href={{ pathname: "/menu", query: { id: r.id } }}>
            <div className="grow">
              <p className="name">{r.name}</p>
              <p className="meta">
                {r.suburb}
                {r.phone ? ` · ${r.phone}` : ""}
              </p>
            </div>
            <svg className="chev" width="18" height="18" viewBox="0 0 24 24" fill="none" aria-hidden="true">
              <path d="M9 6l6 6-6 6" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" />
            </svg>
          </Link>
        </li>
      ))}
    </ul>
  );
}
