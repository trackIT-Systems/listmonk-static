<!doctype html>
<html>
    <head>
        <title>{{ .Campaign.Subject }}</title>
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, minimum-scale=1">
        <base target="_blank">
        <style>
            body {
                background-color: #f5f5f5;
                font-family: system-ui, -apple-system, 'Segoe UI', sans-serif;
                font-size: 15px;
                line-height: 26px;
                margin: 0;
                color: #222222;
            }

            pre {
                background: #f5f5f5;
                border-radius: 16px;
                padding: 2px;
            }

            table {
                width: 100%;
                border: 1px solid #cbd6cd;
            }
            table td {
                border-color: #cbd6cd;
                padding: 5px;
            }

            .wrap {
                background-color: #ffffff;
                border: 1px solid #cbd6cd;
                padding: 30px;
                max-width: 525px;
                margin: 0 auto;
                border-radius: 24px;
            }

            /* Primary action: signal-orange, dark-gray label (brand DESIGN.md, Buttons) */
            .button {
                background: #ff6b35;
                color: #222222 !important;
                display: inline-block;
                border-radius: 16px;
                padding: 11px 22px;
                font-size: 15px;
                font-weight: 600;
                line-height: 22px;
                text-align: center;
                text-decoration: none !important;
            }

            .footer {
                text-align: center;
                font-size: 12px;
                color: #222222;
            }
                .footer a {
                    color: #222222;
                    margin-right: 5px;
                }

            .gutter {
                padding: 30px;
            }

            h1, h2, h3, h4 {
                color: #1e5d39;
            }

            img {
                max-width: 100%;
                height: auto;
            }

            a {
                color: #1e5d39;
                text-decoration: underline;
            }
            @media screen and (max-width: 600px) {
                .wrap {
                    max-width: auto;
                }
                .gutter {
                    padding: 10px;
                }
            }
        </style>
    </head>
<body style="background-color: #f5f5f5;font-family: system-ui, -apple-system, 'Segoe UI', sans-serif;font-size: 15px;line-height: 26px;margin: 0;color: #222222;">
    <div class="gutter" style="padding: 30px;">&nbsp;</div>
    <div class="wrap" style="background-color: #ffffff;padding: 30px;max-width: 525px;margin: 0 auto;border: 1px solid #cbd6cd;border-radius: 24px;">
        {{ template "content" . }}
    </div>
    
    <div class="footer" style="text-align: center;font-size: 12px;color: #222222;">
        <p style="margin: 28px 0 0 0; color: #222222; line-height: 1.5;">
            <strong>trackIT Systems GmbH</strong><br>
            Zelterstraße 1, 35043 Marburg-Schröck, Germany
        </p>
        <p style="margin: 14px 0 0 0;">
            {{ L.T "public.poweredBy" }} <a href="https://listmonk.app" target="_blank" rel="noreferrer" style="color: #222222;">listmonk</a>
            &nbsp;·&nbsp;
            <a href="https://trackit.systems/en/contact/" style="color: #222222;" target="_blank" rel="noreferrer">Imprint</a>
            &nbsp;·&nbsp;
            <a href="https://trackit.systems/datenschutzerklaerung/" style="color: #222222;" target="_blank" rel="noreferrer">Privacy Policy</a>
        </p>
    </div>
</body>
</html>
