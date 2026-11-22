//! Custody ledger — append-only seal/verify, hash-chained (Gap 2 + Gap 8).
//! Pure compute: FNV-1a u64 over event bytes. No imports, no WASI.
//! seal()/verify() are g1-stable: byte-identical outputs to v0.1 forever,
//! so every previously recorded seal still verifies.

wit_bindgen::generate!({
    world: "custody",
    path: "../wit",
});

struct Custody;

fn fnv1a(s: &str) -> u64 {
    let mut h: u64 = 0xcbf29ce484222325;
    for b in s.bytes() {
        h ^= b as u64;
        h = h.wrapping_mul(0x100000001b3);
    }
    h
}

impl exports::forensics::pipe::custody_ops::Guest for Custody {
    fn seal(event: String) -> u64 {
        fnv1a(&event)
    }
    fn verify(expected: u64, event: String) -> bool {
        fnv1a(&event) == expected
    }
    fn seal_chained(prev: u64, event: String) -> u64 {
        fnv1a(&format!("{}:{}", prev, event))
    }
    fn verify_chained(expected: u64, prev: u64, event: String) -> bool {
        Self::seal_chained(prev, event) == expected
    }
}

export!(Custody);
