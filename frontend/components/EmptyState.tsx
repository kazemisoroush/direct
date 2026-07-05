// EmptyState is shown when no restaurants deliver to the address yet. Honest for M1:
// nothing is seeded, and it names what's coming next without pretending otherwise.
export function EmptyState() {
  return (
    <div className="empty">
      <svg className="glyph" width="40" height="40" viewBox="0 0 48 48" fill="none" aria-hidden="true">
        <circle cx="24" cy="24" r="15" stroke="currentColor" strokeWidth="2.2" />
        <circle cx="24" cy="24" r="7" stroke="currentColor" strokeWidth="2.2" />
        <path d="M8 12v10M8 12c2 0 3 1.5 3 4s-1 4-3 4" stroke="currentColor" strokeWidth="2.2" strokeLinecap="round" />
      </svg>
      <h2>No kitchens here yet</h2>
      <p>
        We&rsquo;re onboarding your first local restaurant. Hills Kebabs lands next &mdash; you&rsquo;ll order
        straight from them.
      </p>
      <p className="save-note">That&rsquo;s the ~35% delivery-app tax, gone.</p>
    </div>
  );
}
