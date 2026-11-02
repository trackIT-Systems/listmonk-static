package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

// mustRoot returns the listmonk-static repo root (directory containing public/ and email-templates/).
func mustRoot() string {
	if wd, err := os.Getwd(); err == nil {
		if st, err := os.Stat(filepath.Join(wd, "public", "templates")); err == nil && st.IsDir() {
			return wd
		}
		parent := filepath.Clean(filepath.Join(wd, ".."))
		if st, err := os.Stat(filepath.Join(parent, "public", "templates")); err == nil && st.IsDir() {
			return parent
		}
	}
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		log.Fatal("runtime.Caller failed")
	}
	return filepath.Clean(filepath.Join(filepath.Dir(file), ".."))
}

type previewRoute struct {
	Slug string // URL path segment after /public/ or /email/
	File string // filename under public/templates or email-templates
	Name string // {{ define "..." }} name
	Mock func(*I18n) any
}

var publicPreviews = []previewRoute{
	{Slug: "home", File: "home.html", Name: "home", Mock: mockPublicHome},
	{Slug: "archive", File: "archive.html", Name: "archive", Mock: mockPublicArchive},
	{Slug: "optin", File: "optin.html", Name: "optin", Mock: mockPublicOptin},
	{Slug: "message", File: "message.html", Name: "message", Mock: mockPublicMessage},
	{Slug: "subscription", File: "subscription.html", Name: "subscription", Mock: mockPublicSubscription},
	{Slug: "subscription-form", File: "subscription-form.html", Name: "subscription-form", Mock: mockPublicSubscriptionForm},
	{Slug: "login", File: "login.html", Name: "admin-login", Mock: mockAdminLogin},
	{Slug: "forgot-password", File: "forgot-password.html", Name: "admin-forgot-password", Mock: mockAdminForgot},
	{Slug: "reset-password", File: "reset-password.html", Name: "admin-reset-password", Mock: mockAdminReset},
	{Slug: "login-setup", File: "login-setup.html", Name: "admin-login-setup", Mock: mockAdminLoginSetup},
	{Slug: "twofa", File: "twofa.html", Name: "admin-twofa", Mock: mockAdminTwoFA},
}

var emailPreviews = []previewRoute{
	{Slug: "subscriber-optin", File: "subscriber-optin.html", Name: "subscriber-optin", Mock: mockEmailSubscriberOptin},
	{Slug: "optin-campaign", File: "subscriber-optin-campaign.html", Name: "optin-campaign", Mock: mockEmailOptinCampaign},
	{Slug: "campaign-status", File: "campaign-status.html", Name: "campaign-status", Mock: mockEmailCampaignStatus},
	{Slug: "import-status", File: "import-status.html", Name: "import-status", Mock: mockEmailImportStatus},
	{Slug: "subscriber-data", File: "subscriber-data.html", Name: "subscriber-data", Mock: mockEmailEmpty},
	{Slug: "smtp-test", File: "smtp-test.html", Name: "smtp-test", Mock: mockEmailEmpty},
	{Slug: "forgot-password", File: "forgot-password.html", Name: "forgot-password", Mock: mockEmailForgotPassword},
}

func parsePublicTemplates(root, pageFile string, lang *I18n) (*template.Template, error) {
	// Dotless {{ L.T "..." }} in index.html header/footer requires L in FuncMap (not only .L on data).
	fm := template.FuncMap{
		"L": func() *I18n { return lang },
	}
	base := filepath.Join(root, "public/templates/index.html")
	page := filepath.Join(root, "public/templates", pageFile)
	return template.New("").Funcs(fm).ParseFiles(base, page)
}

func parseEmailTemplates(root string, lang *I18n, pageFile string) (*template.Template, error) {
	fm := template.FuncMap{
		"L": func() *I18n { return lang },
		"RootURL": func() string {
			return "http://localhost:8088"
		},
		"LogoURL": func() string { return "" },
	}
	base := filepath.Join(root, "email-templates/base.html")
	page := filepath.Join(root, "email-templates", pageFile)
	return template.New("").Funcs(fm).ParseFiles(base, page)
}

func findRoute(list []previewRoute, slug string) (previewRoute, bool) {
	for _, r := range list {
		if r.Slug == slug {
			return r, true
		}
	}
	return previewRoute{}, false
}

// canonPath trims trailing slashes (except root).
func canonPath(p string) string {
	for len(p) > 1 && strings.HasSuffix(p, "/") {
		p = strings.TrimSuffix(p, "/")
	}
	return p
}

// listmonkRedirect maps listmonk-style URLs to this preview server's routes.
// Preserves query strings where relevant (e.g. /admin/reset?token=...).
func listmonkRedirect(path, rawQuery string) (target string, ok bool) {
	p := canonPath(path)
	withQuery := func(dst string) string {
		if rawQuery != "" {
			return dst + "?" + rawQuery
		}
		return dst
	}
	switch p {
	case "/subscription/form":
		return withQuery("/public/subscription-form"), true
	case "/archive":
		return withQuery("/public/archive"), true
	case "/archive.xml":
		return withQuery("/public/archive"), true
	case "/admin", "/admin/login":
		return withQuery("/public/login"), true
	case "/admin/forgot":
		return withQuery("/public/forgot-password"), true
	case "/admin/reset":
		return withQuery("/public/reset-password"), true
	case "/admin/login/twofa":
		return withQuery("/public/twofa"), true
	case "/auth/oidc":
		return withQuery("/public/login"), true
	default:
		return "", false
	}
}

// normalizePublicSlug maps paths used under /public/... to preview slugs.
func normalizePublicSlug(slug string) string {
	switch slug {
	case "subscription/form":
		return "subscription-form"
	case "archive.xml":
		return "archive"
	case "admin":
		return "login"
	default:
		return slug
	}
}

// optinCampaignSubFirstNameRE matches the literal placeholder listmonk leaves in
// subscriber-optin-campaign.html for a second compile pass when the fragment is embedded in a campaign.
var optinCampaignSubFirstNameRE = regexp.MustCompile(`\{\{\s*\.Subscriber\.FirstName\s*\}\}`)

func patchOptinCampaignEmailPreview(html string) string {
	return optinCampaignSubFirstNameRE.ReplaceAllString(html, previewMockSubscriberFirstName)
}

// isSafeLocaleCode restricts -lang to a basename under i18n/ (no path separators).
func isSafeLocaleCode(s string) bool {
	if len(s) < 2 || len(s) > 24 {
		return false
	}
	for _, r := range s {
		switch {
		case r >= 'a' && r <= 'z', r >= 'A' && r <= 'Z', r >= '0' && r <= '9', r == '-':
			continue
		default:
			return false
		}
	}
	return true
}

func main() {
	root := mustRoot()

	langFlag := flag.String("lang", "en", "i18n locale: basename of a JSON file under i18n/ (e.g. de, fr, zh-CN, cs-cz)")
	addrFlag := flag.String("addr", ":8088", "listen address")
	flag.Parse()
	code := strings.TrimSpace(*langFlag)
	if code == "" {
		code = "en"
	}
	if !isSafeLocaleCode(code) {
		log.Fatalf("invalid -lang %q: use only letters, digits, and hyphen (see i18n/*.json)", code)
	}
	langPath := filepath.Join(root, "i18n", code+".json")
	lang, err := LoadI18n(langPath)
	if err != nil {
		log.Fatalf("load i18n: %v", err)
	}
	log.Printf("i18n locale: %s", code)

	staticDir := filepath.Join(root, "public/static")
	fs := http.FileServer(http.Dir(staticDir))

	mux := http.NewServeMux()
	mux.Handle("/public/static/", http.StripPrefix("/public/static/", fs))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if dst, ok := listmonkRedirect(r.URL.Path, r.URL.RawQuery); ok {
			http.Redirect(w, r, dst, http.StatusFound)
			return
		}
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		var b strings.Builder
		b.WriteString("<!DOCTYPE html><html><head><meta charset=\"utf-8\"><title>listmonk template preview</title>")
		b.WriteString("<style>body{font-family:system-ui,sans-serif;max-width:56rem;margin:2rem auto;padding:0 1rem}code{background:#eee;padding:0 .2em}a{display:inline-block;margin:.25rem 0}</style>")
		b.WriteString("</head><body><h1>listmonk static template preview</h1>")
		b.WriteString("<p>Repo root: <code>")
		b.WriteString(template.HTMLEscapeString(root))
		b.WriteString("</code></p><h2>Public pages</h2><ul>")
		for _, p := range publicPreviews {
			fmt.Fprintf(&b, `<li><a href="/public/%s">%s</a> <small>(%s → %s)</small></li>`, p.Slug, p.Slug, p.File, p.Name)
		}
		b.WriteString("</ul><h2>System e-mails</h2><ul>")
		for _, p := range emailPreviews {
			fmt.Fprintf(&b, `<li><a href="/email/%s">%s</a> <small>(%s → %s)</small></li>`, p.Slug, p.Slug, p.File, p.Name)
		}
		b.WriteString("</ul><h2>Campaign templates</h2><ul>")
		for _, c := range campaignPreviews {
			fmt.Fprintf(&b, `<li><a href="/campaign/%s">%s</a> <small>(%s, %s)</small></li>`, c.Slug, c.Slug, c.File, c.Note)
		}
		b.WriteString(`<li><a href="/compare/default-visual">default-visual: .tpl and .json side by side</a></li>`)
		b.WriteString("</ul></body></html>")
		_, _ = w.Write([]byte(b.String()))
	})

	mux.HandleFunc("/public/", func(w http.ResponseWriter, r *http.Request) {
		slug := strings.TrimPrefix(r.URL.Path, "/public/")
		slug = strings.TrimSuffix(slug, "/")
		if slug == "" {
			http.NotFound(w, r)
			return
		}
		slug = normalizePublicSlug(slug)
		pr, ok := findRoute(publicPreviews, slug)
		if !ok {
			http.NotFound(w, r)
			return
		}
		tmpl, err := parsePublicTemplates(root, pr.File, lang)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := tmpl.ExecuteTemplate(w, pr.Name, pr.Mock(lang)); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	mux.HandleFunc("/email/", func(w http.ResponseWriter, r *http.Request) {
		slug := strings.TrimPrefix(r.URL.Path, "/email/")
		slug = strings.TrimSuffix(slug, "/")
		if slug == "" {
			http.NotFound(w, r)
			return
		}
		pr, ok := findRoute(emailPreviews, slug)
		if !ok {
			http.NotFound(w, r)
			return
		}
		tmpl, err := parseEmailTemplates(root, lang, pr.File)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if pr.Name == "optin-campaign" {
			var buf bytes.Buffer
			if err := tmpl.ExecuteTemplate(&buf, pr.Name, pr.Mock(lang)); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			_, _ = w.Write([]byte(patchOptinCampaignEmailPreview(buf.String())))
			return
		}
		if err := tmpl.ExecuteTemplate(w, pr.Name, pr.Mock(lang)); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	mux.HandleFunc("/campaign/", func(w http.ResponseWriter, r *http.Request) {
		c, ok := findCampaign(strings.Trim(strings.TrimPrefix(r.URL.Path, "/campaign/"), "/"))
		if !ok {
			http.NotFound(w, r)
			return
		}
		tmpl, err := parseCampaignTemplate(root, c, lang)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := tmpl.Execute(w, mockCampaign()); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// The visual template exists twice: the editor JSON and the HTML listmonk sends. Both should look the same.
	mux.HandleFunc("/compare/default-visual", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = w.Write([]byte(`<!DOCTYPE html><html><head><meta charset="utf-8"><title>default-visual: .tpl and .json</title>` +
			`<style>body{font-family:system-ui,sans-serif;margin:0}div{display:flex;gap:1rem;padding:1rem}section{flex:1}iframe{width:100%;height:90vh;border:1px solid #ccc}h2{font-size:1rem;margin:0 0 .5rem}</style></head><body><div>` +
			`<section><h2>default-visual.tpl (sent by listmonk)</h2><iframe src="/campaign/default-visual"></iframe></section>` +
			`<section><h2>default-visual.json (visual editor, approximate)</h2><iframe src="/campaign/default-visual-json"></iframe></section>` +
			`</div></body></html>`))
	})

	log.Printf("listmonk template preview listening on %s", *addrFlag)
	log.Fatal(http.ListenAndServe(*addrFlag, mux))
}
