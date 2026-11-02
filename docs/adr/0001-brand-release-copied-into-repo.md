# 0001. Brand release copied into the repo

- Status: Accepted
- Date: 2026-11-02
- Deciders: Jonas Höchst

## Context

The public pages and e-mails must follow [trackIT-Systems/brand](https://github.com/trackIT-Systems/brand). Consumers install a versioned brand release ([brand ADR-0010](https://github.com/trackIT-Systems/brand/blob/main/docs/adr/0010-releases-on-github.md)).

listmonk serves this repo as plain files (`--static-dir`). There is no build step and no `node_modules/`. Most mail clients ignore CSS variables, so e-mails need literal color values.

## Decision

- [scripts/update-brand.sh](../../scripts/update-brand.sh) copies one brand release into the repo: `DESIGN.md`, `VOICE.md`, `glossary/` and `tokens.json` into [brand/](../../brand/), `tokens.css` into [public/static/brand/](../../public/static/brand/), and the black logo to `public/static/logo.svg`. [brand/VERSION](../../brand/VERSION) records the version.
- Public pages load `tokens.css` and use only its variables.
- E-mail templates inline the palette hex values. [scripts/check-colors.sh](../../scripts/check-colors.sh) fails on any hex color that is not in the palette; CI runs it on every pull request.
- Files in `brand/` and `public/static/brand/` are never edited by hand.

## Consequences

- listmonk keeps working without a build step.
- A brand update is one pull request: run the script with the new version, then fix whatever the color check reports.
- The check catches colors outside the palette, not wrong roles (e.g. a palette color on the wrong background). Reviewers check those against `brand/DESIGN.md`.
