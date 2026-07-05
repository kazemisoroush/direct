import { formatPrice } from "../lib/money";
import type { MenuItem } from "../lib/restaurants/restaurant";

// groupByCategory groups items by category, preserving first-seen order for both.
function groupByCategory(menu: MenuItem[]): [string, MenuItem[]][] {
  const groups = new Map<string, MenuItem[]>();
  for (const item of menu) {
    const list = groups.get(item.category) ?? [];
    list.push(item);
    groups.set(item.category, list);
  }
  return [...groups.entries()];
}

// MenuView renders a restaurant's menu as receipt-style rows grouped by category.
export function MenuView({ menu }: { menu: MenuItem[] }) {
  if (menu.length === 0) {
    return <p className="loading">This restaurant hasn&rsquo;t listed a menu yet.</p>;
  }
  return (
    <div className="menu">
      {groupByCategory(menu).map(([category, items]) => (
        <section key={category}>
          <div className="rule">{category}</div>
          <ul className="menu-list">
            {items.map((item) => (
              <li key={item.id} className="menu-row">
                <div className="grow">
                  <p className="mi-name">{item.name}</p>
                  {item.description ? <p className="mi-desc">{item.description}</p> : null}
                </div>
                <span className="mi-price">{formatPrice(item.priceCents)}</span>
              </li>
            ))}
          </ul>
        </section>
      ))}
    </div>
  );
}
