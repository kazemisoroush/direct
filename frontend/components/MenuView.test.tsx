import { render, screen } from "@testing-library/react";
import { describe, expect, it } from "vitest";

import type { MenuItem } from "../lib/restaurants/restaurant";
import { MenuView } from "./MenuView";

describe("MenuView", () => {
  it("renders items grouped by category with formatted prices", () => {
    const menu: MenuItem[] = [
      { id: "beef-kebab", name: "Beef Kebab", priceCents: 1500, category: "Kebabs" },
      { id: "pide", name: "Pide", priceCents: 1200, category: "Pide" },
    ];
    render(<MenuView menu={menu} />);
    expect(screen.getByText("Beef Kebab")).toBeInTheDocument();
    expect(screen.getByText("$15.00")).toBeInTheDocument();
    expect(screen.getByText("Kebabs")).toBeInTheDocument();
    expect(screen.getByText("$12.00")).toBeInTheDocument();
    // "Pide" is both a category and an item name here, so it appears twice.
    expect(screen.getAllByText("Pide")).toHaveLength(2);
  });

  it("handles a restaurant with no menu", () => {
    render(<MenuView menu={[]} />);
    expect(screen.getByText(/hasn.t listed a menu yet/i)).toBeInTheDocument();
  });
});
