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

## Default color palette

Hex colors are inlined in [`email-templates/`](email-templates/) and [`public/static/style.css`](public/static/style.css); there is no shared token file. Swatches below match those defaults (classic e-mail templates use the cool gray backdrop; [`default-visual`](email-templates/default-visual.json) uses the warmer neutrals).

| | Hex | Typical use |
|---|-----|-------------|
| <svg width="28" height="28" xmlns="http://www.w3.org/2000/svg"><title>Primary brand (trackit green)</title><rect width="28" height="28" rx="4" fill="#1e5d39" stroke="#ccc" stroke-width="1"/></svg> | `#1e5d39` | Links, buttons, focus rings |
| <svg width="28" height="28" xmlns="http://www.w3.org/2000/svg"><title>Link underline</title><rect width="28" height="28" rx="4" fill="#5a7a6b" stroke="#ccc" stroke-width="1"/></svg> | `#5a7a6b` | Link underline (`public/` only) |
| <svg width="28" height="28" xmlns="http://www.w3.org/2000/svg"><title>Error</title><rect width="28" height="28" rx="4" fill="#ff6b35" stroke="#ccc" stroke-width="1"/></svg> | `#ff6b35` | Error text (`public/` only) |
| <svg width="28" height="28" xmlns="http://www.w3.org/2000/svg"><title>Page backdrop</title><rect width="28" height="28" rx="4" fill="#f9f9f9" stroke="#ccc" stroke-width="1"/></svg> | `#f9f9f9` | Public page background |
| <svg width="28" height="28" xmlns="http://www.w3.org/2000/svg"><title>Email backdrop</title><rect width="28" height="28" rx="4" fill="#F0F1F3" stroke="#ccc" stroke-width="1"/></svg> | `#F0F1F3` | Classic e-mail outer background |
| <svg width="28" height="28" xmlns="http://www.w3.org/2000/svg"><title>Visual backdrop</title><rect width="28" height="28" rx="4" fill="#F5F5F5" stroke="#ccc" stroke-width="1"/></svg> | `#F5F5F5` | Visual builder e-mail backdrop / footer blocks |
| <svg width="28" height="28" xmlns="http://www.w3.org/2000/svg"><title>Shadow</title><rect width="28" height="28" rx="4" fill="#f3f3f3" stroke="#ccc" stroke-width="1"/></svg> | `#f3f3f3` | Input / card shadow tint |
| <svg width="28" height="28" xmlns="http://www.w3.org/2000/svg"><title>Code block</title><rect width="28" height="28" rx="4" fill="#f4f4f4" stroke="#ccc" stroke-width="1"/></svg> | `#f4f4f4` | Code/pre background in classic templates (some files mistakenly use `#f4f4f4f4`) |
| <svg width="28" height="28" xmlns="http://www.w3.org/2000/svg"><title>Surface</title><rect width="28" height="28" rx="4" fill="#ffffff" stroke="#ccc" stroke-width="1"/></svg> | `#fff` | Cards, content panels |
| <svg width="28" height="28" xmlns="http://www.w3.org/2000/svg"><title>Divider</title><rect width="28" height="28" rx="4" fill="#CCCCCC" stroke="#ccc" stroke-width="1"/></svg> | `#CCCCCC` | Strong divider (visual e-mails) |
| <svg width="28" height="28" xmlns="http://www.w3.org/2000/svg"><title>Light border</title><rect width="28" height="28" rx="4" fill="#eee" stroke="#ccc" stroke-width="1"/></svg> | `#eee` | Light borders, separators |
| <svg width="28" height="28" xmlns="http://www.w3.org/2000/svg"><title>Input border</title><rect width="28" height="28" rx="4" fill="#ddd" stroke="#ccc" stroke-width="1"/></svg> | `#ddd` | Inputs, borders |
| <svg width="28" height="28" xmlns="http://www.w3.org/2000/svg"><title>Body text</title><rect width="28" height="28" rx="4" fill="#444" stroke="#ccc" stroke-width="1"/></svg> | `#444` | Secondary labels, classic e-mail body copy |
| <svg width="28" height="28" xmlns="http://www.w3.org/2000/svg"><title>Muted</title><rect width="28" height="28" rx="4" fill="#666" stroke="#ccc" stroke-width="1"/></svg> | `#666` | Hints, archive dates |
| <svg width="28" height="28" xmlns="http://www.w3.org/2000/svg"><title>Footer</title><rect width="28" height="28" rx="4" fill="#888" stroke="#ccc" stroke-width="1"/></svg> | `#888` | Footer links, muted text |
| <svg width="28" height="28" xmlns="http://www.w3.org/2000/svg"><title>Footer subtle</title><rect width="28" height="28" rx="4" fill="#aaa" stroke="#ccc" stroke-width="1"/></svg> | `#aaa` | Public footer text |
| <svg width="28" height="28" xmlns="http://www.w3.org/2000/svg"><title>Primary text</title><rect width="28" height="28" rx="4" fill="#111" stroke="#ccc" stroke-width="1"/></svg> | `#111` | Primary body text, hovers |
| <svg width="28" height="28" xmlns="http://www.w3.org/2000/svg"><title>Visual text</title><rect width="28" height="28" rx="4" fill="#262626" stroke="#ccc" stroke-width="1"/></svg> | `#262626` | Visual e-mail main text |
| <svg width="28" height="28" xmlns="http://www.w3.org/2000/svg"><title>Dark bar</title><rect width="28" height="28" rx="4" fill="#222" stroke="#ccc" stroke-width="1"/></svg> | `#222` | Dark header bar (some e-mails) |
| <svg width="28" height="28" xmlns="http://www.w3.org/2000/svg"><title>Dark UI</title><rect width="28" height="28" rx="4" fill="#333" stroke="#ccc" stroke-width="1"/></svg> | `#333` | Button hover, dark accents |

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
