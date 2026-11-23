//! Acquisition ingest — the front door of the pipeline (Gap 1 + Gap 2).
//! Pure compute: splits the blob into 64-byte chunks, FNV-1a hashes each,
//! folds a root seal over the chunk hashes. No imports, no WASI.

wit_bindgen::generate!({
    world: "acquire",
    path: "../wit",
});

struct Acquire;

const CHUNK: usize = 64;

fn fnv1a_bytes(bytes: &[u8]) -> u64 {
    let mut h: u64 = 0xcbf29ce484222325;
    for b in bytes {
        h ^= *b as u64;
        h = h.wrapping_mul(0x100000001b3);
    }
    h
}

impl exports::forensics::pipe::acquire_ops::Guest for Acquire {
    fn ingest(blob: String) -> String {
        let bytes = blob.as_bytes();
        let mut out = String::new();
        let mut root: u64 = 0xcbf29ce484222325;
        let mut n = 0u64;
        for chunk in bytes.chunks(CHUNK) {
            let h = fnv1a_bytes(chunk);
            // Fold: mix chunk hash into the running root.
            for b in h.to_le_bytes() {
                root ^= b as u64;
                root = root.wrapping_mul(0x100000001b3);
            }
            if n > 0 {
                out.push(' ');
            }
            out.push_str(&format!("c{}={}", n, h));
            n += 1;
        }
        if n == 0 {
            format!("chunks=0 root={}", root)
        } else {
            format!("chunks={} root={} {}", n, root, out)
        }
    }
}

export!(Acquire);
