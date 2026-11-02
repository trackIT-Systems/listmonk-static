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

Open <http://localhost:8088/> for links to each preview. Use `-addr :8099` to listen on another port. Static assets are served from `public/static/` at `/public/static/`.

## Campaign templates

The `*.tpl` files in [`email-templates/`](../email-templates/) are shown under `/campaign/`. Where a template includes the campaign body (`{{ template "content" . }}`), the renderer fills in mock content with headings, a link, a button, a list, a table and code, so every styled element is visible.

`default-visual.json` is the visual editor source of `default-visual.tpl`. The renderer turns its blocks into HTML the way the editor does (approximately; only the block types the file uses), and `/compare/default-visual` shows both side by side. If they differ, the editor will overwrite the `.tpl` the next time someone saves the template there.

Listmonk-style URLs used in the templates (e.g. `{{ .RootURL }}/subscription/form`, `/admin/login`) are redirected to the matching preview page under `/public/...` so in-browser navigation works without editing the HTML.
