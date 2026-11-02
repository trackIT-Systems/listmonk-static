package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// campaignPreview is a campaign template (email-templates/*.tpl or the visual editor JSON).
type campaignPreview struct {
	Slug string
	File string // filename under email-templates
	Note string // shown on the index page
}

var campaignPreviews = []campaignPreview{
	{Slug: "default", File: "default.tpl", Note: "campaign"},
	{Slug: "default-archive", File: "default-archive.tpl", Note: "archive view"},
	{Slug: "sample-tx", File: "sample-tx.tpl", Note: "transactional"},
	{Slug: "default-visual", File: "default-visual.tpl", Note: "visual, rendered HTML"},
	{Slug: "default-visual-json", File: "default-visual.json", Note: "visual, editor JSON rendered here"},
}

func findCampaign(slug string) (campaignPreview, bool) {
	for _, c := range campaignPreviews {
		if c.Slug == slug {
			return c, true
		}
	}
	return campaignPreview{}, false
}

// mockCampaignContent stands in for the campaign body that listmonk inserts at {{ template "content" . }}.
// It uses the elements the template styles: headings, links, a button, a list, a table and code.
const mockCampaignContent = `{{ define "content" }}
<h2>Hello {{ .Subscriber.FirstName }}</h2>
<p>This is mock campaign content from the preview renderer. Here is <a href="https://trackit.systems">a link</a> in running text.</p>
<p><a class="button" href="https://trackit.systems">Primary action</a></p>
<h3>A smaller heading</h3>
<ul>
	<li>First list item</li>
	<li>Second list item</li>
</ul>
<table>
	<tr><td><strong>Devices</strong></td><td>65</td></tr>
	<tr><td><strong>Period</strong></td><td>2026–2031</td></tr>
</table>
<pre>code block</pre>
{{ end }}`

type campaignCtx struct {
	Campaign struct {
		Subject string
	}
	Subscriber struct {
		Name      string
		FirstName string
		LastName  string
		Email     string
	}
	Tx struct {
		Data map[string]any
	}
}

func mockCampaign() campaignCtx {
	var c campaignCtx
	c.Campaign.Subject = "Mock campaign subject"
	c.Subscriber.Name = previewMockSubscriberFirstName + " Example"
	c.Subscriber.FirstName = previewMockSubscriberFirstName
	c.Subscriber.LastName = "Example"
	c.Subscriber.Email = "alex@example.com"
	c.Tx.Data = map[string]any{"order_id": "1234", "shipping_date": "2026-11-02"}
	return c
}

func campaignFuncs(lang *I18n) template.FuncMap {
	return template.FuncMap{
		"L":              func() *I18n { return lang },
		"UnsubscribeURL": func() string { return "/public/subscription" },
		"MessageURL":     func() string { return "/campaign/default" },
		"TrackView":      func() template.HTML { return "" },
	}
}

// parseCampaignTemplate returns the template for a campaign preview. The JSON source is
// first turned into HTML by renderVisualJSON; both then run through html/template like listmonk does.
func parseCampaignTemplate(root string, c campaignPreview, lang *I18n) (*template.Template, error) {
	b, err := os.ReadFile(filepath.Join(root, "email-templates", c.File))
	if err != nil {
		return nil, err
	}
	src := string(b)
	if strings.HasSuffix(c.File, ".json") {
		if src, err = renderVisualJSON(b); err != nil {
			return nil, fmt.Errorf("%s: %w", c.File, err)
		}
	}
	t, err := template.New(c.File).Funcs(campaignFuncs(lang)).Parse(src)
	if err != nil {
		return nil, err
	}
	if t.Lookup("content") == nil {
		if _, err := t.Parse(mockCampaignContent); err != nil {
			return nil, err
		}
	}
	return t, nil
}

// --- Visual editor JSON ---
//
// renderVisualJSON approximates the HTML that listmonk's visual editor (email-builder-js)
// produces for the block types used in default-visual.json. listmonk sends the saved HTML
// (default-visual.tpl), not this; the preview shows whether the JSON still matches it.

type visualBlock struct {
	Type string `json:"type"`
	Data struct {
		Style       visualStyle     `json:"style"`
		Props       json.RawMessage `json:"props"`
		ChildrenIDs []string        `json:"childrenIds"`

		// EmailLayout keeps its settings directly in data.
		BackdropColor string `json:"backdropColor"`
		CanvasColor   string `json:"canvasColor"`
		TextColor     string `json:"textColor"`
		FontFamily    string `json:"fontFamily"`
		BorderColor   string `json:"borderColor"`
		BorderRadius  *int   `json:"borderRadius"`
	} `json:"data"`
}

type visualStyle struct {
	Color           string `json:"color"`
	BackgroundColor string `json:"backgroundColor"`
	FontSize        int    `json:"fontSize"`
	FontWeight      string `json:"fontWeight"`
	TextAlign       string `json:"textAlign"`
	Padding         *struct {
		Top, Bottom, Right, Left int
	} `json:"padding"`
}

type visualProps struct {
	Text                  string `json:"text"`
	Level                 string `json:"level"`
	LineColor             string `json:"lineColor"`
	LineHeight            int    `json:"lineHeight"`
	ButtonBackgroundColor string `json:"buttonBackgroundColor"`
	ButtonTextColor       string `json:"buttonTextColor"`
	ButtonStyle           string `json:"buttonStyle"`
	URL                   string `json:"url"`
}

var visualFonts = map[string]string{
	"MODERN_SANS":    `"Helvetica Neue", "Arial Nova", "Nimbus Sans", Arial, sans-serif`,
	"BOOK_SANS":      `Optima, Candara, "Noto Sans", source-sans-pro, sans-serif`,
	"ORGANIC_SANS":   `Seravek, "Gill Sans Nova", Ubuntu, Calibri, "DejaVu Sans", source-sans-pro, sans-serif`,
	"GEOMETRIC_SANS": `Avenir, "Avenir Next LT Pro", Montserrat, Corbel, "URW Gothic", source-sans-pro, sans-serif`,
	"HEAVY_SANS":     `Bahnschrift, "DIN Alternate", "Franklin Gothic Medium", "Nimbus Sans Narrow", sans-serif-condensed, sans-serif`,
	"ROUNDED_SANS":   `ui-rounded, "Hiragino Maru Gothic ProN", Quicksand, Comfortaa, Manjari, "Arial Rounded MT Bold", Calibri, source-sans-pro, sans-serif`,
	"MODERN_SERIF":   `Charter, "Bitstream Charter", "Sitka Text", Cambria, serif`,
	"BOOK_SERIF":     `"Iowan Old Style", "Palatino Linotype", "URW Palladio L", P052, serif`,
	"MONOSPACE":      `"Nimbus Mono PS", "Courier New", "Cutive Mono", monospace`,
}

var headingSizes = map[string]int{"h1": 32, "h2": 24, "h3": 20}

var buttonRadius = map[string]string{"rounded": "4px", "pill": "64px", "rectangle": "0"}

func (s visualStyle) css() string {
	var b strings.Builder
	if s.Color != "" {
		fmt.Fprintf(&b, "color:%s;", s.Color)
	}
	if s.BackgroundColor != "" {
		fmt.Fprintf(&b, "background-color:%s;", s.BackgroundColor)
	}
	if s.FontSize != 0 {
		fmt.Fprintf(&b, "font-size:%dpx;", s.FontSize)
	}
	if s.FontWeight != "" {
		fmt.Fprintf(&b, "font-weight:%s;", s.FontWeight)
	}
	if s.TextAlign != "" {
		fmt.Fprintf(&b, "text-align:%s;", s.TextAlign)
	}
	b.WriteString(s.padding())
	return b.String()
}

func (s visualStyle) padding() string {
	p := s.Padding
	if p == nil {
		return ""
	}
	return fmt.Sprintf("padding:%dpx %dpx %dpx %dpx;", p.Top, p.Right, p.Bottom, p.Left)
}

var markdownLinkRE = regexp.MustCompile(`\[([^\]]+)\]\(([^)\s]+)\)`)

// markdownToHTML handles what the default template uses: paragraphs, links and inline HTML.
func markdownToHTML(s string) string {
	var b strings.Builder
	for _, p := range strings.Split(s, "\n\n") {
		p = markdownLinkRE.ReplaceAllString(p, `<a href="$2" target="_blank">$1</a>`)
		fmt.Fprintf(&b, "<p>%s</p>", p)
	}
	return b.String()
}

func renderVisualJSON(src []byte) (string, error) {
	var doc map[string]visualBlock
	if err := json.Unmarshal(src, &doc); err != nil {
		return "", err
	}
	root, ok := doc["root"]
	if !ok || root.Type != "EmailLayout" {
		return "", fmt.Errorf(`no "root" EmailLayout block`)
	}
	l := root.Data

	var b strings.Builder
	b.WriteString(`<!doctype html><html><head><meta name="viewport" content="width=device-width, initial-scale=1" /></head><body>`)
	fmt.Fprintf(&b, `<div style='background-color:%s;color:%s;font-family:%s;font-size:16px;font-weight:400;letter-spacing:0.15008px;line-height:1.5;margin:0;padding:32px 0;min-height:100%%;width:100%%'>`,
		l.BackdropColor, l.TextColor, visualFonts[l.FontFamily])
	canvas := fmt.Sprintf("margin:0 auto;max-width:600px;background-color:%s", l.CanvasColor)
	if l.BorderColor != "" {
		canvas += ";border:1px solid " + l.BorderColor
	}
	if l.BorderRadius != nil {
		canvas += fmt.Sprintf(";border-radius:%dpx;overflow:hidden", *l.BorderRadius)
	}
	fmt.Fprintf(&b, `<table align="center" width="100%%" style="%s" role="presentation" cellspacing="0" cellpadding="0" border="0"><tbody><tr style="width:100%%"><td>`, canvas)

	for _, id := range l.ChildrenIDs {
		blk, ok := doc[id]
		if !ok {
			return "", fmt.Errorf("block %q is listed in root.childrenIds but missing", id)
		}
		var p visualProps
		if len(blk.Data.Props) > 0 {
			if err := json.Unmarshal(blk.Data.Props, &p); err != nil {
				return "", fmt.Errorf("block %q: %w", id, err)
			}
		}
		st := blk.Data.Style
		switch blk.Type {
		case "Heading":
			level := p.Level
			if level == "" {
				level = "h2"
			}
			pad := st.padding()
			st.Padding, st.FontWeight = nil, ""
			fmt.Fprintf(&b, `<%s style="%sfont-weight:bold;margin:0;font-size:%dpx;%s">%s</%s>`, level, st.css(), headingSizes[level], pad, p.Text, level)
		case "Text":
			fmt.Fprintf(&b, `<div style="%s">%s</div>`, st.css(), markdownToHTML(p.Text))
		case "Divider":
			h := p.LineHeight
			if h == 0 {
				h = 1
			}
			fmt.Fprintf(&b, `<div style="%s"><hr style="width:100%%;border:none;border-top:%dpx solid %s;margin:0"/></div>`, st.css(), h, p.LineColor)
		case "Button":
			fg := p.ButtonTextColor
			if fg == "" {
				fg = "#FFFFFF"
			}
			style := p.ButtonStyle
			if style == "" {
				style = "rounded"
			}
			fmt.Fprintf(&b, `<div style="%s"><a href="%s" target="_blank" style="color:%s;font-size:16px;font-weight:bold;background-color:%s;border-radius:%s;display:inline-block;padding:12px 20px;text-decoration:none">%s</a></div>`,
				st.css(), p.URL, fg, p.ButtonBackgroundColor, buttonRadius[style], p.Text)
		default:
			fmt.Fprintf(&b, `<div style="padding:16px 24px;border:1px dashed red">Block type %q is not supported by the preview renderer.</div>`, blk.Type)
		}
	}
	b.WriteString(`</td></tr></tbody></table></div></body></html>`)
	return b.String(), nil
}
