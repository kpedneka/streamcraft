# ADR-001: Monorepo vs. polyrepo for backend + frontend

**Status:** Accepted
**Date:** 2026-09-19

## Context

StreamCraft has two independently-versioned stacks — a Go backend (`backend/`) and
a Next.js frontend (`web/`) — plus project docs (`docs/`). Repo scaffolding needed
to decide whether these live in one repository or two before module structure,
linting, and CI could be set up.

This is a single-developer learning project, not a project with separate backend
and frontend teams. Streaming features (e.g. auth, playback, upload) commonly
touch both the API and the UI in the same unit of work, and CI/tooling setup
needs to be done once, not duplicated across repos.

## Decision

Monorepo. `backend/`, `web/`, and `docs/` all live under one repo root
(this repository). Each stack keeps its own dependency manifest
(`backend/go.mod`, `web/package.json`) and its own CI job in
`.github/workflows/ci.yml`, scoped via `working-directory`, so the two stacks
build, test, and lint independently despite sharing a repo.

A root-level `package.json` exists only for repo-wide tooling (commit message
linting via commitlint) — it is not an app and must not accumulate
frontend/backend dependencies.

## Alternatives considered

- **Polyrepo** (separate `streamcraft-backend` and `streamcraft-web`
  repositories): rejected. For a solo developer, cross-repo coordination
  (matching PRs, pinning one repo's commit in the other, duplicated CI/lint
  setup) adds overhead with no corresponding benefit — there's no team
  boundary or independent release cadence to protect. A feature that changes
  an API response shape and its frontend consumer would need two PRs merged
  in the right order instead of one atomic commit.

## Consequences

- A single PR can span backend and frontend changes atomically, which fits
  how most features here will actually be built.
- CI must scope each job's install/build/test/lint to its own subdirectory
  (`working-directory: backend` / `working-directory: web`) rather than
  running one blanket command at the repo root — already how
  `.github/workflows/ci.yml` is structured.
- The root `package.json`/`package-lock.json` must stay limited to repo-root
  tooling (commitlint). Frontend test/build dependencies belong in
  `web/package.json` only — this was previously gotten wrong (vitest and its
  peers were briefly added at the repo root) and broke `web`'s CI job, which
  installs via `npm ci` scoped to `web/` and won't see root-level packages.
- If the project ever grows beyond one developer with genuinely independent
  backend/frontend release cycles, this decision would be worth revisiting —
  not expected during the current learning-project phase.
