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

---

## Garbage Collection

### Documentation freshness

- **What it checks**: Docs referencing files or functions that no
  longer exist
- **Frequency**: weekly
- **Enforcement**: agent
- **Tool**: harness-gc agent
- **Auto-fix**: false

---

## Status

Last audit: 2026-03-30
Constraints enforced: 7/7
Garbage collection active: 1/1
Drift detected: no
