# Harness Health Snapshot — 2026-04-14 (revised)

## Enforcement

- Constraints: 9/13 enforced (69%)
- Technical: 5 deterministic | 2 agent | 1 weekly deterministic | 1 unverified
- Governance: 1 deterministic | 1 agent | 3 unverified
- Unverified: SBOM generation, secrets prevention (pre-commit hook),
  secrets detection (CI gitleaks), approved dependency licenses
- Drift detected: no

## Garbage Collection

- Rules active: 5/5
- Findings since last snapshot: 0

## Mutation Testing

- Go kill rate: available via CI artifacts (mutation-testing.yml, weekly)
- Trend: no new data since last snapshot — check CI artifacts

## Compound Learning

- REFLECTION_LOG entries: 8 (3 new since last snapshot)
- Entries with signal (Surprise + Proposal): 8/8 (100%)
- AGENTS.md gotchas: 3 (2 new since last snapshot)
- AGENTS.md arch decisions: 4
- Promotions since last snapshot: 2 (L4-to-L5, two-speed-pipeline)
- Unpromoted reflections awaiting review: 3 (entries 6, 7, 8; oldest
  from 2026-04-14)

## Session Quality

- Depletion check documented: yes (CLAUDE.md)
- 90-minute pause cadence: documented

## Operational Cadence

- Last /harness-audit: 2026-04-14 (today, quarterly — on schedule)
- Last /assess: 2026-04-14 (today, quarterly — on schedule)
- Last /reflect: 2026-04-14 (today)
- Last snapshot: 2026-04-14 (today, revised after governance constraints)
- Quarterly cadence documented: yes (CLAUDE.md)
- Outer loop overdue: no

## Cost Indicators

- Model routing configured: yes (MODEL_ROUTING.md)
- Cost snapshots captured: 0 (first capture due 2026-06-30)

## Meta

- Snapshot currency: current (today)
- Cadence compliance: all cadences on schedule
- Learning flow: active — 3 new reflections since last snapshot, 2 new
  promotions to AGENTS.md. The promotion bottleneck identified in the
  previous snapshot has been addressed.
- GC effectiveness: 5 rules configured, 0 findings (2 consecutive
  snapshots with 0 findings — not yet at the 3-snapshot silent threshold)
- Trend alerts: none (insufficient snapshots for 3-point decline
  detection)
- Enforcement ratio: 69% — below the 70% healthy threshold due to 4
  new unverified governance constraints. This is expected: governance
  constraints start unverified and promote over time.
- **Aggregate health: Attention** (enforcement ratio < 70%)

## Trends (vs 2026-04-14 previous)

| Metric | Previous | Current | Delta |
| --- | --- | --- | --- |
| Constraints total | 8 | 13 | +5 (governance) |
| Constraints enforced | 7/8 (88%) | 9/13 (69%) | -19% (4 new unverified) |
| GC rules active | 5/5 | 5/5 | stable |
| REFLECTION_LOG entries | 5 | 8 | +3 |
| Entries with signal | 100% | 100% | stable |
| AGENTS.md gotchas | 1 | 3 | +2 |
| AGENTS.md promotions | 0 since prev | 2 since prev | +2 (flow unblocked) |
| Unpromoted entries | 4 | 3 | -1 (improving) |
| Snapshot age (days) | 0 | 0 | current |
| Cadence compliance | on schedule | on schedule | stable |
| Cost snapshots | 0 | 0 | stable (first due 2026-06-30) |
