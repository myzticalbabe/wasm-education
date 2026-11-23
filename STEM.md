# STEM learning guide — wasm-education

A WebAssembly-powered learning platform with two tracks: four game-shaped adventures
for ages 8–13, and an undergraduate/developer platform that opens the same models up
for inspection. This guide is for teachers, librarians, club and troop leaders, parents
and students who want to use it for real STEM learning — in particular for girls and
other groups still under-represented in computing, engineering and mathematics.

It says what the platform actually teaches, how to run a session, and what to prepare
so that nobody is excluded by accident. Deeper material lives in
[`docs/FACILITATOR-GUIDE.md`](docs/FACILITATOR-GUIDE.md),
[`docs/STANDARDS-ALIGNMENT.md`](docs/STANDARDS-ALIGNMENT.md) and
[`docs/INCLUSION-GIRLS-IN-STEM.md`](docs/INCLUSION-GIRLS-IN-STEM.md).

## What this project teaches

**Subject areas**

- **Computational neuroscience** — neuron types, membrane dynamics, spiking, and why a
  network does not behave like the sum of its cells.
- **Graph theory and algorithms** — shortest path, spanning trees, strongly connected
  components, and how a search actually proceeds.
- **Cognitive psychology** — semantic memory, spreading activation, priming, schema
  acquisition.
- **Educational psychology** — cognitive load theory: intrinsic, extraneous and
  germane load, and what "this feels hard" is telling you.

**Skills practised**

- Building and revising a model, then defending it with evidence from the screen.
- Predicting before testing — turning a demo into an experiment.
- Reading a visualisation (raster plot, network diagram, load meter) and separating
  signal from decoration.
- Comparing two algorithms on the same input and justifying a choice.
- Decomposing "why do I remember some words faster?" into testable sub-questions.
- Explaining a result in one written sentence — which is the actual assessment.

## Prerequisites

For the kids track: none. Ages 8–13, no accounts, no installs, no prior programming,
and no internet connection after the first page load. A browser and a pointing device
(or a keyboard) are enough. For the undergraduate platform: comfort with functions and
graphs helps; nothing else is assumed for the entry-level activity.

## Six practices this repository is set up to support

1. **Connect it to their lives.** Start from memory they already have — a song lyric
   that unlocks a whole verse, a smell that brings back a room. "Your brain is a
   network. Prove it." is the opening move in the facilitator guide because every
   learner already has the evidence.
2. **Let them investigate.** Every module has a prediction point before a result is
   shown. Ask what will light up, which route will win, what happens when a distractor
   is added — then run it and make them reconcile the difference.
3. **Make struggle normal.** The Brain Power Meter exists to teach that difficulty is
   the *input* to learning, not a verdict on the learner. Say it out loud: "this is the
   part where the model gets interesting," not "this is the part where you're behind."
4. **Challenge the stereotype.** Name who does this work — the platform ships a Real
   Scientists view for exactly this reason — and praise the method ("you tested three
   routes before deciding") rather than the person ("you're a natural").
5. **Make it social.** Mixed-ability pairs, not ability groups: the modules reward
   explanation, not speed. Close each block by having two pairs read their sentence and
   the room decide which explanation is strongest.
6. **Show diverse role models.** The kids track names six working scientists whose
   actual contribution maps onto one of the four modules — Bassett, Marder,
   Levi-Montalcini, Milner, Perlman, Allen — each with a "try this" tie-in.

## A 90-minute session

This is Format B from [`docs/FACILITATOR-GUIDE.md`](docs/FACILITATOR-GUIDE.md): two
modules, 45 minutes each.

| Min | What | Notes |
|---|---|---|
| 0–5 | Hook: "Your brain is a network. Prove it." | Ask for two words they think are connected; write both up. |
| 5–10 | Whole-group demo of **Word Web** on one screen | Activate one concept; ask them to predict what lights up next. |
| 10–30 | Pairs on **Word Web** | Circulate. Ask "why did *that* one light up?", not "did you finish?". |
| 30–38 | Pairs write one sentence: *"Activation spreads to ___ because ___"* | This sentence is the assessment. Collect it. |
| 38–45 | Two pairs read out; the room decides which explanation is strongest and why | |
| 45–50 | Hook for block two: "The web was ideas. Now it's cells. What's the same?" | |
| 50–80 | Pairs on **Neuron Lab**, same circulate-and-predict pattern | Ask them to change one input and explain the change. |
| 80–88 | Closing sentence: *"Neurons fire when ___ because ___."* | |
| 88–90 | One-sentence exit ticket: name a thing you got wrong and what told you | |

For a four-session club sequence (Word Web → Neuron Lab → Thought Paths → Brain Power
Meter), use Format C in the facilitator guide. That order moves representation →
mechanism → algorithm → self-regulation, and it ends on metacognition, which is the
part that transfers.

## Vocabulary to define before the session

Define each of these the first time it appears, in one sentence, in the learners' own
words. Specialist vocabulary is a hidden entry test if nobody explains it.

- **Network** — a set of things, plus the connections between them.
- **Activation** — how "switched on" an idea currently is; it fades unless something
  keeps it going.
- **Spread** — activation passing along a connection to a neighbour.
- **Neuron** — a cell that carries a signal; it fires when its input crosses a
  threshold.
- **Spike** — one firing event. A spike train is a sequence of them over time.
- **Path** — an ordered list of connections that gets you from one node to another.
- **Shortest path** — the path with the lowest total cost, which is not always the one
  with the fewest steps.
- **Intrinsic load** — difficulty that comes from the material itself.
- **Extraneous load** — difficulty added by how the material is presented; this is the
  one you can remove.
- **Germane load** — effort that builds the mental structure you actually keep.

## Accessibility

Read [`ACCESSIBILITY.md`](ACCESSIBILITY.md) before running a session, and
[`docs/ACCESSIBILITY.md`](docs/ACCESSIBILITY.md) for the measured audit. Two rules
matter most in a mixed-ability room:

- Every activity must be completable with a keyboard alone, with a visible focus ring.
  Each canvas has an equivalent toolbar button for every action it accepts.
- Every instruction names the control or the heading ("the **Word Web** button in the
  toolbar"), never a colour or a position.

The platform speaks its instructions aloud for learners who need reading support; that
is a reading aid, not a replacement for a screen reader. Pair learners rather than
singling anyone out, and always offer a text route for a visual step.

## Contributing as a learner

Small, well-scoped contributions are the point, not a consolation prize:

- Browse issues labelled
  [`good first issue`](https://github.com/myzticalbabe/wasm-education/labels/good%20first%20issue).
- Read [CONTRIBUTING.md](CONTRIBUTING.md) before opening a pull request.
- The single most useful contribution is a
  [facilitator report](.github/ISSUE_TEMPLATE/facilitator_report.yml) from a real
  session — no detail is too small.
- Ask in [SUPPORT.md](SUPPORT.md) if anything is unclear. Asking is expected here; it
  is not a failure.

## Licence and classroom reuse

This project is licensed under the **SOUL FORGE License v1.0** — see [LICENSE](LICENSE)
for the exact terms. GitHub's detector reports `NOASSERTION` because the text is not a
standard SPDX template; `LICENSE` is authoritative. Classroom, club and workshop use is
welcome; check whether your institution requires a separate agreement before
redistributing modified copies.
