# 0002. Pill-shaped buttons

- Status: Accepted
- Date: 2026-11-02
- Deciders: Jonas Höchst

## Context

The brand gives buttons 16px corners (`radius.control`, [brand/DESIGN.md](../../brand/DESIGN.md) › Buttons). listmonk's visual editor cannot produce that: its buttons have 4px corners, 64px corners (pill) or none. A template saved in the editor overwrites any other value.

## Decision

All buttons in this repo are pills: public pages use `var(--radius-pill)`, e-mails `9999px`, and the visual editor its `pill` style (64px). Colors, label and size still follow the brand.

## Consequences

- Buttons look the same in every template, including ones saved in the visual editor.
- This differs from the brand rule. TODO(Nicole): decide whether the brand allows pill buttons, or whether this repo should go back to 16px wherever the editor is not involved.
