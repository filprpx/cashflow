---
status: proposed
---

# Theme-Aware Terminal Presentation

Cashflow will centralize visual styling behind application theme configurations using semantic roles such as background, surface, primary, muted, accent, warning, error, and focus. Screens and Bubbles components will receive theme-derived styles through the presentation and composition layer rather than hardcoding colors individually. Cashflow should remain compatible with terminal environments and system themes such as Omarchy without depending on Omarchy-specific files or APIs.

## Consequences

- Screens do not own global color decisions.
- Components remain reusable across themes.
- The default presentation should respect terminal color capabilities where possible.
- Transparent component surfaces are preferred unless a theme explicitly defines a filled surface.
- Theme selection and configuration can be added without changing domain modules.
