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
```

The repo includes [`go.work`](go.work) so `go run ./renderer` works from the repository root. Then open the URL printed in the log (by default `http://localhost:8088/`) for an index of public and e-mail previews.
