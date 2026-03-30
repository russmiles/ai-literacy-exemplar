# AGENTS.md

<!-- Compound learning memory. Human-curated, agent-proposed. -->

## STYLE

- Go standard library preferred over third-party packages
- Literate programming preambles on every .go file
- CUPID review on every change

## GOTCHAS

<!-- Populated from REFLECTION_LOG.md -->

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
