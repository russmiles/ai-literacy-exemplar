# mdcheck — Conventions

## Literate Programming

All code follows literate programming conventions:

1. Every file opens with a narrative preamble — why it exists, key
   design decisions, what it deliberately does NOT do
2. Documentation explains reasoning, not signatures
3. Presentation follows logical understanding
4. Each file has one clearly stated concern
5. Inline comments explain WHY, not WHAT

## Code Review

All reviews apply the CUPID lens:

- **Composable** — can it be used independently?
- **Unix philosophy** — does it do one thing well?
- **Predictable** — does it behave as its name suggests?
- **Idiomatic** — does it follow Go conventions?
- **Domain-based** — do names come from Markdown and link checking?

## Conventional Comments

All review feedback uses [Conventional Comments](https://conventionalcomments.org/) labels. Every comment is prefixed with a label that signals intent, and a decoration that signals whether it blocks merge.

**Labels:** `praise`, `issue`, `suggestion`, `nitpick`, `question`, `thought`, `todo`, `chore`, `note`

**Decorations:** `(blocking)` — must fix before merge. `(non-blocking)` — should not prevent merge. `(if-minor)` — fix only if trivial.

**Rules:**

1. Every review includes at least one `praise:` comment
2. Every `issue:` is paired with a `suggestion:` when possible
3. Use `question:` when uncertain whether something is a problem
4. Use `nitpick:` for trivial preferences — do not dress up preferences as issues

## Workflow

- **Spec-first**: changes flow through specs/001-link-checker/spec.md
  before touching implementation code
- **TDD**: red-green-refactor strictly
- **Branch discipline**: never commit directly to main; create an issue
  first, then a descriptive branch
- **Commit messages**: concise, what changed and why, no postamble

## PR Health Check

After every push:

1. Run `gh pr checks <number> --watch`
2. Fix any failures
3. Do not declare ready until all checks green

## Session Discipline

After 90 minutes of continuous AI-assisted work, pause and self-assess:
am I still reviewing critically, or am I rubber-stamping? If the answer
is uncertain, end the session. Resume fresh.

Signs of depletion: accepting output without reading diffs, skipping
test verification, feeling impatient with review cycles. These are
signals to stop, not push through.

## Quarterly Operating Cadence

Run these on a quarterly basis to keep the habitat healthy:

1. **`/assess`** — run a full AI literacy assessment
2. **`/harness-audit`** — verify declared enforcement matches reality
3. **Review REFLECTION_LOG.md** — promote worthy entries to AGENTS.md
   (also review monthly — entries older than 30 days without a
   promote/dismiss decision are stale)
4. **Check mutation testing trends** — compare weekly mutation scores over the quarter
5. **Run `/harness-health`** — generate a health snapshot and review trends
6. **Capture cost data** — record AI provider spend via `/cost-capture`
   or manually into `observability/costs/` (first capture due 2026-06-30)
