# Accessibility — current state and what is not done

Measured against `demo-pages/kids-index.html` on 2026-11-15. Written down so the
claims in the README are auditable and so regressions are visible.

## What is implemented

| Feature | Where | Notes |
|---|---|---|
| Landmarks | `<header>`, `<nav>`, `<main>`, `<footer>` | one of each, in order |
| Heading order | h1 → h2 → h3 → h4 | no skips (measured: 1 / 8 / 12 / 20) |
| Canvas text alternatives | `aria-label` + `role="img"` + fallback text | see below |
| Keyboard path to every module action | toolbar `<button>` per module | the canvas is never required |
| Reduced motion | `@media (prefers-reduced-motion: reduce)` | collapses all animation to 0.01 ms |
| High contrast | `@media (prefers-contrast: high)` | forces black/white palette |
| Dark mode | `@media (prefers-color-scheme: dark)` | separate palette, not an inversion |
| Base text size | `html { font-size: 18px }` | sized for the age band |
| Selection state exposed | `aria-pressed` on nav buttons | toggled by `showView()` |
| Spoken instructions | speech-synthesis toggle in the header | for reading support, not a screen-reader replacement |

## Canvas handling

The four visualisations are `<canvas>`, which exposes nothing to assistive
technology on its own. Each one carries:

1. `role="img"` and an `aria-label` naming what it shows and how it responds;
2. fallback text inside the element for renderers that ignore ARIA;
3. a `<p class="sr-only">` summary of the current state, updated by the module;
4. an equivalent set of toolbar buttons that perform every action the canvas
   accepts by click.

Point 4 is the important one. `role="img"` on an interactive surface is a
compromise: it tells a screen reader "this is a picture, here is its description,
don't try to tab into it". That is only honest because the buttons provide the
same affordances. If you add a canvas interaction with no button equivalent, the
canvas role must change to `role="application"` and a keyboard model must be
documented.

## Known limitations

- **Not tested with a real screen reader.** The ARIA is correct by inspection and
  by automated checks; no NVDA/JAWS/VoiceOver session has been run. Treat as
  unverified until someone does.
- **No automated axe/Lighthouse gate in CI.** CI checks that the pages parse and
  that the WASM builds; it does not check accessibility. Adding `axe-core` as a
  CI step is the obvious next improvement.
- **Colour contrast** is chosen for the age band and passes for body text, but the
  decorative gradient text in headers (`.hero h2`, `.logo-text h1`) uses
  `-webkit-text-fill-color: transparent` over a gradient. It is decorative and
  duplicated by the plain `<title>`/heading text, but a strict audit will flag it.
- **Motion:** several elements animate infinitely (`bounce`, `float`). The
  reduced-motion query neutralises them, but the default is animated.

## If you are extending this

1. Any new interactive surface needs a non-pointer path.
2. Any new colour needs checking against both the light and dark palette.
3. Any new animation needs to survive `prefers-reduced-motion`.
4. Do not add `tabindex` to make a `<div>` clickable. Use a `<button>`.
