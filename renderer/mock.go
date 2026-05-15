package main

import (
	"database/sql"
	"html/template"
	"time"
)

// previewMockSubscriberFirstName is used for e-mail previews (including optin-campaign placeholder substitution).
const previewMockSubscriberFirstName = "Alex"

type publicMeta struct {
	Title       string
	Description string
}

type homePageData struct {
	publicMeta
}

type subscriptionPageData struct {
	publicMeta
	ShowManage       bool
	AllowBlocklist   bool
	AllowPreferences bool
	Subscriber       struct {
		Name string
	}
	Subscriptions []subscriptionRow
	AllowExport     bool
	AllowWipe       bool
	SubUUID         string
}

type subscriptionRow struct {
	UUID                 string
	Name                 string
	SubscriptionStatus struct{ Value string }
}

type subscriptionFormData struct {
	publicMeta
	Lists   []publicListRow
	Captcha captchaBlock
}

type publicListRow struct {
	UUID        string
	Name        string
	Description string
	Type        string
}

type captchaBlock struct {
	Enabled  bool
	Provider string
	Key      string
}

type optinPageData struct {
	publicMeta
	Lists []publicListRow
}

type messagePageData struct {
	publicMeta
	Message template.HTML
}

type archivePageData struct {
	publicMeta
	Campaigns  []archiveCampaignRow
	TotalPages int
	Pagination template.HTML
}

type archiveCampaignRow struct {
	URL       string
	Subject   string
	SendAt    sql.NullTime
	CreatedAt sql.NullTime
}

type adminLoginData struct {
	publicMeta
	PasswordEnabled  bool
	Nonce            string
	NextURI          string
	OIDCProvider     string
	OIDCProviderLogo string
	Error            string
}

type adminForgotData struct {
	publicMeta
	Error string
}

type adminResetData struct {
	publicMeta
	Token string
	Email string
	Error string
}

type adminLoginSetupData struct {
	publicMeta
	Nonce string
	Error string
}

type adminTwoFAData struct {
	publicMeta
	Token   string
	NextURI string
	Error   string
}

// PageCtx is the root context for public HTML templates (matches listmonk).
type PageCtx struct {
	L                   *I18n
	SiteName            string
	RootURL             string
	AssetVersion        string
	LogoURL             string
	FaviconURL          string
	EnablePublicArchive bool
	EnablePublicSubPage bool
	Data                any
}

// --- Email template contexts ---

type emailSubscriberOptinCtx struct {
	Subscriber struct {
		FirstName string
	}
	Lists []struct {
		Name string
		UUID string
		Type string
	}
	OptinURL string
	UnsubURL string
}

type emailOptinCampaignCtx struct {
	Lists []struct {
		Name string
		Type string
	}
	OptinURLAttr template.HTMLAttr
}

type emailImportStatusCtx struct {
	Name      string
	Status    string
	Imported  int
	Total     int
}

type emailForgotPasswordCtx struct {
	ResetURL string
}

func newPageCtx(lang *I18n, data any) PageCtx {
	return PageCtx{
		L:                   lang,
		SiteName:            "Mock Newsletter",
		RootURL:             "http://localhost:8088",
		AssetVersion:        "preview",
		LogoURL:             "",
		FaviconURL:          "",
		EnablePublicArchive: true,
		EnablePublicSubPage: true,
		Data:                data,
	}
}

func mockPublicHome(lang *I18n) any {
	return newPageCtx(lang, homePageData{
		publicMeta: publicMeta{
			Title:       "Home",
			Description: "Public landing page preview",
		},
	})
}

func mockPublicSubscription(lang *I18n) any {
	return newPageCtx(lang, subscriptionPageData{
		publicMeta: publicMeta{
			Title:       "Subscription",
			Description: "Manage subscription and privacy",
		},
		ShowManage:       true,
		AllowBlocklist:   true,
		AllowPreferences: true,
		Subscriber: struct {
			Name string
		}{Name: "Jane Subscriber"},
		Subscriptions: []subscriptionRow{
			{UUID: "11111111-1111-1111-1111-111111111111", Name: "Weekly digest", SubscriptionStatus: struct{ Value string }{Value: "subscribed"}},
			{UUID: "22222222-2222-2222-2222-222222222222", Name: "Product updates", SubscriptionStatus: struct{ Value string }{Value: "subscribed"}},
		},
		AllowExport: true,
		AllowWipe:   true,
		SubUUID:     "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee",
	})
}

func mockPublicSubscriptionForm(lang *I18n) any {
	return newPageCtx(lang, subscriptionFormData{
		publicMeta: publicMeta{
			Title:       "Subscribe",
			Description: "Public subscription form",
		},
		Lists: []publicListRow{
			{UUID: "33333333-3333-3333-3333-333333333333", Name: "Announcements", Description: "Company news and releases.", Type: "public"},
			{UUID: "44444444-4444-4444-4444-444444444444", Name: "Beta program", Description: "", Type: "public"},
		},
		Captcha: captchaBlock{Enabled: false},
	})
}

func mockPublicOptin(lang *I18n) any {
	return newPageCtx(lang, optinPageData{
		publicMeta: publicMeta{
			Title:       "Confirm subscription",
			Description: "Double opt-in confirmation",
		},
		Lists: []publicListRow{
			{UUID: "55555555-5555-5555-5555-555555555555", Name: "Newsletter", Description: "", Type: "public"},
			{UUID: "66666666-6666-6666-6666-666666666666", Name: "VIP list", Description: "", Type: "private"},
		},
	})
}

func mockPublicMessage(lang *I18n) any {
	return newPageCtx(lang, messagePageData{
		publicMeta: publicMeta{
			Title:       "Message",
			Description: "Generic status message",
		},
		Message: template.HTML("<p>Your subscription preferences were updated.</p>"),
	})
}

func mockPublicArchive(lang *I18n) any {
	t1 := time.Date(2025, 3, 1, 12, 0, 0, 0, time.UTC)
	t2 := time.Date(2025, 4, 10, 9, 30, 0, 0, time.UTC)
	return newPageCtx(lang, archivePageData{
		publicMeta: publicMeta{
			Title:       "Archive",
			Description: "Public campaign archive",
		},
		Campaigns: []archiveCampaignRow{
			{URL: "http://localhost:8088/archive/campaign-1", Subject: "March updates", SendAt: sql.NullTime{Valid: true, Time: t1}, CreatedAt: sql.NullTime{}},
			{URL: "http://localhost:8088/archive/campaign-2", Subject: "April product news", SendAt: sql.NullTime{Valid: false}, CreatedAt: sql.NullTime{Valid: true, Time: t2}},
		},
		TotalPages: 2,
		Pagination: template.HTML(`<a href="?page=1">1</a> · <a href="?page=2">2</a>`),
	})
}

func mockAdminLogin(lang *I18n) any {
	return newPageCtx(lang, adminLoginData{
		publicMeta: publicMeta{
			Title:       "Login",
			Description: "Admin login",
		},
		PasswordEnabled:  true,
		Nonce:            "mock-nonce",
		NextURI:          "/admin",
		OIDCProvider:     "Google",
		OIDCProviderLogo: "google.com.png",
		Error:            "",
	})
}

func mockAdminForgot(lang *I18n) any {
	return newPageCtx(lang, adminForgotData{
		publicMeta: publicMeta{
			Title:       "Forgot password",
			Description: "Password recovery",
		},
		Error: "",
	})
}

func mockAdminReset(lang *I18n) any {
	return newPageCtx(lang, adminResetData{
		publicMeta: publicMeta{
			Title:       "Reset password",
			Description: "Set a new password",
		},
		Token: "mock-reset-token",
		Email: "admin@example.com",
		Error: "",
	})
}

func mockAdminLoginSetup(lang *I18n) any {
	return newPageCtx(lang, adminLoginSetupData{
		publicMeta: publicMeta{
			Title:       "Setup",
			Description: "First-time admin account",
		},
		Nonce: "mock-nonce", Error: "",
	})
}

func mockAdminTwoFA(lang *I18n) any {
	return newPageCtx(lang, adminTwoFAData{
		publicMeta: publicMeta{
			Title:       "Two-factor authentication",
			Description: "Enter TOTP code",
		},
		Token:   "mock-twofa-token",
		NextURI: "/admin",
		Error:   "",
	})
}

func mockEmailSubscriberOptin(lang *I18n) any {
	_ = lang
	return emailSubscriberOptinCtx{
		Subscriber: struct{ FirstName string }{FirstName: previewMockSubscriberFirstName},
		Lists: []struct {
			Name string
			UUID string
			Type string
		}{
			{Name: "Newsletter", UUID: "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa", Type: "public"},
			{Name: "Insiders", UUID: "bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb", Type: "private"},
		},
		OptinURL: "https://example.com/confirm?token=mock",
		UnsubURL: "https://example.com/subscription/mock-token",
	}
}

func mockEmailOptinCampaign(lang *I18n) any {
	_ = lang
	return emailOptinCampaignCtx{
		Lists: []struct {
			Name string
			Type string
		}{
			{Name: "Newsletter", Type: "public"},
			{Name: "Secret club", Type: "private"},
		},
		OptinURLAttr: template.HTMLAttr(`href="https://example.com/confirm-campaign"`),
	}
}

func mockEmailCampaignStatus(lang *I18n) any {
	_ = lang
	return map[string]any{
		"ID":     42,
		"Name":   "Spring sale",
		"Status": "running",
		"Sent":   1200,
		"ToSend": 5000,
		"Reason": "SMTP rate limit hit (mock)",
	}
}

func mockEmailImportStatus(lang *I18n) any {
	_ = lang
	return emailImportStatusCtx{
		Name:     "subscribers-2025-05.csv",
		Status:   "finished",
		Imported: 950,
		Total:    1000,
	}
}

func mockEmailForgotPassword(lang *I18n) any {
	_ = lang
	return emailForgotPasswordCtx{
		ResetURL: "https://example.com/admin/reset?token=mock",
	}
}

func mockEmailEmpty(lang *I18n) any {
	_ = lang
	return struct{}{}
}
