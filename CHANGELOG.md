# Changelog

---

## 31 March 2026 (Q2 update)

### Documentation

- **Improved README with five targeted gaps** — added a Try It Yourself
  exercise with agent pipeline and manual instructions for adding image
  link support; named all source files with one-line descriptions; showed
  illustrative example output with exit codes; documented the L3-to-L4
  assessment and remediation progression; explained the compound learning
  three-stage cycle with concrete REFLECTION_LOG and AGENTS.md state;
  clarified the distinction between the 7 constraints and 5 GC rules, and
  noted Dependabot's role in the investigative loop.

---

## 1 April 2026 (afternoon)

### AI Literacy Assessment

- **Second AI literacy assessment: Level 4 — Specification Architect** —
  all three L3 gaps addressed (compound learning active, cadence
  documented, pipeline exercised on fragment validation). Context
  engineering 4/5, architectural constraints 5/5, guardrail design 4/5.
  Remaining gaps to L5: no OTel, no cost data, cadence not yet proven
  across a full quarter.
- **README badge updated** — AI Literacy badge now shows Level 4 with
  link to the new assessment document.
- **Assessment reflection captured** — third REFLECTION_LOG entry
  documenting the L3-to-L4 progression and the qualitative difference
  between L4 and L5.

---

## 31 March 2026 (afternoon)

### New feature — fragment link validation

- **Added fragment link validation to mdcheck (FR-011)** — links like
  `file.md#heading` now verify that the heading exists in the target
  file. Headings are normalised to GitHub-compatible slugs. Missing
  headings are reported as broken with a clear reason message.
- **Spec, plan, and tests updated before implementation** — two new
  acceptance scenarios, FR-011 added to spec, plan updated with FR
  mapping and test cases, three failing tests written before any
  production code.
- **Coverage maintained at 96.1%** — all new functions (slugify,
  fileHasFragment, updated checkLocalFile) at 87-100% coverage.
- **Reflection captured** — spec-first workflow surfaced slug
  normalisation concern during scenario writing, preventing scope
  creep during implementation.

---

## 1 April 2026

### Level 5 infrastructure

- **Fixed README agent and skill counts** — agent count corrected from
  six to seven (assessor added), skill count corrected from four to five
  (ai-literacy-assessment skill added), mechanism map updated with all
  five GC rules and correct agent tier guidance.
- **Added quarterly operating cadence to CLAUDE.md** — documents the
  five recurring actions (assess, harness-audit, reflection review,
  mutation trend check, harness-health) needed to keep the habitat
  operational rather than merely configured.
- **Promoted first reflection to AGENTS.md GOTCHAS** — the
  configured-vs-operational insight from the 2026-03-31 assessment is
  now a permanent GOTCHA for future agents. Compound learning
  demonstrated for the first time.
- **Expanded HARNESS.md from 1 to 5 GC rules** — added convention
  drift, stale AGENTS.md, snapshot staleness, and dependency currency.
  All five rules have enforcement configured; status section updated to
  5/5 active, audit date 2026-04-01.
- **Replaced sample snapshot with accurate health data** — the
  2026-04-01 snapshot now reflects actual state: 7/7 constraints,
  5/5 GC rules, 1 reflection promoted, mutation testing deferred to
  CI artifacts rather than fabricated numbers.

---

## 31 March 2026

### AI Literacy Assessment

- **First AI literacy assessment: Level 3 — Habitat Engineer** — full
  evidence-based assessment documenting L2-L5 signals, three-discipline
  scoring (context 4/5, constraints 5/5, guardrails 3/5), strengths,
  gaps, and recommendations. Guardrail design ceiling identified:
  configured but not yet operational.
- **Assessment badge and mechanism map updated** — AI Literacy Level 3
  badge added to README, assessor agent and /assess command added to
  mechanism map.
- **First reflection captured** — REFLECTION_LOG.md populated with
  assessment reflection, activating the compound learning system.

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
