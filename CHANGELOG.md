# Changelog

---

## 30 March 2026

### Initial habitat and application

- **Project scaffolded with ai-literacy-superpowers** — CLAUDE.md,
  HARNESS.md (7 constraints, all enforced), AGENTS.md, MODEL_ROUTING.md,
  REFLECTION_LOG.md, six-agent team, four skills, CI workflows,
  Dependabot configuration, and weekly mutation testing.
- **mdcheck link checker implemented** — Go CLI that parses Markdown
  files for inline links, reference links, and autolinks, then checks
  each via HTTP HEAD (URLs) or file existence (local paths). Reports
  broken links with file, line, text, and reason. Spec-first workflow
  with TDD — 15 tests, 96% coverage on parser and checker.
