# bypass-kit expectations (E2R-6: saturate the adversary, not the population)

`parse()` is total: it never traps, never returns empty. The bar is
no-crash + deterministic, not "correct interpretation" — triage output is
always re-verified downstream by chained seals.

| # | case | why it matters | expected |
|---|------|----------------|----------|
| 1 | empty | zero-length records are real (failed taps) | `parsed len=0 head=`, exit 0 |
| 2 | 10 KB single blob | volume triage must not OOM or truncate silently | `parsed len=10240`, exit 0, stable |
| 3 | cyrillic lookalikes (`раураl.сom`, `аt`) | homograph impersonation of senders/cells | echoed verbatim, stable — parser must not normalize away evidence |
| 4 | RTL override + combining marks | display-order attacks that flip meaning on screen | echoed verbatim, stable |
| 5 | embedded newlines | log-injection / line-splitting into ledger format | parse echoes; **ledger boundary is enforced by reporter line-splitting, and each line re-verifies** |
| 6 | delimiter soup | parser must not interpret `\|`, `:`, `,` structurally | echoed verbatim, stable |
| 7 | `prev=0 event=spoof seal=1` as parse input | spoofed ledger line fed as triage text | treated as inert text, stable |
| 8 | real event containing ` seal=` | proves `parse_line` right-splits the seal tag: `rsplit_once(" seal=")` takes the LAST tag, so event text cannot smuggle a seal | `claims=1 verified=1` |

Known non-goal: Unicode normalization (NFC/NFD) is NOT performed — two
byte-distinct spellings of one word seal differently. That is documented
behavior (seal the bytes you saw), not a finding, until a normalization
policy is declared in the WIT contract.
