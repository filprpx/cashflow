---
status: proposed
---

# Encrypted Vault Storage With a Separate Catalog

Each vault stores its configuration and financial data in one SQLCipher-backed encrypted SQLite database. A separate unencrypted catalog stores only discovery metadata, such as the vault name and database location, because the application must discover vaults before it can unlock one.

## Considered Options

- Application-level encryption was rejected because it would leave SQLite structure and metadata exposed and would require the application to reimplement more storage security.
- Storing the catalog inside each vault was rejected because locked vaults must still be discoverable.

## Consequences

- The catalog must never contain passphrases, encryption keys, credentials, or financial data.
- Opening a vault requires successful database authentication with its passphrase.
- SQLCipher becomes a deliberate storage technology dependency.
