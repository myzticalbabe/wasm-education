# WASM Education Platform

**Cognitive & Semantic Learning Platform** — WebAssembly-powered educational modules for
cognitive science, computational neuroscience, graph theory, and educational psychology.

Two tracks ship in this repository: an **undergraduate / developer platform** and a
**kids track for ages 8–13** with a facilitation kit for teachers, club leaders and
parents.

[![CI](https://github.com/myzticalbabe/wasm-education/actions/workflows/ci.yml/badge.svg)](https://github.com/myzticalbabe/wasm-education/actions/workflows/ci.yml)
[![Pages Deploy](https://github.com/myzticalbabe/wasm-education/actions/workflows/pages-deploy.yml/badge.svg)](https://github.com/myzticalbabe/wasm-education/actions/workflows/pages-deploy.yml)
[![License: SOUL FORGE](https://img.shields.io/badge/License-SOUL_FORGE-violet.svg)](LICENSE)
[![GitHub release](https://img.shields.io/github/v/release/myzticalbabe/wasm-education)](https://github.com/myzticalbabe/wasm-education/releases)

---

## Who this is for

| Audience | Entry point | What they get |
|---|---|---|
| Children, ages 8–13 | [`demo-pages/kids-index.html`](demo-pages/kids-index.html) | Four game-shaped adventures, read-aloud, badges, no accounts |
| Educators, librarians, club and troop leaders, parents | [Facilitator guide](docs/FACILITATOR-GUIDE.md) | Timed session plans, differentiation, assessment rubric |
| Undergraduates and developers | [`demo-pages/index.html`](demo-pages/index.html) | Izhikevich dynamics, Tarjan SCC, spreading activation, semantic annotations |
| Curriculum reviewers | [Standards alignment](docs/STANDARDS-ALIGNMENT.md) | NGSS / CSTA / ISTE codes per module |

## STEM for girls

Girls do not leave STEM because their ability drops; the measured drop is in
**belonging**, and it is concentrated in ages 11–14. Two avoidable contributors are
absent representation and ability-framed feedback. This repository addresses both
explicitly rather than by being gender-neutral, because neutral-by-omission leaves
the default signals in place.

Concretely, the kids track includes:

- a **Real Scientists** view naming six working scientists whose actual contribution
  maps onto one of the four modules (Bassett, Marder, Levi-Montalcini, Milner,
  Perlman, Allen), each with a "try this" tie-in;
- **growth framing** in copy — brains build connections by using them, and the
  Brain Power Meter teaches that difficulty is the *input* to learning;
- **no ranking surface** — badges mark exploration milestones, never comparison;
- **no identity disclosure** — no accounts, no name field, no demographic prompts.

The reasoning, the evidence, and the facilitation moves that make or break it are in
[`docs/INCLUSION-GIRLS-IN-STEM.md`](docs/INCLUSION-GIRLS-IN-STEM.md).

> This project does not claim to close the gender gap. It claims something narrower
> and checkable: it removes two avoidable contributors from a resource that
> otherwise has none.

## Live Demo

🌐 **Kids track**: https://myzticalbabe.github.io/wasm-education/kids-index.html
🌐 **Platform**: https://myzticalbabe.github.io/wasm-education/

## Modules

### Kids track (ages 8–13)

| Adventure | Domain | Standard |
|---|---|---|
| **Word Web** | Semantic memory, spreading activation | NGSS MS-LS1-8 |
| **Neuron Lab** | Neuron types, spiking, network dynamics | NGSS MS-LS1-8, MS-LS1-3 |
| **Thought Paths** | Shortest path, spanning trees | CSTA 1B-AP-08, NGSS MS-ETS1-4 |
| **Brain Power Meter** | Cognitive load: intrinsic, extraneous, germane | ISTE 1.5 |

### Platform (undergraduate / developer)

| Module | Domain | Difficulty | Key Concepts |
|--------|--------|------------|--------------|
| **Semantic Network Explorer** | Semantic Memory / Cognitive Psychology | Intermediate | Spreading Activation, Semantic Priming, Working Memory, Schema Theory, Cognitive Scaffolding |
| **Neural Dynamics Simulator** | Computational Neuroscience | Advanced | Action Potentials, Neuron Types, Synaptic Transmission, Network Dynamics, E/I Balance |
| **Graph Algorithm Visualizer** | Graph Theory / Cognitive Architecture | Intermediate | Shortest Path, Strongly Connected Components, MST, Cognitive Process Flow |
| **Cognitive Load Simulator** | Educational Psychology / Instructional Design | Intermediate | Intrinsic Load, Extraneous Load, Germane Load, Schema Acquisition, Automation |

## Documentation

| Document | For |
|---|---|
| [Facilitator guide](docs/FACILITATOR-GUIDE.md) | Running a 45-min, 90-min or four-session programme |
| [Assessment rubric](docs/ASSESSMENT-RUBRIC.md) | Judging whether a learner can *explain* what they saw |
| [Standards alignment](docs/STANDARDS-ALIGNMENT.md) | NGSS, CSTA and ISTE codes per module |
| [Inclusion: girls in STEM](docs/INCLUSION-GIRLS-IN-STEM.md) | Design rationale and the evidence behind it |
| [Accessibility](docs/ACCESSIBILITY.md) | What is implemented, and what is still unverified |
| [Contributing](CONTRIBUTING.md) | Non-negotiables before you open a PR |
| [Security policy](SECURITY.md) | Threat model and private reporting |
| [Code of conduct](CODE_OF_CONDUCT.md) | Expected behaviour in issues and sessions |

## Architecture

```
wasm-education/
├── rust-wasm/          # Rust → WASM (wasm-bindgen, serde, nalgebra, petgraph)
├── go-wasm/            # Go → WASM (TinyGo, wazero, syscall/js)
├── demo-pages/         # Vanilla JS demo platform + kids track
│   ├── index.html      # undergraduate / developer platform
│   ├── kids-index.html # kids track, ages 8–13
│   ├── manifest.json   # PWA manifest
│   └── icons/          # app icons (192, 512)
├── docs/               # standards, facilitation, assessment, inclusion, a11y
└── .github/            # CI/CD pipelines and issue templates
```

## Installation & Development

### Prerequisites
- Rust 1.75+ with `wasm32-unknown-unknown` target
- TinyGo 0.31+
- Node.js 20+ (for demo pages)
- wasm-pack (`cargo install wasm-pack`)

### Build Rust WASM
```bash
cd rust-wasm
wasm-pack build --target web --out-dir pkg
```

### Build Go WASM
```bash
cd go-wasm
tinygo build -o main.wasm -target wasm -no-debug -opt=s .
```

### Run Demo Pages
```bash
cd demo-pages
npx http-server .
# Open http://localhost:8080/kids-index.html
```

## Usage

Each module exports a JavaScript API:

```javascript
// Semantic Network
const sn = wasmEducation.newSemanticNetwork();
sn.activate('working_memory', 1.0);
sn.advanceScaffold();
const graphData = sn.getGraphData();

// Neural Dynamics
const nd = wasmEducation.newNeuralDynamics(50);
nd.setDT(0.1);
const result = nd.step(inputCurrents);

// Graph Algorithms
const gav = wasmEducation.newGraphAlgorithmVisualizer();
gav.runDijkstra('Working Memory', 'Metacognition');
gav.nextStep();

// Cognitive Load
const cls = wasmEducation.newCognitiveLoadSimulator();
cls.simulateTask(5, 0.7, 0.6, 0.3);
cls.setTimePressure(0.2);
```

## Learning Objectives

Each module includes:
- **Cognitive scaffolding** — progressive disclosure levels with guided questions
- **Semantic labels** — concept type, difficulty, prerequisites, learning objectives, cognitive load
- **Interactive check questions** — multiple choice, free text, visual answers with explanations
- **Real-time visualization** — Canvas/WebGL with parameter controls

## Privacy

The kids track collects nothing. No accounts, no forms, no names. One `localStorage`
key (`brainExplorerProgress`) holds XP and badge state, client-side only, and the page
makes no network requests after load. There is no backend and there are no secrets in
this repository. See [`SECURITY.md`](SECURITY.md) for the full threat model.

## Theoretical Foundations

| Module | Theory | Key References |
|--------|--------|----------------|
| Semantic Network | Spreading Activation, Schema Theory | Collins & Loftus (1975), Anderson (1983) |
| Neural Dynamics | Izhikevich Neuron Model | Izhikevich (2003, 2007) |
| Graph Algorithms | Dijkstra, Tarjan SCC, Kruskal | Cormen et al. (2009) |
| Cognitive Load | Cognitive Load Theory | Sweller (1988), Sweller et al. (2011) |
| Inclusion design | Role-model and stereotype-threat effects | Stout et al. (2011), Cheryan et al. (2011), Mueller & Dweck (1998) |

## License

This project is licensed under the **SOUL FORGE License v1.0** — see [LICENSE](LICENSE) for details.

> *"Code is crystallized intention. Tend the fire."*

GitHub's license detector only recognises OSI/SPDX-standard texts, so the repository
sidebar reports `NOASSERTION` for this project rather than a named license. The
[LICENSE](LICENSE) file is authoritative and the badge above is the accurate one.

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md). The most valuable contribution is a
[facilitator report](.github/ISSUE_TEMPLATE/facilitator_report.yml) from a real
session — no detail is too small.

## Links

- **GitHub**: https://github.com/myzticalbabe/wasm-education
- **Live Demo**: https://myzticalbabe.github.io/wasm-education/
- **Issues**: https://github.com/myzticalbabe/wasm-education/issues
- **Discussions**: https://github.com/myzticalbabe/wasm-education/discussions
- **Author**: @myzticalbabe
