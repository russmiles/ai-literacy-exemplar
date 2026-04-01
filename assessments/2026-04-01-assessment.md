# AI Literacy Assessment — ai-literacy-exemplar

**Date**: 2026-04-01
**Assessed by**: assessor agent (via /assess command)
**Assessed level**: Level 4 — Specification Architect

---

## Observable Evidence

### Repository Signals

| Signal | Found | Level indicator |
| --- | --- | --- |
| CI workflows | Yes — go-tests.yml, lint-markdown.yml, mutation-testing.yml (3) | L2 |
| Test coverage enforcement | Yes — 96% on parser/checker, 85% threshold in CI | L2 |
| Vulnerability scanning | Yes — govulncheck in go-tests.yml, Dependabot for gomod + actions | L2 |
| Mutation testing | Yes — weekly via go-mutesting (mutation-testing.yml) | L2 |
| CLAUDE.md | Yes — 50 lines, literate programming, CUPID, spec-first, TDD, quarterly cadence | L3 |
| HARNESS.md | Yes — 7 constraints, 7/7 enforced (4 deterministic, 2 agent, 1 weekly) | L3 |
| AGENTS.md | Yes — 3 arch decisions, 1 gotcha (promoted from reflections), 3 style entries, test strategy | L3 |
| MODEL_ROUTING.md | Yes — agent routing table + token budget guidance for 6 agents | L3 |
| Custom skills | Yes — 5 (literate-programming, cupid-code-review, github-actions-supply-chain, dependency-vulnerability-audit, ai-literacy-assessment) | L3 |
| Custom agents | Yes — 7 (orchestrator, spec-writer, tdd-agent, go-implementer, code-reviewer, integration-agent, assessor) | L3 |
| Custom commands | Yes — 3 (/assess, /reflect, /worktree) | L3 |
| Hooks configured | Partial — 3 hooks from ai-literacy-superpowers plugin; no local settings.json | L3 |
| REFLECTION_LOG.md | Yes — 2 entries (assessment reflection + fragment validation reflection) | L3 |
| Specifications directory | Yes — specs/001-link-checker/ with spec.md and plan.md | L4 |
| Implementation plans | Yes — plan.md with module structure, data model, FR mapping | L4 |
| Orchestrator with safety gates | Yes — plan approval gate + MAX_REVIEW_CYCLES=3 | L4 |
| Pipeline exercised on multiple features | Yes — initial build + fragment validation (FR-011) | L4 |
| Plugin/platform tooling | Partial — uses ai-literacy-superpowers plugin, does not originate platform tooling | L5 |
| OTel configuration | No | L5 |
| Cross-team governance | No | L5 |

### Evidence Summary

The exemplar has progressed significantly since the Level 3 assessment on
2026-03-31. The three gaps identified in that assessment have all been
addressed:

1. **Compound learning activated** — REFLECTION_LOG.md now has 2 entries
   and the configured-vs-operational insight has been promoted to
   AGENTS.md GOTCHAS. The learning pipeline (reflect, review, promote)
   has been demonstrated end to end.

2. **Operating cadence documented** — CLAUDE.md now includes a quarterly
   operating cadence section with five recurring actions. The first
   assessment, harness audit, reflection capture, and health snapshot
   have all been executed.

3. **Pipeline exercised on a second feature** — fragment link validation
   (FR-011) was delivered end-to-end through the spec-first pipeline:
   spec updated, plan updated, failing tests written, implementation
   made green, reflection captured. This proves the workflow operates
   for incremental changes, not just initial builds.

The repo also now has 5/5 GC rules (up from 1), an accurate health
snapshot, and correct README counts.

What remains thin: the operating cadence has been documented but only
one cycle has been started (not completed). Cost data has not been
collected. There is no OTel or automated trend tracking. The repo
consumes platform tooling (the superpowers plugin) but does not
originate cross-team standards.

---

## Level Assessment

### Primary Level: 4 — Specification Architect

The exemplar meets Level 4 requirements across all three disciplines:

- **Spec-first workflow is operational**: two features have been
  delivered through it (initial build + fragment validation). The spec
  drives tests, tests drive implementation. This is not aspirational;
  it has been exercised.

- **Agent pipeline with safety gates**: the orchestrator coordinates
  seven agents with a plan approval gate and a MAX_REVIEW_CYCLES=3
  guardrail. The pipeline has been used for a real incremental feature.

- **Compound learning is active**: reflections have been captured AND
  promoted. AGENTS.md contains a gotcha that was proposed in
  REFLECTION_LOG.md and curated by a human. The learning loop is
  operational, not just configured.

The exemplar does NOT meet Level 5. Level 5 requires platform-level
governance, cross-team standards, observability with trends (OTel or
equivalent), and a proven multi-quarter operating cadence. The cadence
has been documented and started but not yet completed across a full
quarter. Cost data has not been collected. There is no automated
observability pipeline.

### Discipline Maturity

| Discipline | Strength (1-5) | Evidence |
| --- | --- | --- |
| Context Engineering | 4/5 | CLAUDE.md with quarterly cadence, AGENTS.md with promoted gotcha, MODEL_ROUTING.md, 5 skills, 7 agents, 3 commands. Hooks via plugin. No evidence of multi-session CLAUDE.md refinement beyond the cadence addition. |
| Architectural Constraints | 5/5 | 7/7 constraints enforced (4 deterministic, 2 agent, 1 weekly), 5/5 GC rules, 3 CI workflows, 96% coverage with threshold, govulncheck, Dependabot, mutation testing. Exemplary. |
| Guardrail Design | 4/5 | 7-agent pipeline with plan approval gate and MAX_REVIEW_CYCLES=3. Pipeline exercised on 2 features. 2 reflections captured, 1 promoted. Operating cadence documented. Health snapshot exists. Missing: OTel, cost data, multi-quarter cadence proof. |

### The Weakest Discipline

Context engineering and guardrail design are both at 4/5. Architectural
constraints at 5/5 are the strongest. The ceiling is set by the two
disciplines at 4, which aligns with Level 4.

To reach Level 5, guardrail design would need: automated observability
trends (not just snapshots), proven quarterly cadence (at least two
cycles completed), and cost discipline with actual data.

Context engineering would need: evidence of iterative CLAUDE.md
refinement across multiple quarters, and platform-level context
templates that could serve other teams.

---

## Strengths

- **Architectural constraints remain exemplary** — 7/7 enforced, zero
  unverified, 5 GC rules active. The harness is the strongest aspect
  of this project and has been since the initial build.

- **Spec-first workflow proven on incremental change** — FR-011
  (fragment validation) demonstrated the full cycle: spec update, plan
  update, failing tests, implementation, reflection. This moves the
  spec-first workflow from "declared" to "operational".

- **Compound learning pipeline demonstrated end-to-end** — the
  configured-vs-operational insight went from REFLECTION_LOG.md to
  AGENTS.md GOTCHAS. Future agents will benefit from this without
  needing to rediscover it.

- **Rapid remediation** — the three gaps identified in the L3
  assessment were all addressed within 24 hours: cadence documented,
  reflections promoted, pipeline exercised on a second feature.

- **Honest self-assessment** — the L3 assessment correctly identified
  the distinction between configured and operational, which is itself
  a Level 4 insight.

## Gaps

- **No cost data collected** — MODEL_ROUTING.md exists with tier
  guidance, but no actual cost data has been reviewed. Cost discipline
  is theoretical.

- **Operating cadence not yet proven across a full quarter** — the
  cadence was documented on 2026-04-01 and the first actions have been
  taken, but a quarterly cycle requires at least one completed quarter
  to verify the rhythm is sustained.

- **No automated observability** — the health snapshot is a point-in-time
  document, not an automated trend. There is no OTel, no dashboard, no
  automated staleness alerting. The investigative loop depends on human
  memory to trigger.

- **Only two features through the pipeline** — two is enough to prove
  the workflow is operational (not a one-off), but a larger sample
  would build confidence that the pipeline handles diverse change types.

- **Hooks not locally declared** — the three hooks come from the
  ai-literacy-superpowers plugin and are not visible in the repo's own
  configuration. This makes the enforcement surface harder to audit
  from the repo alone.

## Recommendations

1. **Collect cost data in the next quarterly review** — check the AI
   provider's usage dashboard. Record the spend in the health snapshot.
   Compare with MODEL_ROUTING.md guidance. This is the simplest gap to
   close and moves cost discipline from theoretical to observed.

2. **Complete the first quarterly cadence** — by 2026-06-30, run all
   five quarterly actions: /assess, /harness-audit, reflection review,
   mutation trend check, /harness-health. Document results. A completed
   quarter proves the cadence is sustainable.

3. **Add automated snapshot staleness check to CI** — a simple workflow
   that fails if the most recent snapshot is older than 30 days would
   close the meta-observability gap without requiring full OTel.

4. **Exercise the pipeline on a third feature** — choose a feature that
   requires a different kind of change (e.g. a new output format, a
   configuration option) to prove the pipeline handles variety.

5. **Add local hooks configuration** — create a settings.json that
   declares the three plugin hooks explicitly, making the advisory loop
   auditable from the repo without needing to inspect plugin internals.

---

## Immediate Adjustments Applied

1. Updated AI Literacy badge in README.md from Level 3 to Level 4
2. Updated badge link to point to this assessment document

## Workflow Operation Recommendations

| Recommendation | Frequency | First due |
| --- | --- | --- |
| Collect AI usage cost data and record in health snapshot | Quarterly | 2026-06-30 |
| Complete all five quarterly cadence actions | Quarterly | 2026-06-30 |
| Add CI check for snapshot staleness (<30 days) | Once | Next PR cycle |
| Exercise pipeline on a third feature | Once | Next feature |
| Declare hooks locally in settings.json | Once | Next PR cycle |

## Reflection

This assessment reveals a healthy progression from Level 3 to Level 4 in
under 48 hours. The speed of remediation is itself a signal: the
infrastructure was already in place, it just needed to be activated.
The configured-vs-operational distinction identified in the L3 assessment
was exactly right — the delta between L3 and L4 was practice, not
infrastructure.

The remaining gap to Level 5 is qualitatively different. Level 5
requires sustained operation over time (multiple quarters), automated
observability (not just snapshots), and platform-level governance that
serves teams beyond this repo. These are not things that can be
demonstrated in a sprint — they require elapsed time and organisational
scope.

Future assessments should watch for: whether the quarterly cadence is
actually sustained, whether cost data appears in snapshots, and whether
the pipeline handles diverse change types without friction.

---

## Next Assessment

Suggested re-assessment date: 2026-07-01 (quarterly)

Previous assessment: [2026-03-31](2026-03-31-assessment.md) (Level 3 — Habitat Engineer)
