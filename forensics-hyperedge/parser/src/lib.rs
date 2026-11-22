//! Triage parser — normalizes one artifact blob (Gap 1 + Gap 6).
//! Pure compute: deterministic prefix + length tag. No imports.

wit_bindgen::generate!({
    world: "parser",
    path: "../wit",
});

struct Parser;

impl exports::forensics::pipe::parse_ops::Guest for Parser {
    fn parse(blob: String) -> String {
        let head: String = blob.chars().take(24).collect();
        format!("parsed len={} head={}", blob.len(), head)
    }
}

export!(Parser);
