# Code Review Guidelines

Mirrors the "Code Review Guidelines" page in the StreamCraft Notion project — keep both in sync if this changes.

## Correctness
- Does the code do what the linked backlog item describes? Any silently missing edge cases?
- Go errors checked and wrapped with context (`fmt.Errorf("...: %w", err)`), not swallowed.
- Concurrency: goroutines have a clear owner/lifecycle; context cancellation respected in long-running work (especially ffmpeg jobs).

## Security (real attack surface here)
- **Uploads**: file type/size validated server-side; never trust client-supplied MIME type alone.
- **ffmpeg invocation**: arguments passed as a slice to `exec.Command`, never string-concatenated into a shell — avoids command injection.
- **Auth**: passwords hashed (bcrypt/argon2); sessions/tokens expire; admin-only endpoints check role server-side.
- **Streaming URLs**: signed/expiring URLs or auth-gated proxying so video isn't trivially hot-linkable.
- Secrets never hardcoded or logged.

## Go idiom & style
- Standard Go project layout; package names short, not stutter-y.
- Exported identifiers documented where non-obvious.
- `go vet` / `staticcheck` clean; `gofmt`-formatted.

## Data model
- Schema changes come with a migration, reversible where practical.
- Watch-history writes are idempotent/safe under retries.

## Testing
- Core business logic (progress calc, auth checks, upload validation) has unit tests.
- ffmpeg-invoking code is structured so command-building logic is testable without running ffmpeg.

## Infra (OpenTofu)
- No hardcoded account IDs/ARNs.
- IAM policies scoped to what the resource needs, not `*`.
- Remote state + locking before this goes beyond a single developer.

## Definition of done
1. Reviewed against this checklist
2. Tests pass and exist for new logic
3. Backlog status updated in Notion
4. ADR written if a real tradeoff decision was involved
