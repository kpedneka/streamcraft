.PHONY: build test lint run migrate hooks licenses

COVERAGE_THRESHOLD := 80

all: fmt lint build test

fmt:
	gofmt -s -w .

build:
	go build -o ./bin/api ./cmd/api

test:
	go test ./...

lint:
	golangci-lint run ./...

# Installs the pre-commit hooks defined in lefthook.yml (gitleaks, gofmt,
# go vet) into this clone's .git/hooks. One-time, per clone — lefthook.yml
# itself is checked in, but git hooks live outside version control, so
# every clone needs this run once. Requires lefthook (`brew install
# lefthook`).
hooks:
	lefthook install

# Regenerates THIRD_PARTY_LICENSES.csv: every Go dependency reachable from
# any cmd/ entrypoint, plus the license type go-licenses detected for it.
# Filters out the project's own module — go-licenses reports it too (now
# that a real LICENSE file exists to detect), but this repo's own license
# isn't a third-party notice. Requires go-licenses
# (`go install github.com/google/go-licenses/v2@latest`). Not run in CI —
# dependency licenses don't change often enough to justify a check on
# every push; rerun by hand after adding a new dependency.
licenses:
	go-licenses report ./cmd/... 2>/dev/null | grep -v '^github.com/kpedneka/streamcraft,' > THIRD_PARTY_LICENSES.csv