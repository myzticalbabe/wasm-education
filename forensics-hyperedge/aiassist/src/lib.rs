//! AI-assist wrapper — ONLY logged path to model output (Gap 8).
//! Pure compute demo: returns pinned digest + prompt echo so every
//! AI-derived lead is replayable. Real host would enforce log() per call.

wit_bindgen::generate!({
    world: "aiassist",
    path: "../wit",
});

struct Ai;

impl Guest for Ai {
    fn assist(prompt: String) -> String {
        format!("model=pin:fnv-demo|prompt={}|logged=true", prompt)
    }
}

export!(Ai);
