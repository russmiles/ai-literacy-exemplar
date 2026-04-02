# Harness — mdcheck

## Context

### Stack

- **Primary language**: Go 1.26
- **Build system**: Go modules
- **Test framework**: Go testing package
- **Container strategy**: none (CLI tool)

### Conventions

- **Naming**: Go standard — PascalCase exported, camelCase unexported
- **File structure**: cmd/ for entry points, internal/ for library code
- **Error handling**: wrap with fmt.Errorf, return to caller
- **Documentation**: literate programming preambles on every file

---

## Constraints

### Markdown formatting

- **Rule**: All Markdown files must pass markdownlint
- **Enforcement**: deterministic
- **Tool**: npx markdownlint-cli2 "**/*.md"
- **Scope**: pr

### Go tests pass

- **Rule**: All tests must pass with zero failures
- **Enforcement**: deterministic
- **Tool**: go test ./...
- **Scope**: pr

### Go test coverage

- **Rule**: Code coverage on parser.go and checker.go must be at
  least 85%
- **Enforcement**: deterministic
- **Tool**: Go coverage tool
- **Scope**: pr

### Go vulnerability scan

- **Rule**: No known vulnerabilities in dependencies
- **Enforcement**: deterministic
- **Tool**: govulncheck ./...
- **Scope**: pr

### Literate programming compliance

- **Rule**: Every source file has a narrative preamble, documentation
  explains reasoning, presentation follows logical understanding
- **Enforcement**: agent
- **Tool**: harness-enforcer
- **Scope**: pr

### CUPID properties

- **Rule**: Code maintains composable, unix, predictable, idiomatic,
  and domain-based properties
- **Enforcement**: agent
- **Tool**: harness-enforcer
- **Scope**: pr

### Mutation testing score tracked

- **Rule**: Weekly mutation testing via go-mutesting; score reported
  as CI artifact
- **Enforcement**: deterministic
- **Tool**: Weekly mutation-testing.yml workflow
- **Scope**: weekly

### SBOM generation

- **Rule**: A Software Bill of Materials is generated for every release
  binary, listing all dependencies and their licenses
- **Enforcement**: unverified
- **Tool**: `go version -m` or syft/cyclonedx-gomod (activate when
  first external dependency is added)
- **Scope**: pr

---

## Garbage Collection

### Documentation freshness

- **What it checks**: Docs referencing files or functions that no
  longer exist
- **Frequency**: weekly
- **Enforcement**: agent
- **Tool**: harness-gc agent
- **Auto-fix**: false

### Convention drift

- **What it checks**: Whether source files follow literate programming
  preambles and CUPID naming conventions
- **Frequency**: weekly
- **Enforcement**: agent
- **Tool**: harness-gc agent
- **Auto-fix**: false

### Stale AGENTS.md

- **What it checks**: Whether REFLECTION_LOG.md entries older than 30
  days have been reviewed for promotion to AGENTS.md
- **Frequency**: weekly
- **Enforcement**: agent
- **Tool**: harness-gc agent
- **Auto-fix**: false

### Snapshot staleness

- **What it checks**: Whether the most recent harness health snapshot
  in observability/snapshots/ is less than 30 days old
- **Frequency**: weekly
- **Enforcement**: deterministic
- **Tool**: file date check
- **Auto-fix**: false

### Dependency currency

- **What it checks**: Whether project dependencies have known
  vulnerabilities or are outdated
- **Frequency**: weekly
- **Enforcement**: deterministic
- **Tool**: govulncheck + Dependabot
- **Auto-fix**: false

---

## Status

Last audit: 2026-04-01
Constraints enforced: 7/8
Garbage collection active: 5/5
Drift detected: no
