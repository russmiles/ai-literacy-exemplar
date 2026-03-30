# ai-literacy-exemplar

[![Go Tests](https://github.com/russmiles/ai-literacy-exemplar/actions/workflows/go-tests.yml/badge.svg)](https://github.com/russmiles/ai-literacy-exemplar/actions/workflows/go-tests.yml)
[![Lint Markdown](https://github.com/russmiles/ai-literacy-exemplar/actions/workflows/lint-markdown.yml/badge.svg)](https://github.com/russmiles/ai-literacy-exemplar/actions/workflows/lint-markdown.yml)
[![Harness](https://img.shields.io/badge/Harness-7%2F7_enforced-2E8B57?style=flat-square)](HARNESS.md)
[![Mutation Testing](https://img.shields.io/badge/Mutation_Testing-weekly-4682B4?style=flat-square)](HARNESS.md)
[![Coverage](https://img.shields.io/badge/Coverage-96%25-2E8B57?style=flat-square)](HARNESS.md)

A worked example of the AI Literacy framework — a Go CLI Markdown link checker built using the [ai-literacy-superpowers](https://github.com/russmiles/ai-literacy-superpowers) plugin, with the full framework-aligned development workflow in action.

---

## The Application: mdcheck

A command-line tool that checks Markdown files for broken links.

```bash
go build ./cmd/mdcheck/
./mdcheck README.md
```

Given a file or directory, it extracts all links (inline, reference, autolinks), verifies each one (HTTP HEAD for URLs, file existence for local paths), and reports any that are broken with file, line, and reason. Exits non-zero if broken links are found.

That's it. The application is deliberately simple — three source files, two test files, readable in ten minutes. The point is not the application. The point is how it was built.

---

## The Framework in Action

This repository demonstrates the AI Literacy framework's development workflow at every level.

### Level 2: Verification Infrastructure

CI workflows enforce quality on every push:

- **markdownlint** — all Markdown files pass lint
- **Go tests** — 15 tests, all passing
- **Coverage** — 96% on parser and checker (85% threshold enforced)
- **govulncheck** — no known vulnerabilities in dependencies
- **Mutation testing** — weekly via go-mutesting (investigative loop)

### Level 3: The Habitat

The development environment is fully configured:

- **CLAUDE.md** — literate programming, CUPID review, spec-first workflow, TDD
- **HARNESS.md** — 7 constraints, all enforced (4 deterministic, 2 agent, 1 weekly)
- **AGENTS.md** — compound learning memory with architectural decisions
- **MODEL_ROUTING.md** — model-tier guidance for the six-agent team
- **4 project-local skills** — literate programming, CUPID, supply chain, dependency audit

### Level 4: Specification-Driven Development

The link checker was specified before it was built:

- **[spec.md](specs/001-link-checker/spec.md)** — user stories, acceptance scenarios, functional requirements
- **[plan.md](specs/001-link-checker/plan.md)** — Go implementation plan with module structure, data model, algorithms

### Level 5: The Agent Team

Six agents coordinate the development lifecycle:

```text
orchestrator
  → spec-writer
  → GATE: plan approval (user reviews before proceeding)
  → tdd-agent
  → go-implementer
  → code-reviewer (CUPID + literate programming)
  → GUARDRAIL: MAX_REVIEW_CYCLES=3
  → integration-agent (includes reflection → REFLECTION_LOG.md)
```

---

## The Three Enforcement Loops

```text
ADVISORY (edit time)
  └── PreToolUse hook warns on HARNESS.md constraint violations

STRICT (merge time)
  ├── markdownlint on all .md files
  ├── Go tests + 96% coverage
  └── govulncheck vulnerability scan

INVESTIGATIVE (weekly)
  ├── go-mutesting mutation testing
  └── Dependabot dependency updates
```

---

## Verification Approach

Three layers form the verification chain:

1. **Tests verify behaviour** — 15 tests derived from spec acceptance scenarios
2. **Coverage verifies execution** — 96% of statements in parser and checker
3. **Mutation testing verifies the tests** — weekly go-mutesting confirms tests detect faults

Coverage measures what was executed. Mutation testing measures whether the tests actually detect changes.

---

## How to Study This Repo

Suggested reading order:

1. **This README** — the big picture
2. **[spec.md](specs/001-link-checker/spec.md)** — what the application does
3. **[plan.md](specs/001-link-checker/plan.md)** — how it's implemented
4. **[CLAUDE.md](CLAUDE.md)** — the conventions
5. **[HARNESS.md](HARNESS.md)** — the enforcement
6. **[internal/parser.go](internal/parser.go)** — literate programming in practice
7. **[internal/checker.go](internal/checker.go)** — CUPID properties in practice
8. **[.claude/agents/](/.claude/agents/)** — the agent team

---

## How to Apply to Your Own Project

Install the [ai-literacy-superpowers](https://github.com/russmiles/ai-literacy-superpowers) plugin:

```bash
claude plugin install ai-literacy-superpowers
```

Then initialize:

```text
cd your-project
/superpowers-init
```

The init command discovers your stack, asks about your conventions, and scaffolds the full habitat.

---

## Intellectual Foundations

This exemplar draws on:

- **Christopher Alexander** — the quality without a name; design for inhabitants
- **Donald Knuth** — literate programming; code as literature for readers
- **Richard P. Gabriel** — habitability; code as a place to live in
- **Daniel Terhorst-North** — CUPID; code as a place of joy
- **Birgitta Boeckeler** — harness engineering; the three enforcement loops
- **Addy Osmani** — agent orchestration; quality gates and compound learning

The mission: building habitats where human and AI intelligence thrive together.
