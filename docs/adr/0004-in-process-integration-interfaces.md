---
status: proposed
---

# Integrations Use In-Process Interfaces Initially

External systems such as MeuPluggy and Wealthfolio are represented by in-process Go integrations behind application-owned interfaces. The first version will not load dynamic plugins or external extension processes; extensibility comes from registering additional implementations of those interfaces.

## Considered Options

- Dynamic plugins were deferred because they add deployment, versioning, security, and lifecycle complexity before the extension points are understood.
- A single integration interface for every capability was rejected in favor of source and target capabilities that can vary independently.

## Consequences

- New integrations can be added without changing workflow logic.
- Integrations remain part of the application's deployment and release process.
- A dynamic extension mechanism can be added later if a concrete use case justifies it.
