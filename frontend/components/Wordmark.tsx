// Wordmark is the Direct logotype: a pin-with-plate mark (local + food) plus an accent dot.
export function Wordmark() {
  return (
    <span className="wordmark">
      <svg className="mark" width="22" height="22" viewBox="0 0 24 24" fill="none" aria-hidden="true">
        <path
          d="M12 2C8.1 2 5 5 5 8.8c0 4.9 7 13.2 7 13.2s7-8.3 7-13.2C19 5 15.9 2 12 2z"
          fill="currentColor"
        />
        <circle cx="12" cy="9" r="2.7" fill="var(--surface)" />
      </svg>
      Direct<span className="dot">.</span>
    </span>
  );
}
