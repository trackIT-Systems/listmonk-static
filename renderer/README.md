# Template preview server

Renders listmonk system templates from this repo with mock data and [i18n/en.json](../i18n/en.json) strings.

From the **repository root** (directory that contains `public/` and `email-templates/`):

```bash
go run ./renderer
```

The root [`go.work`](../go.work) file registers the `renderer` module so this works without `cd renderer`.

Alternatively:

```bash
cd renderer && go run .
```

Open <http://localhost:8088/> for links to each preview. Static assets are served from `public/static/` at `/public/static/`.

Listmonk-style URLs used in the templates (e.g. `{{ .RootURL }}/subscription/form`, `/admin/login`) are redirected to the matching preview page under `/public/...` so in-browser navigation works without editing the HTML.
