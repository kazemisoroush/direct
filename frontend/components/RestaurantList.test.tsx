import { render, screen } from "@testing-library/react";
import { describe, expect, it } from "vitest";

import type { Restaurant } from "../lib/restaurants/restaurant";
import { RestaurantList } from "./RestaurantList";

describe("RestaurantList", () => {
  it("shows the empty state when no restaurants deliver here (M1)", () => {
    render(<RestaurantList restaurants={[]} />);
    expect(screen.getByRole("heading", { name: /no kitchens here yet/i })).toBeInTheDocument();
  });

  it("renders a card per restaurant", () => {
    const restaurants: Restaurant[] = [
      { id: "1", name: "Hills Kebabs", suburb: "Kellyville", address: "", deliversToPostcodes: ["2155"] },
    ];
    render(<RestaurantList restaurants={restaurants} />);
    expect(screen.getByText("Hills Kebabs")).toBeInTheDocument();
    expect(screen.getByText("Kellyville")).toBeInTheDocument();
  });
});
