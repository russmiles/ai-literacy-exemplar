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
