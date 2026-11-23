#!/usr/bin/env bash
# replay.sh — end-to-end case replay for 13-forensics-hyperedge (Gap item 3).
# Re-derives every seal of case-001, fails closed on mismatch. Pinned values
# were recorded from live runs on 2026-11-22 (wasmtime 51.0.0); any change in
# guest behavior breaks this script by design.
set -euo pipefail

S="$(cd "$(dirname "$0")" && pwd)"
TC="$S/../_toolchain"
WR="${FORENSICS_WASMTIME:-$TC/wasmtime-v51.0.0-x86_64-windows/wasmtime.exe}"
[ -x "$WR" ] || { echo "wasmtime not found at $WR"; exit 2; }

pass=0
check() { # check <name> <got> <want>
  if [ "$2" == "$3" ]; then pass=$((pass+1)); echo "ok   $1";
  else echo "FAIL $1: got [$2] want [$3]"; exit 1; fi
}

E0='case-001|acquire|2026-11-22'
E1='case-001|parse|MSG from +1555 at cell-42'
E2='case-001|triage|cell-42 lead'
S0=2566118661192581389
S1=13517821046654747850
S2=16657771995194439954

# 1. g1 stability: the v0.1 seal value must never move.
check "seal-g1" "$("$WR" run --invoke "seal(\"$E0\")" "$S/custody.component.wasm")" "4526540249018236259"
# 2. Chain re-derivation.
check "chain-s0" "$("$WR" run --invoke "seal-chained(0, \"$E0\")" "$S/custody.component.wasm")" "$S0"
check "chain-s1" "$("$WR" run --invoke "seal-chained($S0, \"$E1\")" "$S/custody.component.wasm")" "$S1"
check "chain-s2" "$("$WR" run --invoke "seal-chained($S1, \"$E2\")" "$S/custody.component.wasm")" "$S2"
# 3. Positive verifications.
check "verify-s1" "$("$WR" run --invoke "verify-chained($S1, $S0, \"$E1\")" "$S/custody.component.wasm")" "true"
# 4. Negative controls: tampered event and tampered seal both fail.
check "tamper-event" "$("$WR" run --invoke "verify-chained($S1, $S0, \"$E1 TAMPERED\")" "$S/custody.component.wasm")" "false"
check "tamper-seal" "$("$WR" run --invoke "verify-chained(1, $S0, \"$E1\")" "$S/custody.component.wasm")" "false"
# 5. Ingest: single-chunk, multi-chunk, empty.
check "ingest-1" "$("$WR" run --invoke 'ingest("MSG from +1555 at cell-42")' "$S/acquire.component.wasm")" '"chunks=1 root=15411808922966539639 c0=631032329683602341"'
BLOB200=$(python3 -c "print('A'*200, end='')")
check "ingest-200" "$("$WR" run --invoke "ingest(\"$BLOB200\")" "$S/acquire.component.wasm")" '"chunks=4 root=12213112917550233364 c0=9172658670259733477 c1=9172658670259733477 c2=9172658670259733477 c3=3064734786292905933"'
check "ingest-empty" "$("$WR" run --invoke 'ingest("")' "$S/acquire.component.wasm")" '"chunks=0 root=14695981039346656037"'
# 6. Composed timeline + assist g1 + annotate pin.
check "timeline" "$("$WR" run --invoke 'run("MSG from +1555 at cell-42")' "$S/composed-timeline.component.wasm")" '"parsed len=25 head=MSG from +1555 at cell-4|timeline"'
check "assist-g1" "$("$WR" run --invoke 'assist("summarize cell-42")' "$S/aiassist.component.wasm")" '"model=pin:fnv-demo|prompt=summarize cell-42|logged=true"'
check "annotate" "$("$WR" run --invoke 'annotate("fnv-demo", "temp=0", "summarize cell-42")' "$S/aiassist.component.wasm")" '"model-hash=16113363829767271346|model-id=fnv-demo|prompt=summarize cell-42|logged=true"'
# 7. Full case report: 3 claims, 3 verified, pinned report hash.
LED="prev=0 event=$E0 seal=$S0\nprev=$S0 event=$E1 seal=$S1\nprev=$S1 event=$E2 seal=$S2"
TL='parsed len=25 head=MSG from +1555 at cell-4|timeline'
AN='model-hash=16113363829767271346|model-id=fnv-demo|prompt=summarize cell-42|logged=true'
REP="$("$WR" run --invoke "report(\"case-001\", \"$LED\", \"$TL\", \"$AN\")" "$S/composed-reporter.component.wasm")"
echo "$REP" | grep -q 'claims=3 verified=3' || { echo "FAIL report-counts: $REP"; exit 1; }
pass=$((pass+1)); echo "ok   report-counts"
echo "$REP" | grep -q 'report-hash=6535537349267962746' || { echo "FAIL report-hash: $REP"; exit 1; }
pass=$((pass+1)); echo "ok   report-hash"
# 8. Tampered ledger fails closed: the bad line is named, never dropped.
BAD_LED="prev=0 event=$E0 seal=$S0\nprev=$S0 event=$E1 seal=1\nprev=$S1 event=$E2 seal=$S2"
BAD="$("$WR" run --invoke "report(\"case-001\", \"$BAD_LED\", \"$TL\", \"$AN\")" "$S/composed-reporter.component.wasm")"
echo "$BAD" | grep -q 'claims=3 verified=2' || { echo "FAIL tamper-counts: $BAD"; exit 1; }
pass=$((pass+1)); echo "ok   tamper-counts"
echo "$BAD" | grep -q 'claim\[1\]=FAIL' || { echo "FAIL tamper-named: $BAD"; exit 1; }
pass=$((pass+1)); echo "ok   tamper-named"

echo "replay: $pass checks passed"
