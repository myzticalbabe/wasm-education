//! Timeline builder — IMPORTS parse-ops, exports run (composition demo).
//! run(blob) = parse(blob) + "|timeline". Parser supplied at compose time.

wit_bindgen::generate!({
    world: "timeline",
    path: "../wit",
});

struct Timeline;

impl Guest for Timeline {
    fn run(blob: String) -> String {
        let p = forensics::pipe::parse_ops::parse(&blob);
        format!("{}|timeline", p)
    }
}

export!(Timeline);
