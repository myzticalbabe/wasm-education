# Contributing

Thanks for looking. This repository is an education platform, so the bar for a
change is "does this help a learner understand something", not "does this add a
feature".

## Ways to contribute that are genuinely useful

1. **Run a session and report what broke.** Facilitator reports are the most
   valuable contribution this project can receive. Use the
   [facilitator guide](docs/FACILITATOR-GUIDE.md), then open an issue with what
   confused learners.
2. **Standards corrections.** If a mapping in
   [`docs/STANDARDS-ALIGNMENT.md`](docs/STANDARDS-ALIGNMENT.md) is wrong for your
   jurisdiction, say which code and why.
3. **Accessibility testing with real assistive technology.** See the limitations
   list in [`docs/ACCESSIBILITY.md`](docs/ACCESSIBILITY.md) — no screen-reader
   session has been run yet.
4. **WASM performance or correctness** in `rust-wasm/` and `go-wasm/`.
5. **New educational modules** that fit the existing semantic-annotation pattern.

## Development setup

```bash
# Rust → WASM
cd rust-wasm && wasm-pack build --target web --out-dir pkg

# Go → WASM
cd go-wasm && tinygo build -o main.wasm -target wasm -no-debug -opt=s .

# Demo pages (no build step)
cd demo-pages && npx http-server .
```

Prerequisites: Rust 1.75+ with `wasm32-unknown-unknown`, TinyGo 0.31+, Node 20+.

## Non-negotiables

These are the properties that make the resource usable by the people it is for.
A pull request that breaks one will be asked to change:

- **No network calls from the kids page.** Progress is one `localStorage` key.
  No analytics, no fonts from a CDN, no telemetry.
- **No accounts, no names, no demographic prompts.** The audience includes minors.
- **No leaderboards or public comparison.** Badges mark exploration, not rank.
- **No pointer-only interactions.** Every canvas action needs a keyboard path.
- **`prefers-reduced-motion` and `prefers-contrast` must keep working.**
- **Copy stays growth-framed.** Effort and strategy are the mechanism; "smart" is
  not an explanation.

## Style

- Vanilla JS, no framework, no bundler. The pages must run from `file://`.
- ES2020 is fine. Do not introduce a build step for `demo-pages/`.
- Match the existing CSS custom-property structure rather than adding literals.
- Prefer editing the four existing modules over adding a fifth one that overlaps.

## Pull requests

- One concern per PR.
- Say which module and which standard the change serves.
- If you changed a page, include before/after screenshots for both the light and
  dark palettes.
- If you changed WASM code, run the CI workflow locally first (`ci.yml` is short).

## Commit messages

Present tense, imperative, one line of what and one line of why:

```
Fix canvas fallback text so screen readers announce module state

The four canvases had aria-label but no role and no fallback
content, so a renderer that ignored ARIA got an empty element.
```

## License

Contributions are accepted under the [SOUL FORGE License](LICENSE).
