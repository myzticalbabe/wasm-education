#!/usr/bin/env python3
"""Structural checks for the static demo pages.

Guards against three defect classes that have actually occurred in this repo:

1. A nav button pointing at a view that does not exist. `data-view="modules"`
   with no `id="view-modules"` leaves the page blank when clicked.
2. A local asset reference (`href`/`src`) that 404s, e.g. a manifest or icon
   that was never added.
3. A `<canvas>` with no text alternative, which is invisible to assistive tech.

Stdlib only, no network, no dependencies. Exit 1 on any finding.
"""
from __future__ import annotations

import html.parser
import pathlib
import re
import sys

ROOT = pathlib.Path(__file__).resolve().parent.parent
PAGES = sorted((ROOT / "demo-pages").glob("*.html"))

# Views that are served by another section by design.
VIEW_ALIASES = {"modules": "home"}

# Nav targets intentionally resolved in JavaScript rather than by a matching
# section. showView() in demo-pages/index.html redirects "modules" to the first
# module card, which then calls showView() again with a real view name.
JS_RESOLVED = {"modules"}

findings: list[str] = []


class PageParser(html.parser.HTMLParser):
    def __init__(self) -> None:
        super().__init__(convert_charrefs=True)
        self.view_targets: list[str] = []
        self.view_ids: set[str] = set()
        self.refs: list[tuple[str, str]] = []
        self.canvases: list[dict[str, str]] = []

    def handle_starttag(self, tag: str, attrs: list[tuple[str, str | None]]) -> None:
        a = {k: (v or "") for k, v in attrs}

        if a.get("data-view"):
            self.view_targets.append(a["data-view"])
        if a.get("id", "").startswith("view-"):
            self.view_ids.add(a["id"][len("view-"):])

        for attr in ("href", "src"):
            if a.get(attr):
                self.refs.append((attr, a[attr]))

        if tag == "canvas":
            self.canvases.append(a)


def check_page(page: pathlib.Path) -> None:
    rel = page.relative_to(ROOT).as_posix()
    parser = PageParser()
    parser.feed(page.read_text(encoding="utf-8"))

    # 1. nav targets resolve to a section (directly or via an alias).
    for target in parser.view_targets:
        if target in parser.view_ids:
            continue
        if VIEW_ALIASES.get(target) in parser.view_ids:
            continue
        if target in JS_RESOLVED:
            continue
        findings.append(
            f"{rel}: nav target data-view=\"{target}\" has no "
            f"id=\"view-{target}\", no alias in VIEW_ALIASES, and is not in "
            "JS_RESOLVED — clicking it will blank the page"
        )

    # 2. local references exist on disk.
    for attr, value in parser.refs:
        if re.match(r"^(https?:|mailto:|tel:|data:|javascript:|#|//)", value):
            continue
        target = (page.parent / value.split("#")[0].split("?")[0]).resolve()
        if not target.exists():
            findings.append(f"{rel}: {attr}=\"{value}\" does not resolve to a file")

    # 3. canvases have a text alternative and an exposed role.
    for canvas in parser.canvases:
        label = canvas.get("aria-label") or canvas.get("aria-labelledby")
        if not label:
            findings.append(
                f"{rel}: <canvas id=\"{canvas.get('id', '?')}\"> has no aria-label"
            )
        if canvas.get("role") not in {"img", "application"}:
            findings.append(
                f"{rel}: <canvas id=\"{canvas.get('id', '?')}\"> has no "
                "role=\"img\" (or \"application\" with a keyboard model)"
            )


def main() -> int:
    if not PAGES:
        print("no pages found under demo-pages/", file=sys.stderr)
        return 1

    for page in PAGES:
        check_page(page)

    if findings:
        print(f"FAIL — {len(findings)} structural finding(s):\n")
        for f in findings:
            print(f"  - {f}")
        return 1

    print(f"PASS — {len(PAGES)} page(s) checked ({', '.join(p.name for p in PAGES)})")
    return 0


if __name__ == "__main__":
    sys.exit(main())
