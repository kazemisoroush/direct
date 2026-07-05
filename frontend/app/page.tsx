"use client";

import { useCallback, useEffect, useState } from "react";
import { useRouter } from "next/navigation";

import { AddressBar } from "../components/AddressBar";
import { RestaurantList } from "../components/RestaurantList";
import { ThemeToggle } from "../components/ThemeToggle";
import { Wordmark } from "../components/Wordmark";
import { useAuth } from "../lib/auth/context";
import { listRestaurants } from "../lib/restaurants/listRestaurants";
import type { Restaurant } from "../lib/restaurants/restaurant";

// The selected delivery address. M1 fixes it to mine; choosing an address is a later milestone.
const deliveryAddress = "Kellyville NSW 2155";

export default function Home() {
  const { ready, authenticated, api, signOut } = useAuth();
  const router = useRouter();
  const [restaurants, setRestaurants] = useState<Restaurant[] | null>(null);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    if (ready && !authenticated) {
      router.replace("/login");
    }
  }, [ready, authenticated, router]);

  const refresh = useCallback(async () => {
    if (!api) return;
    setError(null);
    try {
      setRestaurants(await listRestaurants(api, deliveryAddress));
    } catch (err) {
      setError(err instanceof Error ? err.message : "could not load restaurants");
    }
  }, [api]);

  useEffect(() => {
    if (authenticated) {
      void refresh();
    }
  }, [authenticated, refresh]);

  if (!ready) {
    return <main className="loading">Loading…</main>;
  }
  if (!authenticated) {
    return null;
  }

  return (
    <>
      <header className="topbar">
        <Wordmark />
        <div className="topbar-actions">
          <AddressBar address={deliveryAddress} />
          <ThemeToggle />
          <button className="link" onClick={signOut}>
            Sign out
          </button>
        </div>
      </header>
      <main className="page">
        <div className="page-head">
          <h1>Order straight from the kitchen</h1>
          <p>The restaurant&rsquo;s real price plus real delivery. No 35% platform tax.</p>
        </div>
        <div className="rule">Delivering to you</div>
        {error && <p role="alert">{error}</p>}
        {restaurants === null ? (
          <p className="loading">Finding restaurants near you…</p>
        ) : (
          <RestaurantList restaurants={restaurants} />
        )}
      </main>
    </>
  );
}
