# Direct

Order Uber-style **directly from restaurants**, skipping the ~35% delivery-platform tax
(commission + delivery fee + service fee). The customer pays the restaurant's real price
plus real delivery — nothing else.

Planning, constraints and milestones live in Linear (personal workspace, team Soroush):
project **Direct**. This repo is the implementation.

## Steel thread

One user (me), one restaurant (Hills Kebabs Kellyville): log in → one page with my address
on top → restaurants that deliver to my address → menu → cart → checkout → place order
(no payment). Complete at milestone **M3**.

## Inside the Box (closed world)

Go Lambda + API Gateway (HTTP API), DynamoDB, Amazon Cognito, a Next.js static site on
S3 + CloudFront, all provisioned by AWS CDK (Go) and deployed by GitHub Actions via OIDC —
the same box proven in the Vault project. One LLM vision API is admitted at M4 for AI menu
reading. Payments, multi-restaurant onboarding and delivery logistics stay out of the box.

## Layout

- `backend/` — Go API (`cmd/lambda` for AWS, `cmd/server` for local), HTTP handlers in `internal/api`.
- `infra/` — AWS CDK app: `DirectStack` (frontend hosting + API) and `DirectCICDStack` (GitHub OIDC deploy role).
- `frontend/` — Next.js static-export SPA.
- `openapi.yaml` — the API contract; source of truth for backend and frontend.

## Status

- **M0** — baseline: static page over CloudFront + a `GET /health` Lambda, deployed to AWS on merge to `main`, CI on PRs.
- **M1** — walking skeleton: Amazon Cognito auth and the one authenticated page (delivery address on top + a data-driven restaurant list). `GET /health` stays public; every other route requires a Cognito access token.
- **M2** — restaurant CRUD is API-first: `POST /restaurants` creates a restaurant with its menu, `GET /restaurants/{id}` returns it, and the home card opens a menu page. The address→restaurant filter shows a restaurant only when it delivers to the selected postcode. The catalogue is populated by calling the API (no committed seed data or CLI).

Cart, checkout and placing an order land in M3.

## Commands

```sh
make build   # build backend and infra
make test    # run backend tests
make lint    # lint backend
make synth   # cdk synth (runs the cdk-nag security gate)
make deploy  # cdk deploy
```

Frontend: `cd frontend && npm ci && npm run dev` (dev), `npm run build` (static export to `out/`).
