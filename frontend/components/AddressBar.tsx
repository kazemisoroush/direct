// AddressBar shows the delivery address at the top of the page. Changing it is a later
// milestone; for now it displays the selected address (the whole list depends on it).
export function AddressBar({ address }: { address: string }) {
  return (
    <button className="address" type="button" aria-label="Delivery address">
      <svg className="pin" width="15" height="15" viewBox="0 0 24 24" fill="none" aria-hidden="true">
        <path
          d="M12 2C8.1 2 5 5 5 8.8c0 4.9 7 13.2 7 13.2s7-8.3 7-13.2C19 5 15.9 2 12 2z"
          fill="currentColor"
        />
        <circle cx="12" cy="9" r="2.5" fill="var(--surface)" />
      </svg>
      <span className="who">Deliver to</span>
      <span className="val">{address}</span>
      <svg className="caret" width="12" height="12" viewBox="0 0 24 24" fill="none" aria-hidden="true">
        <path d="M6 9l6 6 6-6" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" />
      </svg>
    </button>
  );
}
