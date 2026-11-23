# 13 — Forensics Hyper-Edge (LE / oversight pipeline, Component Model)

**Type:** WASI-0.2-style pure-compute components — WIT-typed, composed, capability-denied.

Maps to Europol Oct-2026 gaps: standard interchange (1), custody-by-construction (2),
volume triage (6), logged AI (8).

## Contract

`wit/forensics.wit` (`forensics:pipe@0.1.0`):

- `custody-ops`: `seal`/`verify` (FNV-1a, g1-stable) + `seal-chained(prev, event)` / `verify-chained` (v0.2)
- `parse-ops`: `parse(blob: string) -> string`
- `acquire-ops` (v0.2): `ingest(blob) -> "chunks=N root=R c0=H0 ..."` (64-byte chunks, folded root)
- `report-ops` (v0.2): `report(case-id, ledger, timeline, notes)` — per-claim OK/FAIL + `report-hash`
- worlds `custody`, `parser`, `acquire`, `timeline` (imports `parse-ops`, exports `run`),
  `reporter` (imports `custody-ops`, exports `report-ops`), `aiassist` (`assist` g1-stable + v0.2 `annotate(model-id, params, prompt)`)

## Build (proven 2026-11-22)

```bash
ROOT=13-forensics-hyperedge
WT=../_toolchain/wasm-tools-1.263.0-x86_64-windows/wasm-tools.exe
# wit-bindgen pinned to =0.63.0 in every crate (see Cargo.toml) so CI rebuilds
# byte-identical guests; CARGO_TARGET_DIR shares one wit-bindgen build.
export CARGO_TARGET_DIR=$ROOT/.target-shared
for c in acquire custody parser timeline reporter aiassist; do
  (cd $ROOT/$c && cargo build --locked --release --target wasm32-unknown-unknown)
  $WT component new $CARGO_TARGET_DIR/wasm32-unknown-unknown/release/$c.wasm -o $ROOT/$c.component.wasm
done
cp $ROOT/parser.component.wasm $ROOT/parserdep.wasm
cp $ROOT/custody.component.wasm $ROOT/custodydep.wasm
$WT compose $ROOT/timeline.component.wasm -d $ROOT/parserdep.wasm -o $ROOT/composed-timeline.component.wasm
$WT compose $ROOT/reporter.component.wasm -d $ROOT/custodydep.wasm -o $ROOT/composed-reporter.component.wasm
bash $ROOT/replay.sh            # 17 checks, pinned expectations
bash $ROOT/bypass-kit/run.sh    # 8 adversarial checks
```

Invoke syntax (wasmtime 51): call-expression form —
`wasmtime run --invoke 'seal("event")' custody.component.wasm`.
Multi-line strings use `\n` escapes inside the invoke literal.

## Proofs (wasmtime 51.0.0, wasm-tools 1.263.0)

- `seal("case-001|acquire|2026-11-22")` → `4526540249018236259`
- `run("MSG from +1555 at cell-42")` → `"parsed len=25 head=MSG from +1555 at cell-4|timeline"`
- `assist("summarize device-2 messages")` → `"model=pin:fnv-demo|prompt=summarize device-2 messages|logged=true"`
- `component wit composed-timeline.component.wasm` → `world root { export run }` (`parse-ops` internalized — composition proof)
- `validate --features component-model` → VALID-OK (both)

## Proofs v0.2 (2026-11-22, same toolchain)

- g1: `seal("case-001|acquire|2026-11-22")` → `4526540249018236259` (unchanged)
- `ingest("MSG from +1555 at cell-42")` → `"chunks=1 root=15411808922966539639 c0=631032329683602341"`
- chain: `seal-chained(0, E0)` → `2566118661192581389` → `13517821046654747850` → `16657771995194439954`
- `annotate("fnv-demo", "temp=0", "summarize cell-42")` → `"model-hash=16113363829767271346|..."`
- `report(case-001, ledger, timeline, notes)` → `claims=3 verified=3`, `report-hash=6535537349267962746`
- tampered ledger → `claims=3 verified=2`, `claim[1]=FAIL` (fails closed, named)
- bypass-kit: empty / 10KB / homographs / RTL+combining / newlines / delimiter-soup / `seal=` smuggle — all exit 0 + deterministic; smuggled event verifies via right-split
- `replay.sh` 17/17, `bypass-kit/run.sh` 8/8, all 8 components `validate` VALID-OK
- CI job `forensics-components` rebuilds + replays + budgets every component (≤128 KiB) on push

## Gap mapping

- B custody ledger → Gap 2 (chain of custody born at acquisition, verifiable)
- A acquire world pattern → Gap 1 (interchangeable imagers behind one WIT)
- C triage compose (parser→timeline) → Gaps 1+6 (new artifact parser = one component, no glue)
- F aiassist wrapper → Gap 8 (only logged, pinned, replayable path to model output)
- v0.2 closes the loop: acquire front door, chained seals, replay fixture, court-ready report, CI, bypass-kit, model pinning, demo page
- Same WIT reused for legislative `oversight-production` twin (document-intake, deposition-index, entity-link).
