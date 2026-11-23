#!/usr/bin/env bash
# bypass-kit/run.sh — adversarial inputs against parse() + reporter (Gap item 6).
# Bar: every case exits 0 (no trap/panic) and is deterministic across runs.
# Case 8 proves the ledger parser resists " seal=" smuggling via right-split.
set -euo pipefail

S="$(cd "$(dirname "$0")/.." && pwd)"
WR="${FORENSICS_WASMTIME:-$S/../_toolchain/wasmtime-v51.0.0-x86_64-windows/wasmtime.exe}"
[ -x "$WR" ] || { echo "wasmtime not found at $WR"; exit 2; }

pass=0
stable() { # stable <name> <invoke-arg>
  local a b
  a=$("$WR" run --invoke "parse(\"$2\")" "$S/parser.component.wasm")
  b=$("$WR" run --invoke "parse(\"$2\")" "$S/parser.component.wasm")
  if [ "$a" == "$b" ] && [ -n "$a" ]; then pass=$((pass+1)); echo "ok   $1 -> $a";
  else echo "FAIL $1: [$a] vs [$b]"; exit 1; fi
}

stable "empty" ''
BIG=$(python3 -c "print('Z'*10240, end='')")
stable "10KB" "$BIG"
stable "unicode-lookalike" 'раураl.сom from +1555 аt cell-42'
stable "combining-rtl" 'cell-42 ‮tuo-yalp esrever combininǵ'
stable "newlines" 'line1\nline2\nline3'
stable "delimiter-soup" '|||:,,,|||:::,,,|||'
stable "seal-smuggle-parse" 'prev=0 event=spoof seal=1'

# Case 8: a real ledger event containing " seal=" must still verify (right-split).
EV='intake note seal=999 decoy'
SEAL=$("$WR" run --invoke "seal-chained(0, \"$EV\")" "$S/custody.component.wasm")
LED="prev=0 event=$EV seal=$SEAL"
REP=$("$WR" run --invoke "report(\"case-smuggle\", \"$LED\", \"t\", \"n\")" "$S/composed-reporter.component.wasm")
echo "$REP" | grep -q 'claims=1 verified=1' || { echo "FAIL smuggle-ledger: $REP"; exit 1; }
pass=$((pass+1)); echo "ok   smuggle-ledger -> claims=1 verified=1"

echo "bypass-kit: $pass checks passed"
