# Reflection Log

<!-- Agents append reflections here. Human reviews and promotes to
     AGENTS.md. Append-only. -->

## Entry Format

- **Date**: YYYY-MM-DD
- **Agent**: which agent
- **Task**: what was done
- **Surprise**: what was unexpected
- **Proposal**: what to add to AGENTS.md
- **Improvement**: what would improve the process

---

- **Date**: 2026-03-31
- **Agent**: assessor (via /assess command)
- **Task**: First AI literacy assessment of the exemplar — Level 3 Habitat Engineer
- **Surprise**: The exemplar scored 10/10 on L3 habitat signals but the weakest discipline (guardrail design at 3/5) pulled the ceiling down from L4. The gap isn't missing infrastructure — it's missing practice. REFLECTION_LOG.md was empty, the harness hadn't been revisited, and the orchestrator pipeline had only been used once. The distinction between "configured" and "operational" is the difference between L3 and L4.
- **Proposal**: Add to GOTCHAS: "A fully configured habitat (all signals present) does not mean an operational one. The assessment distinguishes between infrastructure that exists and infrastructure that's been exercised. Empty REFLECTION_LOG.md and unrevised HARNESS.md are the clearest signals that the habitat is configured but not yet lived in."
- **Improvement**: The assessment revealed that the exemplar has no operating cadence documented anywhere. Add a "Quarterly Operating Cadence" section to CLAUDE.md that specifies when to run /harness-audit, review reflections, and check cost data.

---

- **Date**: 2026-04-01
- **Agent**: assessor (via /assess command)
- **Task**: Second AI literacy assessment — Level 4 Specification Architect
- **Surprise**: The delta between L3 and L4 was entirely practice, not infrastructure. All three gaps from the L3 assessment (empty reflections, no cadence, pipeline not re-exercised) were closed within 48 hours because the infrastructure was already in place. The configured-vs-operational gotcha promoted to AGENTS.md turned out to be the single most important insight: it predicted exactly what needed to happen. The remaining gap to L5 is qualitatively different — it requires elapsed time (multi-quarter cadence) and organisational scope (platform tooling, OTel), not sprint-level work.
- **Proposal**: Add to AGENTS.md: "The gap between Level 4 and Level 5 is not infrastructure — it is elapsed time and organisational scope. L5 cannot be sprinted to. It requires sustained quarterly practice, actual cost data, and tooling that serves teams beyond the originating repo."
- **Improvement**: Future assessments should check whether the quarterly cadence has been sustained across a full quarter, not just documented. A documented cadence with zero completed cycles is still L3.

---

- **Date**: 2026-03-31
- **Agent**: go-implementer (via spec-first TDD pipeline)
- **Task**: Add fragment link validation to mdcheck (FR-011)
- **Surprise**: Two of the three fragment tests passed immediately because the existing code already strips fragments and checks file existence — only the "missing heading" case failed, which is the actual new behaviour. This is a healthy sign: the existing design was close to correct and the new feature slotted in cleanly. The slug normalisation was straightforward because the spec constrained it to GitHub-compatible slugs, avoiding the rabbit hole of supporting every possible slug algorithm.
- **Proposal**: Add to AGENTS.md: "When existing tests pass for new feature scenarios, it means the existing code partially implements the feature. The failing test is the one that matters — it defines the boundary between what exists and what's new."
- **Improvement**: The spec-first workflow surfaced the slug normalisation concern during scenario writing, before any code was written. Writing FR-011 forced a decision about which slug algorithm to support, which prevented scope creep during implementation.

---

- **Date**: 2026-04-14
- **Agent**: harness-health (via /harness-health command)
- **Task**: Harness health snapshot generation — second snapshot, first trend comparison
- **Surprise**: Despite 100% signal quality on reflections (all 3 entries have Surprise + Proposal), the learning flow has stalled: two reflections from 2026-03-31 are now 14 days old with no promotion review. The go-implementer's "failing test defines the boundary" insight and the assessor's "L4→L5 requires elapsed time" insight are both sitting unpromoted. The compound learning pipeline is producing high-quality inputs but the human review step — the bottleneck — hasn't fired. This is the configured-vs-operational gotcha applied to the learning system itself.
- **Proposal**: Add to GOTCHAS: "The compound learning pipeline has two speeds: agent-speed (reflections appear immediately) and human-speed (promotions require review). If reflections accumulate without promotion reviews, the pipeline produces signal that never becomes institutional memory. The /harness-health nudge for unpromoted reflections exists precisely to close this gap — act on it."
- **Improvement**: The health snapshot should track "days since last promotion review" as a first-class metric, not just count promotions. A promotion count of zero is ambiguous — it could mean nothing worth promoting, or it could mean nobody looked. Tracking review dates would distinguish the two.

---

- **Date**: 2026-04-14
- **Agent**: assessor (via /assess command)
- **Task**: Third AI literacy assessment — Level 4 Specification Architect (deepening)
- **Surprise**: The project is deepening at L4 rather than stalling or regressing. SBOM constraint, Conventional Comments, and health snapshot trends all added since last assessment — the habitat is being actively refined. But the compound learning bottleneck (unpromoted reflections) persisted for 14 days despite being flagged by the health check. The meta-observability worked (it surfaced the issue) but the human action loop didn't close. This is the configured-vs-operational pattern recurring at the observability layer: the system that detects staleness is itself subject to staleness if the human doesn't act on its output.
- **Proposal**: Add to GOTCHAS: "Meta-observability only works if humans act on what it surfaces. A health check that flags unpromoted reflections is useless if nobody reads the flag. The final link in every observability chain is a human decision — design the chain to make that decision obvious and low-friction."
- **Improvement**: The Depletion Check and monthly reflection review cadence added in this assessment are the first workflow conventions that address human factors rather than tooling. Future assessments should check whether these conventions change behaviour or just add documentation.

---

- **Date**: 2026-04-14
- **Agent**: harness-health + human review session
- **Task**: Health snapshot refresh, promotion review of 4 unpromoted reflections, and adding Observability section to HARNESS.md
- **Surprise**: The promotion review took three PRs in sequence (snapshot update, promotions, observability section) but the actual decision-making — which entries to promote, which to skip, where each belongs in AGENTS.md — was fast because every reflection had a concrete Proposal field. The structured format (Surprise + Proposal) eliminated the "what did this entry mean again?" overhead that makes promotion reviews drag. Entry 5 was correctly identified as redundant with Entry 4, which validates that not every reflection needs promotion — curation means saying no too.
- **Proposal**: Add to GOTCHAS: "Structured reflection entries (with explicit Surprise and Proposal fields) make promotion reviews fast because the decision is pre-framed. Unstructured entries require the reviewer to re-derive intent, which adds friction and delays the review. Signal quality at capture time determines throughput at promotion time."
- **Improvement**: The three-PR workflow (snapshot, promotions, observability) could have been a single PR if the work had been planned as a batch. Future health-check sessions that include promotion reviews should branch once, make all changes, and open one PR — reducing CI round-trips from 3 to 1.
