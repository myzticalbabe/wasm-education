# Security Policy

## Reporting a vulnerability

Do **not** open a public issue for a security problem.

Use GitHub's private reporting: **Security** tab → **Report a vulnerability**.
If that is unavailable, contact [@myzticalbabe](https://github.com/myzticalbabe)
directly.

Include what you did, what happened, and what you expected. A proof of concept is
helpful but not required.

Expect an acknowledgement within 7 days. This is a single-maintainer project, so
please allow time before escalating.

## Supported versions

The live demo pages are served from `main` and are the supported version. There are
no maintained release branches. Security fixes land on `main` only.

## What this project handles

Knowing the threat model saves everyone time. The demo pages are static HTML with
no backend:

| Surface | Reality |
|---|---|
| Data collected | None. No accounts, no forms, no names. |
| Storage | One `localStorage` key, `brainExplorerProgress`, holding XP/badge state. Client-side only. |
| Network egress | None from the kids page. No analytics, no CDN, no third-party scripts, no telemetry. |
| Cookies | None. |
| Backend / API | None. |
| Secrets | None. There is nothing in this repository to leak. |

The realistic risk surface is therefore **supply chain and the WASM modules**, not
data breach:

1. **`rust-wasm/` and `go-wasm/`** — code compiled to WebAssembly and run in the
   learner's browser. A malicious or careless dependency here runs with the page's
   origin. Dependency review is in scope for reports.
2. **The CI workflows** — `.github/workflows/*.yml` run on push. Workflow-level
   issues (unpinned actions, injection via a branch name) are in scope.
3. **`innerHTML` in `demo-pages/`** — several render paths build markup from module
   state. If any of that state can be influenced by an attacker-controlled value,
   it becomes DOM XSS in a page used by children. In scope.

Out of scope: reports that require the user to have already modified their own
copy of the page, and reports about the GitHub Pages host itself.

## Hardening rules for contributors

- Pin GitHub Actions to a commit SHA, not a tag.
- Never interpolate an untrusted value into `run:` in a workflow.
- Prefer `textContent` over `innerHTML`; when `innerHTML` is unavoidable, escape.
- Do not add a dependency to `demo-pages/`. It has none today, and that is a
  security property, not an accident.
