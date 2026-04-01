# ai-literacy-exemplar

[![Go Tests](https://github.com/russmiles/ai-literacy-exemplar/actions/workflows/go-tests.yml/badge.svg)](https://github.com/russmiles/ai-literacy-exemplar/actions/workflows/go-tests.yml)
[![Lint Markdown](https://github.com/russmiles/ai-literacy-exemplar/actions/workflows/lint-markdown.yml/badge.svg)](https://github.com/russmiles/ai-literacy-exemplar/actions/workflows/lint-markdown.yml)
[![Harness](https://img.shields.io/badge/Harness-7%2F7_enforced-2E8B57?style=flat-square)](HARNESS.md)
[![Harness Health](https://img.shields.io/badge/Harness_Health-Healthy-2E8B57?style=flat-square)](observability/snapshots/2026-04-01-snapshot.md)
[![Mutation Testing](https://img.shields.io/badge/Mutation_Testing-weekly-4682B4?style=flat-square)](HARNESS.md)
[![Coverage](https://img.shields.io/badge/Coverage-96%25-2E8B57?style=flat-square)](HARNESS.md)
[![AI Literacy](https://img.shields.io/badge/AI_Literacy-Level_3-20B2AA?style=flat-square)](assessments/2026-03-31-assessment.md)

A worked example of the AI Literacy framework — a Go CLI Markdown link checker built using the [ai-literacy-superpowers](https://github.com/russmiles/ai-literacy-superpowers) plugin, with the full framework-aligned development workflow in action.

---

## The Application: mdcheck

A command-line tool that checks Markdown files for broken links.

```bash
go build ./cmd/mdcheck/
./mdcheck README.md
```

Given a file or directory, it extracts all links (inline, reference, autolinks), verifies each one (HTTP HEAD for URLs, file existence for local paths), validates fragment links (`file.md#heading`) against actual headings, and reports any that are broken with file, line, and reason. Exits non-zero if broken links are found.

That's it. The application is deliberately simple — three source files, two test files, readable in ten minutes. The point is not the application. The point is how it was built.

---

## The Framework in Action

This repository demonstrates the AI Literacy framework's development workflow at every level.

### Level 2: Verification Infrastructure

CI workflows enforce quality on every push:

- **markdownlint** — all Markdown files pass lint
- **Go tests** — 18 tests, all passing
- **Coverage** — 96% on parser and checker (85% threshold enforced)
- **govulncheck** — no known vulnerabilities in dependencies
- **Mutation testing** — weekly via go-mutesting (investigative loop)

### Level 3: The Habitat

The development environment is fully configured:

- **CLAUDE.md** — literate programming, CUPID review, spec-first workflow, TDD
- **HARNESS.md** — 7 constraints, all enforced (4 deterministic, 2 agent, 1 weekly)
- **AGENTS.md** — compound learning memory with architectural decisions
- **MODEL_ROUTING.md** — model-tier guidance for the seven-agent team
- **Five project-local skills** — literate programming, CUPID, supply chain, dependency audit, AI literacy assessment

### Level 4: Specification-Driven Development

The link checker was specified before it was built:

- **[spec.md](specs/001-link-checker/spec.md)** — user stories, acceptance scenarios, functional requirements
- **[plan.md](specs/001-link-checker/plan.md)** — Go implementation plan with module structure, data model, algorithms

### Level 5: The Agent Team

Seven agents coordinate the development lifecycle:

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

### Mechanism Map

```text
ADVISORY LOOP (edit time — warn, do not block)
│
├── Hooks (from ai-literacy-superpowers plugin)
│   ├── PreToolUse constraint gate     Reads HARNESS.md commit-scoped constraints,
│   │                                  warns on violations during Write/Edit
│   ├── Stop drift check               Detects CI/linter/dependency changes at
│   │                                   session end, nudges /harness-audit
│   └── Stop reflection prompt          Detects commits during session,
│                                        nudges /reflect to capture learnings
├── Context (read by agents at session start)
│   ├── CLAUDE.md                       Go conventions, LP, CUPID, spec-first, TDD
│   ├── AGENTS.md                       Compound learning memory
│   ├── MODEL_ROUTING.md                Model-tier guidance for 7 agents
│   └── .claude/skills/
│       ├── literate-programming/       Code generation conventions
│       ├── cupid-code-review/          Code review lens
│       ├── github-actions-supply-chain/ CI hardening checklist
│       ├── dependency-vulnerability-audit/ Go CVE procedures
│       └── ai-literacy-assessment/     AI literacy scoring and remediation
│
└── Commands
    ├── /reflect                        Capture post-task learnings
    ├── /worktree spin|merge|clean      Parallel agent isolation
    └── /assess                         AI literacy assessment + remediation


STRICT LOOP (merge time — block until green)
│
├── CI Workflows (.github/workflows/)
│   ├── lint-markdown.yml               markdownlint on all .md files
│   └── go-tests.yml                    Go tests + 96% coverage + govulncheck
│
├── Agent Pipeline (.claude/agents/)
│   ├── orchestrator                    Coordinates pipeline
│   │   ├── GATE: plan approval         User reviews spec before implementation
│   │   └── GUARDRAIL: MAX_REVIEW_CYCLES=3
│   ├── spec-writer                     Spec + plan updates (no Bash)
│   ├── tdd-agent                       Failing tests from spec scenarios
│   ├── go-implementer                  Makes Go tests green (scoped)
│   ├── code-reviewer                   CUPID + LP review (no Write)
│   ├── integration-agent               CHANGELOG, PR, CI, merge, reflection
│   └── assessor                       AI literacy assessment + remediation
│
└── Harness Constraints (HARNESS.md)
    ├── 4 deterministic                 markdownlint, tests, coverage, govulncheck
    ├── 2 agent-backed                  Literate programming, CUPID
    └── 1 weekly deterministic          Mutation testing


INVESTIGATIVE LOOP (scheduled — sweep for entropy)
│
├── Mutation Testing (mutation-testing.yml, weekly)
│   └── Go — go-mutesting (parser + checker)
│
├── Dependabot (.github/dependabot.yml)
│   ├── github-actions                  Weekly action updates
│   └── gomod                           Weekly Go dependency updates
│
├── Garbage Collection Rules (HARNESS.md)
│   ├── Documentation freshness         Agent — stale references
│   ├── Convention drift                Agent — LP preambles, CUPID naming
│   ├── Stale AGENTS.md                 Agent — unreviewed reflection entries
│   ├── Snapshot staleness              Deterministic — file date check
│   └── Dependency currency             Deterministic — govulncheck + Dependabot
│
└── Compound Learning
    ├── REFLECTION_LOG.md               Agent reflections (append-only)
    └── AGENTS.md                       Human-curated from reflections
```

---

## Verification Approach

Three layers form the verification chain:

1. **Tests verify behaviour** — 18 tests derived from spec acceptance scenarios
2. **Coverage verifies execution** — 96% of statements in parser and checker
3. **Mutation testing verifies the tests** — weekly go-mutesting confirms tests detect faults

Coverage measures what was executed. Mutation testing measures whether the tests actually detect changes.

## Observability

The three enforcement loops generate signals that make the collaboration observable. Without observability, cost discipline is aspirational and verification metrics are snapshots rather than trends.

| Panel | What to track | Source in this repo |
| ----- | ------------- | ------------------- |
| **Cost** | Spend trend, model-tier distribution | Claude Code analytics dashboard |
| **Quality** | Coverage trend (96%), mutation score trend | CI artifacts from go-tests.yml and mutation-testing.yml |
| **Adoption** | Sessions per developer, acceptance rate | Provider analytics |
| **Habitat health** | Harness 7/7 enforced, REFLECTION_LOG.md growth | /harness-status, git log |

The METR study found developers perceive a 20% AI speedup but measure a 19% slowdown. Observability closes this gap — measure first, believe second.

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
