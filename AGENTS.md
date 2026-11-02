# Agent guide: trackIT Systems listmonk-static

This repo holds the files listmonk serves for our newsletter: public pages (subscribe, unsubscribe, archive, login), system e-mails and campaign templates. See [README.md](README.md) for the layout.

## Rules from the brand repo

Look, voice and terms are defined in [trackIT-Systems/brand](https://github.com/trackIT-Systems/brand), not here. A copy of the pinned release is in [brand/](brand/) ([ADR-0001](docs/adr/0001-brand-release-copied-into-repo.md)):

- **Visuals** (colors, radii, fonts, buttons, logo): [brand/DESIGN.md](brand/DESIGN.md). On public pages use the variables from [public/static/brand/tokens.css](public/static/brand/tokens.css); raw color values stay out of [public/static/style.css](public/static/style.css). E-mails inline the palette hex values from [brand/tokens.json](brand/tokens.json), 6 digits, lowercase.
- **Writing any visible text**: [brand/VOICE.md](brand/VOICE.md) and [brand/glossary/terms.yml](brand/glossary/terms.yml).
- Never edit `brand/` or `public/static/brand/` by hand. To update, run `scripts/update-brand.sh VERSION`.
- Before committing, run `scripts/check-colors.sh`.

## Templates

- Files mirror listmonk's bundled `static/` layout. Keep file names and `{{ define }}` names, or listmonk stops picking them up.
- `email-templates/default-visual.json` is the source for the visual editor; `default-visual.tpl` is its rendered HTML. Change both together and check `/compare/default-visual` in the renderer; use only styles the editor can produce (buttons are `pill`, [ADR-0002](docs/adr/0002-pill-buttons.md)), or the next save in the editor overwrites them.
- Preview changes with `go run ./renderer` (see [renderer/README.md](renderer/README.md)).

## Decisions

When you make a decision that a future contributor would ask "why?" about, add an ADR in [docs/adr/](docs/adr/) (copy [0000-template.md](docs/adr/0000-template.md), next free number) and list it in the README table.
