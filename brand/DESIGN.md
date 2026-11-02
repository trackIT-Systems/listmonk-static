# Design

Visual identity of trackIT Systems. Applies to every public surface: websites, datasheets, slides, flyers, newsletters.

Status: **draft**, transcribed from Nicole's [styleguide v2 working draft](docs/sources/2026-10-styleguide-v2-working-draft.pdf), which is based on the new flyer. Values live in [tokens/](tokens/), files in [assets/](assets/); this document explains how to use them.

## Logo

Full lockup (mark plus "trackIT systems") and the mark alone; files and where each fits in [assets/README.md](assets/README.md).

- Black logo on light backgrounds, white logo on dark or green backgrounds.
- Circle versions for avatars and profile images.

TODO(Nicole): clear space, minimum size, logo on slate-green and dark-green, misuse examples, logo color (`#000` in files vs. `#222222` in the styleguide).

## Color

| Token | Hex | Role |
|---|---|---|
| `slate-green` (Schiefergrün) | `#5A7A6B` | Large content boxes and cards. Text inside is white. |
| `dark-green` (Dunkelgrün) | `#1E5D39` | Headings and running text on light backgrounds; product badges; secondary actions. |
| `mint` (Hellgrün) | `#CBD6CD` | Borders and subtle bases, e.g. table headers. |
| `signal-orange` | `#FF6B35` | Primary call to action, and nothing else. |
| `dark-gray` (Dunkelgrau) | `#222222` | Body text, neutral buttons, logo. |
| `soft-white` | `#F5F5F5` | Page background, print and digital. |

Rules:

- Signal orange marks the one primary action in view. Large areas and decorative cards stay green or neutral.
- Dark-green is for type and small elements: headings, links, product badges, secondary buttons. Large areas use slate-green (content boxes, page headers, bands) or mint (`surface-subtle`: table headers, highlighted rows). Otherwise dark-green is everywhere and stops pointing at anything.
- Text on slate-green and dark-green is white (`white`, not `soft-white`: soft-white on slate-green is 4.35:1 and fails for body text). Dark-green text on slate-green is unreadable (1.65:1); dark-gray reaches 3.36:1, large text only.
- Secondary text (captions, sources, metadata) is dark-gray in a smaller size. No other palette color reaches 4.5:1 on soft-white; slate-green is 4.35:1.
- Links are dark-green and always underlined, so they differ from headings in the same color.
- Panels on the page background (tables, forms) are white (`surface`). Subtle bases inside them are mint (`surface-subtle`), with dark-gray or dark-green text.
- Mint borders are decorative. Boundaries of inputs and other controls are slate-green (`border-control`), because WCAG 1.4.11 needs 3:1 and mint on soft-white is 1.37:1.
- Keyboard focus is a dark-green ring (`focus-ring`) on light backgrounds and a white ring on slate-green and dark-green.

Which text color goes on which background is defined by the roles in [tokens/color.tokens.json](tokens/color.tokens.json): `text`, `heading` and `link` on `page-background`, `surface` and `surface-subtle`, and each `on-X` on its `X` (e.g. `on-cta` on `cta`). Other pairs are not approved, even when their contrast passes.

### Contrast (WCAG 2.1, computed)

| Text on background | Ratio | Result |
|---|---|---|
| white on slate-green | 4.74:1 | AA |
| white on dark-green | 7.84:1 | AAA |
| dark-green on soft-white | 7.19:1 | AAA |
| dark-gray on soft-white | 14.59:1 | AAA |
| dark-green on mint | 5.24:1 | AA |
| slate-green on soft-white | 4.35:1 | large text only |
| soft-white on slate-green | 4.35:1 | large text only; use white |
| dark-gray on slate-green | 3.36:1 | large text only; use white |
| **white on signal-orange** | **2.84:1** | **fails**, also for large bold text (needs 3:1) |
| dark-gray on signal-orange | 5.61:1 | AA |

Text on signal-orange is always dark-gray, never white ([ADR-0007](docs/adr/0007-dark-text-on-signal-orange.md)). The working draft listed white on orange as 3.0:1 / pass; the measured value is 2.84:1.

## Shape

- Rounded corners throughout: 16–24px on boxes and cards (`radius.card`, `radius.control`).
- Badges and tags are pills (`radius.pill`), e.g. the dark-green product badge.
- Images are cropped round, with a slate-green ring.
- Sharp-cornered containers do not mix with rounded ones.

## Buttons

In order of emphasis:

| Button | Fill and label | Where |
|---|---|---|
| **Primary CTA** | signal-orange, dark-gray label | Any background. One per view. |
| **Secondary action** | dark-green, white label | Light backgrounds. |
| **Container action** | slate-green, white label | Inside white panels and directly below or next to a slate-green card, for actions that belong to that content. Never on slate-green itself, where it disappears. |
| **Outline** | dark-green border and label, no fill | Light backgrounds. On slate-green and dark-green: white border and label. |

On a slate-green card, use the primary CTA or the white outline button.

Shape and size: `radius.control` (16px), at least 44px high (a touch target), label 15px semibold in the text typeface, padding 11px × 22px. One size; add a second only when a real layout needs it.

States, without extra colors:

- **Hover**: the fill moves 15% away from its label color: signal-orange mixes with white, the greens mix with dark-gray. Contrast to the label only goes up. Outline buttons get a mint fill (on dark: white at 15%).
- **Pressed**: the same, at 30%. A pressed outline button fills dark-green with a white label (on dark: white at 30%).
- **Focus**: the `focus-ring`, 3px with a 2px gap; white on slate-green and dark-green.
- **Disabled**: mint fill, slate-green label, no hover. WCAG exempts disabled controls from contrast, but on public pages prefer an enabled button that explains what is missing.

## Typography

### trackIT Systems: our display font, selected cases

Our display font is **trackIT Systems**, a font of our own in two styles ([ADR-0009](docs/adr/0009-own-font-trackit-systems.md)). It is a display face, used only in selected cases, not for running text. Files and license: [assets/fonts/trackit-systems/](assets/fonts/trackit-systems/).

| Token | Style | File | Weight | Matches |
|---|---|---|---|---|
| `font.family.display` + `font.weight.display-bold` | trackIT Systems Bold | `trackIT Systems Bold` | 700 | "trackIT" in the wordmark |
| `font.family.display` + `font.weight.display-regular` | trackIT Systems Regular | `trackIT Systems Regular` | 400 | "systems" in the wordmark |

The letters follow the outlined wordmark:

- Capitals, figures and cap-height symbols are 596 units tall (`I` and `T` in the wordmark).
- Ascenders (`b d f h k l`) are 600 units tall (`k` in the wordmark); the `t` is 606.
- The `r` ends in a vertical cut and kerns tight against round letters (`ra` in the wordmark).

Typeset text does not replace the wordmark: the logo has its own letter-spacing (about 0.028em in "trackIT", 0.04em in "systems") and a tighter `IT`. Always use the logo files in [assets/logo/](assets/logo/).

TODO(Nicole): list the cases (e.g. hero headlines, product names, flyer titles?).

### Text typeface

TODO(Nicole): typeface for body text, UI and data (with tabular numerals), type scale, weights. The working draft does not specify one yet.

## Layout and spacing

TODO: spacing scale, grid, breakpoints, maximum content width.

## Imagery

Round crops with a slate-green ring (see Shape).

TODO: photography style (field work, devices in place, real sites), what we avoid, image credits.

## Figures and diagrams

Interactive and static figures explain our systems and how results are handled.

TODO: color use in charts (categorical, sequential), line weights, labeling, a static fallback for every interactive figure, reduced-motion behavior.

## Icons

[Tabler Icons](https://tabler.io/icons), outline style ([ADR-0006](docs/adr/0006-tabler-icons.md)).

- Stroke: 2px, the Tabler default. Custom icons use the same stroke.
- Color: inherits the text color (`currentColor`); white on slate-green and dark-green.
- Which icon means what: [assets/icons/icons.yml](assets/icons/icons.yml). Add the concept there before using a new icon.
- Filled variants only where outline is unreadable (very small sizes, active states).

TODO(Nicole): sizes (e.g. 16 / 20 / 24px).

## Motion

TODO: durations, easing, when motion is allowed. Respect `prefers-reduced-motion`.
