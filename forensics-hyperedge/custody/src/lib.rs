//! Custody ledger — append-only seal/verify (Gap 2 + Gap 8).
//! Pure compute: FNV-1a u64 over event bytes. No imports, no WASI.

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
}

export!(Custody);
