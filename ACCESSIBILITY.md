# Accessibility

Accessibility is a condition of this platform being usable in a classroom, not a
nice-to-have. This page states what we aim for, what has actually been measured, what we
already know is missing, and how to report a barrier.

The detailed audit — measured against `demo-pages/kids-index.html` on 2026-11-15 — is in
[`docs/ACCESSIBILITY.md`](docs/ACCESSIBILITY.md). Read it before running a session; it
distinguishes what is implemented from what is still unverified.

## Target

We work toward **WCAG 2.2 Level AA** for anything a person reads or operates. That is a
goal that guides the work; it is not a claim of verified conformance, and this page does
not describe it as one.

## Priorities

- **Keyboard first.** Every module action is reachable from a toolbar `<button>`; the
  canvas is never required. Selection state is exposed with `aria-pressed`.
- **Screen reader support.** Landmarks are used once each and in order (`<header>`,
  `<nav>`, `<main>`, `<footer>`), and the heading order runs h1 → h2 → h3 → h4 with no
  level skipped.
- **Plain language.** Base text is 18px, sized for the age band; terms are defined
  before use (see the vocabulary list in [STEM.md](STEM.md)).
- **Text before pictures.** A canvas is never the only carrier of information.

## How the four canvases are handled

Each visualisation carries four things, and the fourth is the important one:

1. `role="img"` plus an `aria-label` naming what it shows and how it responds;
2. fallback text inside the element, for renderers that ignore ARIA;
3. a `<p class="sr-only">` summary of current state, updated by the module;
4. an equivalent toolbar button for every action the canvas accepts by click.

`role="img"` on an interactive surface is a compromise — it tells a screen reader "this
is a picture, here is its description, don't try to tab into it". That is only honest
because point 4 provides the same affordances. **If you add a canvas interaction with no
button equivalent, the canvas role must change to `role="application"` and the keyboard
model must be documented.**

## Preferences honoured

| Preference | Effect |
|---|---|
| `prefers-reduced-motion: reduce` | All animation collapses to 0.01 ms |
| `prefers-contrast: high` | Forces a black/white palette |
| `prefers-color-scheme: dark` | Separate dark palette, not an inversion |

There is also a speech-synthesis toggle in the header that reads instructions aloud.
That is a reading-support aid, not a replacement for a screen reader.

## What has actually been tested

- **Measured by inspection and automated checks:** landmarks, heading order, canvas
  text alternatives, keyboard path to every module action, the three media queries
  above, `aria-pressed` state, base text size.
- **Not tested with a real screen reader.** No NVDA/JAWS/VoiceOver session has been run.
  The ARIA is correct by inspection; treat screen-reader behaviour as unverified until
  someone does.
- **No automated axe/Lighthouse gate in CI.** CI checks that the pages parse and that
  the WASM builds. It does not check accessibility.

## Known barriers

- **Decorative gradient text.** `.hero h2` and `.logo-text h1` use
  `-webkit-text-fill-color: transparent` over a gradient. It duplicates the plain
  heading text, but a strict audit will flag it.
- **Animated by default.** Several elements animate infinitely (`bounce`, `float`). The
  reduced-motion query neutralises them, but the default is motion.

## Reporting a barrier

Open an issue at <https://github.com/myzticalbabe/wasm-education/issues/new> and include:

- your operating system and browser;
- the assistive technology you were using, if any;
- the step you were on, and what you expected instead;
- whether the barrier blocks the task or only slows it down.

An accessibility issue is not a cosmetic issue. A keyboard trap that blocks a module is
treated with the same priority as a crash.

## For contributors

- Any new interactive surface needs a non-pointer path.
- Any new colour needs checking against both the light and dark palette.
- Any new animation needs to survive `prefers-reduced-motion`.
- Do not add `tabindex` to make a `<div>` clickable. Use a `<button>`.
- Do not "fix" an accessibility report by removing the feature.
