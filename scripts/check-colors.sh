#!/bin/sh
# Fails if a template or stylesheet uses a hex color outside the brand palette.
# E-mails cannot use CSS variables, so they inline the palette values; this
# check keeps them in sync with brand/tokens.json.
set -eu

root=$(cd "$(dirname "$0")/.." && pwd)
cd "$root"

palette=$(sed -n '/"palette"/,/}/p' brand/tokens.json | grep -o '#[0-9a-fA-F]\{6\}' | tr 'A-F' 'a-f')

status=0
for f in $(find public email-templates -type f \( -name '*.html' -o -name '*.tpl' -o -name '*.json' -o -name '*.css' \) ! -path 'public/static/brand/*'); do
  for c in $(grep -o '#[0-9a-fA-F]\{3,8\}\b' "$f" | tr 'A-F' 'a-f' | sort -u); do
    if ! echo "$palette" | grep -qx "$c"; then
      echo "$f: $c is not a brand palette color (use the 6-digit value from brand/tokens.json)"
      status=1
    fi
  done
done
exit $status
