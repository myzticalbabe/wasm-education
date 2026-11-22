# 13 — Forensics Hyper-Edge (LE / oversight pipeline, Component Model)

**Type:** WASI-0.2-style pure-compute components — WIT-typed, composed, capability-denied.

Maps to Europol Oct-2026 gaps: standard interchange (1), custody-by-construction (2),
volume triage (6), logged AI (8).

## Contract

`wit/forensics.wit` (`forensics:pipe@0.1.0`):

- `custody-ops`: `seal(event: string) -> u64` (FNV-1a), `verify(expected, event) -> bool`
- `parse-ops`: `parse(blob: string) -> string`
- worlds `custody`, `parser`, `timeline` (imports `parse-ops`, exports `run`), `aiassist`

## Build (proven 2026-11-22)

```bash
ROOT=13-forensics-hyperedge
WT=../_toolchain/wasm-tools-1.263.0-x86_64-windows/wasm-tools.exe
for c in custody parser timeline aiassist; do
  (cd $ROOT/$c && cargo build --release --target wasm32-unknown-unknown)
  $WT component new $ROOT/$c/target/wasm32-unknown-unknown/release/$c.wasm -o $ROOT/$c.component.wasm
done
cp $ROOT/parser.component.wasm $ROOT/parserdep.wasm
$WT compose $ROOT/timeline.component.wasm -d $ROOT/parserdep.wasm -o $ROOT/composed-timeline.component.wasm
```

## Proofs (wasmtime 51.0.0, wasm-tools 1.263.0)

- `seal("case-001|acquire|2026-11-22")` → `4526540249018236259`
- `run("MSG from +1555 at cell-42")` → `"parsed len=25 head=MSG from +1555 at cell-4|timeline"`
- `assist("summarize device-2 messages")` → `"model=pin:fnv-demo|prompt=summarize device-2 messages|logged=true"`
- `component wit composed-timeline.component.wasm` → `world root { export run }` (`parse-ops` internalized — composition proof)
- `validate --features component-model` → VALID-OK (both)

## Gap mapping

- B custody ledger → Gap 2 (chain of custody born at acquisition, verifiable)
- A acquire world pattern → Gap 1 (interchangeable imagers behind one WIT)
- C triage compose (parser→timeline) → Gaps 1+6 (new artifact parser = one component, no glue)
- F aiassist wrapper → Gap 8 (only logged, pinned, replayable path to model output)
- Same WIT reused for legislative `oversight-production` twin (document-intake, deposition-index, entity-link).
