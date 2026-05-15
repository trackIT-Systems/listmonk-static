# Template preview server

Renders listmonk system templates from this repo with mock data and strings from [`i18n/`](../i18n/) (default `en.json`).

From the **repository root** (directory that contains `public/` and `email-templates/`):

```bash
go run ./renderer
```

Use another locale file by basename (must match a file in `i18n/`, e.g. `de.json`, `zh-CN.json`):

```bash
go run ./renderer -lang=de
go run ./renderer -lang=zh-CN
```

The root [`go.work`](../go.work) file registers the `renderer` module so this works without `cd renderer`.

Alternatively:

```bash
cd renderer && go run .
cd renderer && go run . -lang=de
```

Open <http://localhost:8088/> for links to each preview. Static assets are served from `public/static/` at `/public/static/`.

Listmonk-style URLs used in the templates (e.g. `{{ .RootURL }}/subscription/form`, `/admin/login`) are redirected to the matching preview page under `/public/...` so in-browser navigation works without editing the HTML.
