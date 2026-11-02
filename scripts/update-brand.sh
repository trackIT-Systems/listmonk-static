#!/bin/sh
# Copies a release of trackIT-Systems/brand into this repo (ADR-0001).
# Usage: scripts/update-brand.sh 0.1.0
set -eu

version=${1:?usage: scripts/update-brand.sh VERSION}
root=$(cd "$(dirname "$0")/.." && pwd)
tmp=$(mktemp -d)
trap 'rm -rf "$tmp"' EXIT

curl -fsSL -o "$tmp/brand.tgz" \
  "https://github.com/trackIT-Systems/brand/releases/download/v$version/trackit-systems-brand-$version.tgz"
tar -xzf "$tmp/brand.tgz" -C "$tmp"
pkg="$tmp/package"

rm -rf "$root/brand"
mkdir -p "$root/brand" "$root/public/static/brand"
cp "$pkg/DESIGN.md" "$pkg/VOICE.md" "$pkg/dist/tokens.json" "$root/brand/"
cp -R "$pkg/glossary" "$root/brand/"
echo "$version" > "$root/brand/VERSION"
cp "$pkg/dist/tokens.css" "$root/public/static/brand/tokens.css"
cp "$pkg/assets/logo/trackit-logo-black.svg" "$root/public/static/logo.svg"

echo "brand $version copied. Run scripts/check-colors.sh and review the diff."
