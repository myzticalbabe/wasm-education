# Standards Alignment — WASM Education Platform

Every module in this repository is mapped to named standards below. The codes are
the ones a teacher, curriculum lead, or after-school programme coordinator is
already accountable for, so a module can be dropped into an existing unit rather
than run as an extra.

Standards referenced:

- **NGSS** — Next Generation Science Standards (Achieve, 2013), including the
  Science and Engineering Practices (SEP) and Crosscutting Concepts (CCC).
- **CSTA** — CSTA K–12 Computer Science Standards, revised 2017.
- **ISTE** — ISTE Standards for Students, 2016.

---

## Module → standard map

### 1. Semantic Network Explorer ("Word Web") — semantics, memory, spreading activation

| Standard | Code | What the learner does |
|---|---|---|
| NGSS | **MS-LS1-8** | Gathers and synthesises information about how sensory input is stored as memory |
| NGSS | SEP 2 — Developing and Using Models | Builds a network model of memory and revises it as activation spreads |
| NGSS | CCC-2 — Cause and Effect | Predicts which concepts light up when a neighbour is activated |
| CSTA | **1B-DA-06** | Organises and presents the network visually to highlight relationships |
| CSTA | **1B-DA-07** | Uses the visualisation to propose cause-and-effect relationships between concepts |
| CSTA | **1B-AP-11** | Decomposes "why do I remember some words faster" into sub-problems (priming, load, attention) |
| ISTE | 1.5 Computational Thinker (1.5b) | Collects and represents data in a model to solve a problem |
| ISTE | 1.3 Knowledge Constructor (1.3d) | Builds knowledge by exploring how memory actually works |

### 2. Neural Dynamics Simulator ("Neuron Lab") — membrane dynamics, spiking

| Standard | Code | What the learner does |
|---|---|---|
| NGSS | **MS-LS1-8** | Connects a stimulus to a spike train and to a resulting behaviour |
| NGSS | **4-LS1-2** | Uses a model to describe how information is received, processed, and responded to |
| NGSS | **MS-LS1-3** | Treats the nervous system as one interacting subsystem of the body |
| NGSS | SEP 4 — Analysing and Interpreting Data | Reads a raster plot; separates signal (spikes) from noise |
| NGSS | SEP 5 — Using Mathematics and Computational Thinking | Changes `dt`, input current, and coupling, then explains the change |
| NGSS | CCC-1 — Patterns | Finds the regularities in firing across neuron types |
| NGSS | CCC-4 — Systems and System Models | Explains why network behaviour is not the sum of single neurons |
| CSTA | **2-AP-17** | Systematically tests and refines the model across a range of inputs |
| ISTE | 1.5 Computational Thinker (1.5d) | Understands how automation (the simulation loop) produces the behaviour |

### 3. Graph Algorithm Visualizer ("Thought Paths") — shortest path, SCC, MST

| Standard | Code | What the learner does |
|---|---|---|
| NGSS | **MS-ETS1-4** | Develops a model to generate data for iterative testing of a proposed process |
| NGSS | SEP 5 — Using Mathematics and Computational Thinking | Compares algorithms on the same graph and justifies a choice |
| CSTA | **1B-AP-08** | Compares and refines multiple algorithms for the same task and picks the most appropriate |
| CSTA | **1B-AP-15** | Tests and debugs a path when it turns out to be longer than expected |
| CSTA | **2-AP-10** | Uses flow/logic to describe how the search proceeds |
| CSTA | **2-AP-13** | Decomposes route-finding into frontier, visited set, and cost |
| CSTA | **1B-CS-02** | Models how hardware and software work together to accomplish the task |
| ISTE | 1.5 Computational Thinker (1.5a, 1.5c) | Formulates a problem definition and breaks it into parts |

### 4. Cognitive Load Simulator ("Brain Power Meter") — instructional design

| Standard | Code | What the learner does |
|---|---|---|
| NGSS | SEP 2 — Developing and Using Models | Models three loads (intrinsic, extraneous, germane) and predicts overflow |
| NGSS | CCC-5 — Energy and Matter | Tracks a limited resource being consumed by competing demands |
| CSTA | **1B-AP-12** | Remixes the scenario — adds a distractor, removes one — and observes the effect |
| CSTA | **2-AP-11** | Uses named quantities (load, capacity, pressure) with operations on their values |
| ISTE | 1.5 Computational Thinker (1.5b) | Represents data to highlight a relationship |
| ISTE | 1.4 Innovative Designer (1.4b) | Selects and uses a digital tool to model a real constraint |

---

## Cross-module coverage

Addressing all four modules covers, at minimum:

- **NGSS**: MS-LS1-8, MS-LS1-3, 4-LS1-2, MS-ETS1-4, SEP 2/4/5, CCC-1/2/4/5
- **CSTA 2017**: 1B-AP-08, 1B-AP-11, 1B-AP-12, 1B-AP-15, 1B-CS-02, 1B-DA-06,
  1B-DA-07, 2-AP-10, 2-AP-11, 2-AP-13, 2-AP-17
- **ISTE 2016**: 1.3d, 1.4b, 1.5a, 1.5b, 1.5c, 1.5d

## Grade-band guidance

| Band | Modules | Rationale |
|---|---|---|
| Grades 3–5 (ages 8–10) | Word Web, Brain Power Meter | Concrete manipulation, no algebra; `4-LS1-2` is the anchor |
| Grades 6–8 (ages 11–13) | All four | `MS-LS1-8` and `MS-ETS1-4` are the anchors; graph module needs proportional reasoning |
| Grades 9–12 | All four, plus `demo-pages/index.html` | Undergraduate page adds the Izhikevich model and Tarjan SCC explicitly |

## Verification note

Standards were mapped against the published text of each standard, not from memory
of the code names. If you are adopting this in a US state that has not adopted NGSS,
the SEP/CCC rows and the CSTA rows still transfer; only the `MS-LS1-*` and
`MS-ETS1-*` performance expectations are NGSS-specific.
