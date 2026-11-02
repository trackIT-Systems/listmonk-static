# listmonk-static

Custom [listmonk](https://listmonk.app) **system static files**: public pages, system e-mail templates, and locale strings. These replace the files bundled with listmonk when you point the binary at this directory (see the listmonk docs on [system templates](https://listmonk.app/docs/templating/#system-templates)).

## Layout

| Path | Role |
|------|------|
| [`public/templates/`](public/templates/) | Public HTML (subscription, opt-in, archive, admin login shell, …) |
| [`public/static/`](public/static/) | CSS, JS, logos, favicon, … |
| [`email-templates/`](email-templates/) | System e-mails (`base.html`, opt-in, import status, …) |
| [`i18n/`](i18n/) | Locale JSON files (e.g. `en.json`) |
| [`renderer/`](renderer/) | Small Go app to preview templates in a browser with mock data |
| [`brand/`](brand/) | Copy of the pinned brand release (rules, glossary, tokens); do not edit |
| [`scripts/`](scripts/) | Brand update and color check |
| [`docs/adr/`](docs/adr/) | Decision records |

## Brand

Colors, radii, fonts, buttons and the logo follow [trackIT-Systems/brand](https://github.com/trackIT-Systems/brand). A copy of the pinned release is in [`brand/`](brand/) (version in [`brand/VERSION`](brand/VERSION)); see [ADR-0001](docs/adr/0001-brand-release-copied-into-repo.md).

- Public pages load [`public/static/brand/tokens.css`](public/static/brand/tokens.css) and use its variables in [`style.css`](public/static/style.css).
- E-mails inline the palette hex values, because most mail clients ignore CSS variables.
- [`scripts/check-colors.sh`](scripts/check-colors.sh) fails on any hex color outside the palette. CI runs it on every pull request.

To update the brand, run the script with the release version, check, and open a pull request:

```bash
scripts/update-brand.sh 0.2.0
scripts/check-colors.sh
```

## Decisions

| ADR | Decision | Status |
|---|---|---|
| [0001](docs/adr/0001-brand-release-copied-into-repo.md) | Brand release copied into the repo | Accepted |
| [0002](docs/adr/0002-pill-buttons.md) | Pill-shaped buttons | Accepted |

## Using with listmonk

Copy or mount this repo (or a subset) and start listmonk with your static path, for example:

```bash
listmonk --static-dir=/path/to/listmonk-static
```

Paths under `public/` and `email-templates/` match listmonk’s bundled layout so you can override only what you need.

## Local template preview

To render the same templates without a full listmonk install, use the preview server in [`renderer/`](renderer/README.md):

```bash
go run ./renderer
go run ./renderer -lang=de
```

The repo includes [`go.work`](go.work) so `go run ./renderer` works from the repository root. Then open the URL printed in the log (by default `http://localhost:8088/`) for an index of public and e-mail previews. Use `-lang` with the basename of any file in [`i18n/`](i18n/) (e.g. `de`, `zh-CN`; default `en`).

## Language in listmonk

In the admin UI, set **Settings → General → Language** to the locale you want. listmonk loads the matching JSON from the static directory’s `i18n/` folder (this repo’s [`i18n/`](i18n/) when using `--static-dir` here). Add or edit `*.json` files there to customize translations.
