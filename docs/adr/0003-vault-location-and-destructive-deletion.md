---
status: proposed
---

# User-Selectable Vault Locations With Authenticated Deletion

New vaults use a platform-appropriate user data directory by default, while allowing the user to choose another database location. Deleting a vault is permanent: it requires entering the vault passphrase and removes both the vault database and its catalog entry.

## Consequences

- The default path provides a safe zero-configuration experience.
- Users can place vault data on another disk or storage location when needed.
- A missing or inaccessible database must be reported rather than silently treated as an empty vault.
- The application must make the irreversible nature of deletion clear before requesting confirmation.
