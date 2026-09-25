---
status: proposed
---

# Integrations Exchange Through a Canonical Financial Model

Source integrations do not map directly to target integrations. Retrieved data is preserved and transformed into an application-owned canonical financial model before it is delivered to targets, allowing MeuPluggy and Wealthfolio to evolve independently and enabling additional integrations later.

## Considered Options

- Direct MeuPluggy-to-Wealthfolio mapping was rejected because it couples two external schemas and prevents replaying or reusing imported data for other targets.

## Consequences

- The application owns the meaning and compatibility of its canonical model.
- Source-specific and target-specific representations remain at integration seams.
- Mapping and delivery can be retried without retrieving source data again, subject to the workflow's freshness policy.
