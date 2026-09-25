# Cashflow

Cashflow is a personal financial data control application. It connects to external financial systems, preserves data locally, transforms it into a stable financial model, and sends it to other local or external systems.

## Workspace

**Vault**:
A named, isolated workspace containing the configuration and financial data for one user-controlled environment. A vault is protected by a passphrase.
_Avoid_: Profile, project, account

**Vault catalog**:
The application-level inventory of known vaults and the metadata needed to find them. It does not contain financial data, passphrases, or secrets.
_Avoid_: Vault database, registry

**Vault database**:
The encrypted local database owned by one vault. It contains the vault's configuration and financial data.
_Avoid_: Catalog, application database

**Vault session**:
The period during which a vault has been successfully unlocked and its data is available to the application.
_Avoid_: Login, connection

**Passphrase**:
The user-provided secret used to unlock a vault and authorize destructive vault operations.
_Avoid_: Password, key

## Financial Data

**Source integration**:
An integration that retrieves data from an external financial system.
_Avoid_: Importer, connector, plugin

**Target integration**:
An integration that sends application-owned financial data to an external system.
_Avoid_: Exporter, destination, plugin

**Mapping**:
The transformation that converts source-specific data into the application's canonical financial model or converts canonical data into a target-specific representation.
_Avoid_: Conversion, translation

**Canonical financial model**:
The application-owned representation of financial data that remains stable across source and target integrations.
_Avoid_: Provider model, API model

## Automation

**Integration**:
A capability for exchanging data with an external system. An integration may provide source, target, or both capabilities.
_Avoid_: Plugin, addon, extension

**Workflow**:
A configured sequence of source retrieval, storage, mapping, and target delivery operations.
_Avoid_: Job, pipeline, sync

**Run**:
One execution of a workflow, including its result and any reported failures.
_Avoid_: Task, invocation
