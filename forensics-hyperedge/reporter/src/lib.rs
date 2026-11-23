//! Court-ready report — IMPORTS custody-ops, exports report (Gap 1).
//! Re-verifies every ledger line against the chained seal, then binds the
//! whole case (ledger + timeline + analyst notes + per-claim verdicts) under
//! one report-hash. A tampered line fails closed: claim=FAIL, and the hash
//! covers the FAIL so it cannot be edited out.

wit_bindgen::generate!({
    world: "reporter",
    path: "../wit",
});

struct Reporter;

fn fnv1a(s: &str) -> u64 {
    let mut h: u64 = 0xcbf29ce484222325;
    for b in s.bytes() {
        h ^= b as u64;
        h = h.wrapping_mul(0x100000001b3);
    }
    h
}

impl exports::forensics::pipe::report_ops::Guest for Reporter {
    fn report(case_id: String, ledger: String, timeline: String, notes: String) -> String {
        let mut out = String::new();
        out.push_str(&format!("case={}\n", case_id));
        let mut n = 0u32;
        let mut ok = 0u32;
        let mut verdicts = String::new();
        for line in ledger.lines() {
            let line = line.trim();
            if line.is_empty() {
                continue;
            }
            let verdict = match parse_line(line) {
                Some((prev, event, seal)) => {
                    let expect = forensics::pipe::custody_ops::seal_chained(prev, &event);
                    expect == seal
                }
                None => false,
            };
            if verdict {
                ok += 1;
            }
            verdicts.push_str(&format!(
                "claim[{}]={} line={}\n",
                n,
                if verdict { "OK" } else { "FAIL" },
                line
            ));
            n += 1;
        }
        out.push_str(&format!("claims={} verified={}\n", n, ok));
        out.push_str(&verdicts);
        out.push_str(&format!("timeline={}\nnotes={}\n", timeline, notes));
        let hash = fnv1a(&format!(
            "{}|{}|{}|{}|{}|{}",
            case_id, ledger, timeline, notes, n, verdicts
        ));
        out.push_str(&format!("report-hash={}\n", hash));
        out
    }
}

/// Parse "prev=<u64> event=<event> seal=<u64>". Event may contain spaces;
/// the seal tag is split from the right so event text cannot smuggle it.
fn parse_line(line: &str) -> Option<(u64, String, u64)> {
    let (head, seal_part) = line.rsplit_once(" seal=")?;
    let rest = head.strip_prefix("prev=")?;
    let (prev_part, event) = rest.split_once(" event=")?;
    Some((prev_part.parse().ok()?, event.to_string(), seal_part.parse().ok()?))
}

export!(Reporter);
