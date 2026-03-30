# Model Routing

## Agent Routing

| Agent | Model tier | Rationale |
| --- | --- | --- |
| orchestrator | Most capable | Coordination, judgment |
| spec-writer | Most capable | Specification quality |
| tdd-agent | Standard | Structured test output |
| go-implementer | Standard | Plan-specified implementation |
| code-reviewer | Most capable | Nuance, quality judgment |
| integration-agent | Standard | Procedural workflow |

## Token Budget Guidance

| Role | Typical range | Escalation signal |
| --- | --- | --- |
| spec-writer | 50-100k | >100k: decompose |
| tdd-agent | 50-150k | >150k: split scenarios |
| go-implementer | 100-250k | >250k: decompose task |
| code-reviewer | 50-100k | >100k: batch review |
| integration-agent | 30-80k | >80k: investigate CI |
