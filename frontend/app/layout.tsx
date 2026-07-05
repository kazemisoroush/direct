import type { Metadata } from "next";

import "./globals.css";

export const metadata: Metadata = {
  title: "Direct",
  description: "Order straight from the restaurant — no middleman tax",
};

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  );
}
