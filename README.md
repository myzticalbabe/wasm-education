# WASM Education Platform

**Cognitive & Semantic Learning Platform** — WebAssembly-powered educational modules for cognitive science, computational neuroscience, graph theory, and educational psychology.

[![CI](https://github.com/myzticalbabe/wasm-education/actions/workflows/ci.yml/badge.svg)](https://github.com/myzticalbabe/wasm-education/actions/workflows/ci.yml)
[![Pages Deploy](https://github.com/myzticalbabe/wasm-education/actions/workflows/pages-deploy.yml/badge.svg)](https://github.com/myzticalbabe/wasm-education/actions/workflows/pages-deploy.yml)
[![License: SOUL FORGE](https://img.shields.io/badge/License-SOUL_FORGE-violet.svg)](LICENSE)
[![GitHub release](https://img.shields.io/github/v/release/myzticalbabe/wasm-education)](https://github.com/myzticalbabe/wasm-education/releases)
[![GitHub stars](https://img.shields.io/github/stars/myzticalbabe/wasm-education?style=social)](https://github.com/myzticalbabe/wasm-education/stargazers)

---

## Overview

This platform provides **WebAssembly-powered educational modules** for higher education and advanced developer learning. Each module is designed with:

- **Semantic annotations** — every concept carries cognitive metadata (difficulty, prerequisites, learning objectives, cognitive load)
- **Cognitive scaffolding** — progressive disclosure through guided levels with check questions
- **Interactive visualization** — Canvas/WebGL renderings with real-time parameter control
- **WASM performance** — Rust (wasm-bindgen) and Go (TinyGo) compiled to WebAssembly for near-native speed
- **Educational rigor** — grounded in peer-reviewed theories (CLT, spreading activation, Izhikevich models, graph algorithms)

## Live Demo

🌐 **Interactive Platform**: https://myzticalbabe.github.io/wasm-education/

## Modules

| Module | Domain | Difficulty | Key Concepts |
|--------|--------|------------|--------------|
| **Semantic Network Explorer** | Semantic Memory / Cognitive Psychology | Intermediate | Spreading Activation, Semantic Priming, Working Memory, Schema Theory, Cognitive Scaffolding |
| **Neural Dynamics Simulator** | Computational Neuroscience | Advanced | Action Potentials, Neuron Types, Synaptic Transmission, Network Dynamics, E/I Balance |
| **Graph Algorithm Visualizer** | Graph Theory / Cognitive Architecture | Intermediate | Shortest Path, Strongly Connected Components, MST, Cognitive Process Flow |
| **Cognitive Load Simulator** | Educational Psychology / Instructional Design | Intermediate | Intrinsic Load, Extraneous Load, Germane Load, Schema Acquisition, Automation |

## Architecture

```
wasm-education/
├── rust-wasm/          # Rust → WASM (wasm-bindgen, serde, nalgebra, petgraph)
├── go-wasm/            # Go → WASM (TinyGo, wazero, syscall/js)
├── demo-pages/         # Vanilla JS interactive demo platform
├── docs/               # Documentation
└── .github/workflows/  # CI/CD pipelines
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
# Open http://localhost:8080
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

## Theoretical Foundations

| Module | Theory | Key References |
|--------|--------|----------------|
| Semantic Network | Spreading Activation, Schema Theory | Collins & Loftus (1975), Anderson (1983) |
| Neural Dynamics | Izhikevich Neuron Model | Izhikevich (2003, 2007) |
| Graph Algorithms | Dijkstra, Tarjan SCC, Kruskal | Cormen et al. (2009) |
| Cognitive Load | Cognitive Load Theory | Sweller (1988), Sweller et al. (2011) |

## License

This project is licensed under the **SOUL FORGE License v1.0** — see [LICENSE](LICENSE) for details.

> *"Code is crystallized intention. Tend the fire."*

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## Links

- **GitHub**: https://github.com/myzticalbabe/wasm-education
- **Live Demo**: https://myzticalbabe.github.io/wasm-education/
- **Issues**: https://github.com/myzticalbabe/wasm-education/issues
- **Discussions**: https://github.com/myzticalbabe/wasm-education/discussions
- **Author**: @myzticalbabe