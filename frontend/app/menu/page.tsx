"use client";

import { Suspense, useCallback, useEffect, useState } from "react";
import Link from "next/link";
import { useRouter, useSearchParams } from "next/navigation";

import { MenuView } from "../../components/MenuView";
import { ThemeToggle } from "../../components/ThemeToggle";
import { useAuth } from "../../lib/auth/context";
import { getRestaurant } from "../../lib/restaurants/getRestaurant";
import type { Restaurant } from "../../lib/restaurants/restaurant";

function MenuPageInner() {
  const { ready, authenticated, api, signOut } = useAuth();
  const router = useRouter();
  const id = useSearchParams().get("id");
  // undefined = still loading, null = not found.
  const [restaurant, setRestaurant] = useState<Restaurant | null | undefined>(undefined);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    if (ready && !authenticated) {
      router.replace("/login");
    }
  }, [ready, authenticated, router]);

  const load = useCallback(async () => {
    if (!api || !id) return;
    setError(null);
    try {
      setRestaurant(await getRestaurant(api, id));
    } catch (err) {
      setError(err instanceof Error ? err.message : "could not load the menu");
    }
  }, [api, id]);

  useEffect(() => {
    if (authenticated) {
      void load();
    }
  }, [authenticated, load]);

  if (!ready) {
    return <main className="loading">Loading…</main>;
  }
  if (!authenticated) {
    return null;
  }

  return (
    <>
      <header className="topbar">
        <Link className="back" href="/">
          <span aria-hidden="true">‹</span> All restaurants
        </Link>
        <div className="topbar-actions">
          <ThemeToggle />
          <button className="link" onClick={signOut}>
            Sign out
          </button>
        </div>
      </header>
      <main className="page">
        {error && <p role="alert">{error}</p>}
        {restaurant === undefined && !error && <p className="loading">Loading menu…</p>}
        {restaurant === null && <p className="loading">We couldn&rsquo;t find that restaurant.</p>}
        {restaurant ? (
          <>
            <div className="rest-head">
              <h1>{restaurant.name}</h1>
              <p className="rest-sub">{restaurant.address}</p>
              <span className="direct-badge">Direct prices · no platform markup</span>
            </div>
            <MenuView menu={restaurant.menu ?? []} />
          </>
        ) : null}
      </main>
    </>
  );
}

export default function MenuPage() {
  return (
    <Suspense fallback={<main className="loading">Loading…</main>}>
      <MenuPageInner />
    </Suspense>
  );
}
