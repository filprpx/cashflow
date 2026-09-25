---
status: proposed
---

# Vaults Are Isolated Workspaces

A vault is the unit of isolation for application configuration and personal financial data. Each vault has its own identity, location, passphrase, and database, allowing users to maintain independent environments without mixing their data.

## Consequences

- The application can manage multiple vaults.
- The active vault must be explicitly selected or created.
- Vault operations must not depend on another vault being open.
