# BUILD_LOG 13-forensics-hyperedge — 2026-11-22T18:46:50Z

## versions
wasm-tools 1.263.0 (7a2fa422c 2026-10-30)
wasmtime 51.0.0 (32656dc48 2026-11-20)
cargo 1.99.1 (5f94df478 2026-08-27)
rustc 1.99.1 (58cf8f9c3 2026-10-13)

## artifacts
-rw-r--r-- 1 clint 197609 23643 Nov 22 13:46 C:/Users/clint/Desktop/WASM_SAMPLES/13-forensics-hyperedge/aiassist.component.wasm
-rw-r--r-- 1 clint 197609 51846 Nov 22 13:46 C:/Users/clint/Desktop/WASM_SAMPLES/13-forensics-hyperedge/composed-timeline.component.wasm
-rw-r--r-- 1 clint 197609 22216 Nov 22 13:46 C:/Users/clint/Desktop/WASM_SAMPLES/13-forensics-hyperedge/custody.component.wasm
-rw-r--r-- 1 clint 197609 27052 Nov 22 13:46 C:/Users/clint/Desktop/WASM_SAMPLES/13-forensics-hyperedge/parser.component.wasm
-rw-r--r-- 1 clint 197609 27052 Nov 22 13:46 C:/Users/clint/Desktop/WASM_SAMPLES/13-forensics-hyperedge/parserdep.wasm
-rw-r--r-- 1 clint 197609 24675 Nov 22 13:46 C:/Users/clint/Desktop/WASM_SAMPLES/13-forensics-hyperedge/timeline.component.wasm

## wit composed
package root:component;

world root {
  export run: func(blob: string) -> string;
}

# BUILD_LOG v0.2 — 2026-11-22 (gap closure: acquire, chained seals, replay, reporter, CI, bypass-kit, annotate, demo)
# wit-bindgen =0.63.0 pinned in all 6 crates; CARGO_TARGET_DIR=.target-shared; invoke = call-expression form.
# g1: seal("case-001|acquire|2026-11-22") = 4526540249018236259 (stable)
# chain: 2566118661192581389 -> 13517821046654747850 -> 16657771995194439954
# report(case-001): claims=3 verified=3 report-hash=6535537349267962746; tampered: verified=2 claim[1]=FAIL
# replay.sh 17/17; bypass-kit/run.sh 8/8; validate VALID-OK x8 (acquire custody parser timeline reporter aiassist composed-timeline composed-reporter)
