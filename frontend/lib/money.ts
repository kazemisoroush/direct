// formatPrice renders integer cents as an AUD price string, e.g. 1600 → "$16.00".
export function formatPrice(cents: number): string {
  return new Intl.NumberFormat("en-AU", { style: "currency", currency: "AUD" }).format(cents / 100);
}
