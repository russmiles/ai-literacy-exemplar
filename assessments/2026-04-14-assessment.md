# AI Literacy Assessment — ai-literacy-exemplar

**Date**: 2026-04-14
**Assessed by**: assessor agent (via /assess command)
**Assessed level**: Level 4 — Specification Architect
**Previous assessment**: [2026-04-01](2026-04-01-assessment.md) (Level 4)

---

## Observable Evidence

### Repository Signals

| Signal | Found | Level |
| --- | --- | --- |
| CI workflows | Yes — go-tests.yml, lint-markdown.yml, mutation-testing.yml (3) | L2 |
| Test coverage enforcement | Yes — 96% on parser/checker, 85% threshold in CI | L2 |
| Vulnerability scanning | Yes — govulncheck in CI, Dependabot for gomod + actions | L2 |
| Mutation testing | Yes — weekly via go-mutesting | L2 |
| Supply chain hardening | Yes — all actions pinned to commit hashes | L2 |
| CLAUDE.md | Yes — LP, CUPID, spec-first, TDD, quarterly cadence, Conventional Comments | L3 |
| HARNESS.md | Yes — 8 constraints (7 enforced, 1 unverified), 5 GC rules | L3 |
| AGENTS.md | Yes — 1 gotcha, 4 arch decisions, style, test strategy | L3 |
| MODEL_ROUTING.md | Yes — 6-agent routing table + token budget guidance | L3 |
| Custom skills | Yes — 5 (literate-programming, cupid-code-review, supply chain, dependency audit, assessment) | L3 |
| Custom agents | Yes — 7 (orchestrator, spec-writer, tdd-agent, go-implementer, code-reviewer, integration-agent, assessor) | L3 |
| Custom commands | Yes — 3 (/assess, /reflect, /worktree) | L3 |
| Hooks | Partial — 3 hooks from ai-literacy-superpowers plugin; no local settings.json | L3 |
| REFLECTION_LOG.md | Yes — 4 entries (3 prior + 1 from this session) | L3 |
| Health snapshots | Yes — 2 snapshots with trend comparison | L3 |
| Specifications | Yes — specs/001-link-checker/ with spec.md + plan.md | L4 |
| Implementation plans | Yes — plan.md with FR mapping, data model, algorithms | L4 |
| Orchestrator with gates | Yes — plan approval gate + MAX_REVIEW_CYCLES=3 | L4 |
| Pipeline exercised | Yes — 2 features (initial build + fragment validation) | L4 |
| Spec-first discipline | Yes — user reports always spec first, no exceptions | L4 |
| SBOM constraint declared | Yes — unverified, ready to activate with first dependency | L4 |
| Conventional Comments adopted | Yes — code-reviewer uses labelled feedback | L4 |
| Platform tooling (consumer) | Yes — ai-literacy-superpowers plugin | L5 partial |
| OTel configuration | No | L5 |
| Cross-team governance | No — exemplar is a teaching tool, not org governance | L5 |
| Cost data | No — MODEL_ROUTING.md exists, no spend reviewed | L5 |
| Multi-quarter cadence | No — cadence documented 2026-04-01, 13 days ago | L5 |

### Changes Since Last Assessment (2026-04-01)

| Change | Impact |
| --- | --- |
| SBOM constraint added to HARNESS.md | +1 constraint (unverified), enforcement ratio now 7/8 (88%) |
| Conventional Comments adopted | Code-reviewer uses structured feedback labels |
| REFLECTION_LOG.md grew from 2 to 4 entries | +2 reflections captured |
| Second health snapshot generated | Trend comparison now possible |
| actions/upload-artifact bumped to v7.0.1 | Dependabot keeping actions current |
| 41 commits since last assessment | Active development continuing |

### Clarifying Responses

1. **Cost data**: Not checked since last assessment. Cost discipline
   remains theoretical — MODEL_ROUTING.md guides routing but no actual
   spend has been observed.

2. **Session depletion**: Sometimes pushes past the point of diminishing
   returns on multi-step tasks. Session boundaries are not explicitly
   designed.

3. **Organisational scope**: The exemplar is a teaching tool used by one
   person. It does not govern other projects or teams. L5 requires
   organisational scope.

4. **Spec discipline**: Always spec first, no exceptions. This confirms
   the L4 workflow is genuine habit, not just documented aspiration.

---

## Level Assessment

### Primary Level: 4 — Specification Architect

The exemplar remains at Level 4. All L4 requirements are met with
increasing maturity:

- **Spec-first workflow is habitual**: user reports 100% adherence.
  The pipeline has been exercised on 2 features. The Conventional
  Comments addition shows the workflow is still evolving.

- **Agent pipeline with safety gates is operational**: orchestrator
  with plan approval gate and MAX_REVIEW_CYCLES=3. Seven agents
  coordinate the full lifecycle.

- **Compound learning pipeline is producing**: 4 reflections captured,
  1 promoted. Signal quality is 100% (all entries have Surprise +
  Proposal). However, 2 reflections from 2026-03-31 remain
  unpromoted — the human review step has slowed.

### Why Not Level 5

Level 5 requires three things this project does not yet have:

1. **Elapsed time**: the quarterly cadence was documented 13 days ago.
   A completed quarter (first due 2026-06-30) is required to prove
   the rhythm is sustainable.

2. **Cost data**: no spend has been reviewed. Cost discipline requires
   at least one observation.

3. **Organisational scope**: this is a single-person teaching project.
   L5 requires platform-level governance serving multiple teams. This
   is structural, not a deficiency — the exemplar was never intended
   to be L5 governance infrastructure.

### Discipline Maturity

| Discipline | Score | Evidence |
| --- | --- | --- |
| Context Engineering | 4/5 | CLAUDE.md with quarterly cadence + Conventional Comments. AGENTS.md with promoted gotcha. MODEL_ROUTING.md. 5 skills, 7 agents, 3 commands. Hooks via plugin. No iterative CLAUDE.md refinement across multiple quarters yet. |
| Architectural Constraints | 5/5 | 8 constraints (7 enforced, 1 intentionally unverified), 5/5 GC rules, 3 CI workflows, 96% coverage with threshold, govulncheck, Dependabot, mutation testing, actions pinned to hashes. Exemplary. |
| Guardrail Design | 4/5 | 7-agent pipeline with plan approval gate + review cycle guardrail. Pipeline exercised on 2 features. 4 reflections, 1 promoted. 2 health snapshots with trend comparison. Conventional Comments for structured review. Missing: cost data, OTel, multi-quarter cadence proof. |

### Compared to Previous Assessment

| Dimension | 2026-04-01 | 2026-04-14 | Change |
| --- | --- | --- | --- |
| Level | 4 | 4 | Stable |
| Context Engineering | 4/5 | 4/5 | Stable (Conventional Comments added) |
| Architectural Constraints | 5/5 | 5/5 | Stable (SBOM constraint added) |
| Guardrail Design | 4/5 | 4/5 | Stable (snapshots + trend comparison added) |
| Constraints | 7 | 8 | +1 (SBOM, unverified) |
| Reflections | 2 | 4 | +2 |
| Promotions | 1 | 1 | Stable (bottleneck) |
| Health snapshots | 1 | 2 | +1 |

---

## Strengths

1. **Spec-first discipline is genuine habit** — user reports 100%
   adherence without exceptions. This is the strongest L4 signal:
   the workflow is internalised, not just documented.

2. **Architectural constraints remain exemplary** — 5/5 with the
   addition of the SBOM constraint following the promotion ladder
   (declare unverified, activate when needed). Actions pinned to
   hashes. Dependabot keeping dependencies current.

3. **Health observability is maturing** — two snapshots with trend
   comparison. The harness health check surfaced the unpromoted
   reflections, demonstrating that the meta-observability layer is
   working as designed.

4. **Reflection signal quality is 100%** — every REFLECTION_LOG entry
   has both Surprise and Proposal fields with substantive content.
   The pipeline produces high-quality learning signal.

5. **Continuous evolution** — Conventional Comments, SBOM constraint,
   and health snapshots all added since the last assessment. The
   habitat is being actively refined, not just maintained.

## Gaps

1. **Unpromoted reflections** — 2 reflections from 2026-03-31 are 14
   days old without review. The compound learning pipeline produces
   signal that isn't being converted to institutional memory. The
   health check flagged this — the question is whether the flag gets
   acted on.

2. **No cost data** — MODEL_ROUTING.md guides routing but no actual
   spend has been observed. This was a gap in the previous assessment
   and remains unaddressed.

3. **Session depletion awareness is informal** — user reports sometimes
   pushing past diminishing returns. No explicit session boundaries or
   depletion checks are designed into the workflow.

4. **Quarterly cadence unproven** — 13 days into the first quarter.
   The cadence exists on paper but hasn't been tested by time.

5. **Hooks not locally declared** — still relying on plugin-provided
   hooks with no local settings.json. Audit surface remains opaque.

## Recommendations

1. **Review and promote unpromoted reflections now** — two entries are
   aging. Decide: promote to AGENTS.md, or explicitly dismiss with a
   reason. Either outcome closes the learning loop.

2. **Capture first cost snapshot** — check the Claude usage dashboard,
   record a data point in observability/costs/. Even one observation
   moves cost discipline from theoretical to observed.

3. **Add a Depletion Check convention to CLAUDE.md** — a simple rule:
   "After 90 minutes, pause and self-assess: am I still reviewing
   critically or rubber-stamping?" This addresses the "sometimes"
   depletion signal.

4. **Complete the first quarterly cadence by 2026-06-30** — run all
   five actions. The assessment you're running now counts as one.

5. **Declare hooks locally** — create settings.json with the three
   plugin hooks, making the advisory loop auditable from the repo.

---

## Immediate Adjustments Applied

1. **HARNESS.md Status section** — updated constraint count if stale
2. **README badge** — AI Literacy badge link updated to this assessment
3. **README mechanism map** — verified current, no changes needed
4. **README compound learning section** — updated reflection count
   from 3 to 4

## Workflow Operation Recommendations

See Phase 6 below — presented one at a time for accept/reject.

## Reflection

This third assessment reveals a project deepening at Level 4 rather
than stalling. The additions since 2026-04-01 (SBOM constraint,
Conventional Comments, health snapshots with trends, 2 more reflections)
show active habitat refinement. The level hasn't changed because L5
requires *time* — a structural requirement that no amount of
single-session work can satisfy.

The most actionable finding is the unpromoted reflections. The compound
learning pipeline is producing high-quality signal (100% have Surprise +
Proposal) but the human review bottleneck means that signal decays
instead of compounding. This is the configured-vs-operational pattern
recurring at a different layer.

Session depletion is a new signal worth watching. The user's honest
"sometimes" suggests the workflow would benefit from explicit session
boundaries — a small addition to CLAUDE.md that could prevent the
subtle quality erosion that comes from long sessions.

---

## Next Assessment

Suggested re-assessment date: 2026-07-01 (quarterly)

By then: first quarterly cadence should be complete, cost data should
exist, and the pipeline should have handled at least one more feature.
If all three are true, Level 5 becomes assessable (though organisational
scope may remain the structural ceiling for a single-person project).

Previous assessment: [2026-04-01](2026-04-01-assessment.md) (Level 4)
