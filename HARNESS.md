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

### Human review before merge

- **Rule**: Every PR must have at least one approving human review
  before merge
- **Enforcement**: deterministic
- **Tool**: GitHub branch protection (required reviews) + agent
  pre-merge check
- **Scope**: pr
- **Governance requirement**: Human review of all AI-generated code
  before merge
- **Operational meaning**: GitHub branch protection requires at least
  one approving review. The agent also checks for an existing approval
  before attempting any merge operation and warns if none exists.
- **Verification method**: deterministic — branch protection enforces
  at merge time; agent verifies before invoking merge
- **Evidence**: GitHub PR approval record (reviewer name, timestamp,
  approval status)
- **Failure action**: merge blocked by branch protection; agent warns
  and does not attempt merge without approval
- **Frame check**: confirmed aligned — engineering (approval required),
  compliance (audit trail of reviewer identity and timestamp),
  AI system (agent checks before merge, branch protection enforces)

### No secrets in source — prevention

- **Rule**: A pre-commit hook prevents committing files matching
  known secret patterns (API keys, tokens, private keys, .env files)
- **Enforcement**: unverified
- **Tool**: gitleaks pre-commit hook (to be configured)
- **Scope**: commit
- **Governance requirement**: No secrets committed to source control
- **Operational meaning**: A gitleaks pre-commit hook blocks commits
  containing secret patterns before they enter git history
- **Verification method**: deterministic locally; unverified from CI
  (cannot confirm hook is installed on every machine)
- **Evidence**: pre-commit hook configuration exists in the repository
- **Failure action**: commit blocked locally by the hook
- **Frame check**: confirmed aligned — engineering (hook blocks
  secrets at commit time), compliance (hook config documented in
  repo), AI system (agent never commits files matching secret
  patterns)

### No secrets in source — detection

- **Rule**: A CI gitleaks scan finds zero secret patterns in every PR
- **Enforcement**: unverified
- **Tool**: gitleaks CI step (to be added to go-tests.yml)
- **Scope**: pr
- **Governance requirement**: No secrets committed to source control
- **Operational meaning**: gitleaks scans the PR diff for secret
  patterns; any finding fails the check
- **Verification method**: deterministic — gitleaks in CI workflow
- **Evidence**: gitleaks CI step passes with zero findings
- **Failure action**: PR blocked — merge cannot proceed until secrets
  are removed and history cleaned
- **Frame check**: confirmed aligned — engineering (CI scan catches
  secrets missed by hooks), compliance (every merged PR has passing
  scan in audit trail), AI system (CI gate independently verifies)

### Approved dependency licenses

- **Rule**: Every external dependency must use a permissive
  OSI-approved license from the allowlist: MIT, Apache-2.0,
  BSD-2-Clause, BSD-3-Clause. Copyleft licenses (GPL, AGPL, LGPL)
  are prohibited.
- **Enforcement**: unverified
- **Tool**: go-licenses or licensei (activate when first external
  dependency is added)
- **Scope**: pr
- **Governance requirement**: All dependencies must have approved
  open-source licenses
- **Operational meaning**: Currently stdlib-only; constraint activates
  with the first external dependency. When active, a CI license check
  verifies every dependency against the allowlist.
- **Verification method**: deterministic when activated; unverified
  until first external dependency
- **Evidence**: When activated, CI license check output. Until then,
  go.mod showing zero external dependencies.
- **Failure action**: When activated, PR blocked. Until then, any PR
  adding an external dependency triggers manual license review.
- **Frame check**: confirmed aligned — engineering (allowlist enforced
  by CI tool), compliance (license audit report per PR), AI system
  (agent checks license compatibility before recommending dependencies)

### AI change traceability

- **Rule**: Every AI-assisted PR must have: (1) identifiable commits
  via branch naming or commit metadata, (2) a PR description section
  stating what was AI-generated, (3) a linked session record (design
  spec, plan doc, or reflection log entry)
- **Enforcement**: agent
- **Tool**: harness-enforcer
- **Scope**: pr
- **Governance requirement**: AI-generated changes must be
  attributable — traceability from change to AI session
- **Operational meaning**: The agent reviews PRs for three traceability
  elements: commit-level identification, PR-level attribution, and
  session-level records. Flags gaps as advisory warnings.
- **Verification method**: agent — harness-enforcer checks PRs for
  attribution metadata, AI-session markers, and linked session records
- **Evidence**: Branch names or commit messages indicating AI-assisted
  work, PR description with AI attribution, linked design spec/plan/
  reflection log entry
- **Failure action**: advisory — agent flags missing traceability,
  merge proceeds after human review confirms attribution is adequate
- **Frame check**: confirmed aligned — engineering (commits, PRs, and
  docs linked to AI sessions), compliance (audit trail queryable via
  git history and docs/observability), AI system (agent checks own
  output for traceability before declaring work complete)

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

## Observability

- Snapshot cadence: weekly

---

## Status

Last audit: 2026-04-14
Constraints enforced: 7/8 (technical) + 2/5 (governance)
Governance constraints: 5 (1 deterministic, 1 agent, 3 unverified)
Garbage collection active: 5/5
Drift detected: no
