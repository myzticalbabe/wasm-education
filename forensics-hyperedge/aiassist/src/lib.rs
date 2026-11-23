//! AI-assist wrapper — ONLY logged path to model output (Gap 8).
//! assist() is g1-stable (v0.1 output preserved). annotate() pins the model
//! identity: model-hash = FNV-1a over "<model-id>|<params>", so a changed
//! model or params invalidates old annotations by construction.

wit_bindgen::generate!({
    world: "aiassist",
    path: "../wit",
});

struct Ai;

impl Guest for Ai {
    fn assist(prompt: String) -> String {
        format!("model=pin:fnv-demo|prompt={}|logged=true", prompt)
    }
    fn annotate(model_id: String, params: String, prompt: String) -> String {
        let mut h: u64 = 0xcbf29ce484222325;
        for b in format!("{}|{}", model_id, params).bytes() {
            h ^= b as u64;
            h = h.wrapping_mul(0x100000001b3);
        }
        format!(
            "model-hash={}|model-id={}|prompt={}|logged=true",
            h, model_id, prompt
        )
    }
}

export!(Ai);
