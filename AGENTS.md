# AGENTS.md

<!-- Compound learning memory. Human-curated, agent-proposed. -->

## STYLE

- Go standard library preferred over third-party packages
- Literate programming preambles on every .go file
- CUPID review on every change

## GOTCHAS

<!-- Populated from REFLECTION_LOG.md -->

GOTCHA: configured-vs-operational — A fully configured habitat (HARNESS.md,
agents, skills, hooks) does not guarantee operational maturity. The habitat
must be exercised through routine practice: regular audits, reflection capture
and promotion, pipeline usage across features, and cadence-driven health
checks. Configuration is Level 3; operation is Level 4+.

GOTCHA: L4-to-L5-requires-elapsed-time — The gap between Level 4 and
Level 5 is not infrastructure — it is elapsed time and organisational
scope. L5 cannot be sprinted to. It requires sustained quarterly practice,
actual cost data, and tooling that serves teams beyond the originating
repo. Do not plan L5 work as a sprint-level task.

GOTCHA: two-speed-learning-pipeline — The compound learning pipeline has
two speeds: agent-speed (reflections appear immediately) and human-speed
(promotions require review). If reflections accumulate without promotion
reviews, the pipeline produces signal that never becomes institutional
memory. The /harness-health nudge for unpromoted reflections exists to
close this gap — act on it.

## ARCH_DECISIONS

- Parser extracts links from Markdown without a full AST — regex
  plus state machine for inline/reference/autolink patterns
- Checker uses HTTP HEAD for URLs (not GET) to minimise bandwidth
- Exit code 1 if any broken links found; exit code 0 if all clean
- Fragment-only links (#heading) are skipped — validating them would
  require a heading parser that exceeds this tool's scope

## TEST_STRATEGY

- Spec-first: acceptance scenarios in spec.md drive test design
- TDD strictly: red-green-refactor
- Coverage: 85% on parser.go and checker.go (currently 96%)
- Mutation testing: weekly via go-mutesting
- HTTP tests use httptest.NewServer for deterministic URL checking
- File tests use t.TempDir for isolated filesystem checks
- When existing tests pass for new feature scenarios, the failing test
  is the one that matters — it defines the boundary between what exists
  and what's new. Mixed results during feature work are a healthy sign
  that the existing design was close to correct.
