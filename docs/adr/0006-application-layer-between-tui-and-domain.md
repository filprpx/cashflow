---
status: proposed
---

# Application Layer Between TUI and Domain Modules

The Bubble Tea interface is a delivery mechanism, not the owner of vault, storage, security, or workflow behavior. TUI screens communicate with application operations through narrow interfaces, while the application layer coordinates domain modules and adapters; `cmd/tui` is the initial entrypoint and a future `cmd/cli` may reuse the same application layer.

## Consequences

- TUI behavior can be tested without a real filesystem, database, or terminal.
- A future CLI can reuse vault and workflow operations without importing TUI state.
- Domain modules remain independent of Bubble Tea and presentation concerns.
