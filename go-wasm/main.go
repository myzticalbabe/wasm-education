// Package wasm_education provides educational WASM modules for cognitive/semantic learning
// Built with TinyGo for WebAssembly targeting browser environments
package main

import (
	"encoding/json"
	"math"
	"math/rand"
	"strconv"
	"syscall/js"
	"time"
)

// ─────────────────────────────────────────────────────────────────────────────
// CORE TYPES — Semantic Annotations
// ─────────────────────────────────────────────────────────────────────────────

// SemanticLabel holds cognitive metadata for concepts
type SemanticLabel struct {
	Text              string   `json:"text"`
	ConceptType       string   `json:"concept_type"`       // Definition, Theorem, Algorithm, Visualization, Interactive, Analogy, Proof, Application, MetaCognition
	Difficulty        string   `json:"difficulty"`         // Introductory, Intermediate, Advanced, Research
	Prerequisites     []string `json:"prerequisites"`
	LearningObjectives []string `json:"learning_objectives"`
	CognitiveLoad     float32  `json:"cognitive_load"`     // 0.0 - 1.0
	SemanticTags      []string `json:"semantic_tags"`
}

// CognitiveScaffold provides progressive disclosure
type CognitiveScaffold struct {
	Level             int            `json:"level"`
	RevealedConcepts  []string       `json:"revealed_concepts"`
	HiddenConcepts    []string       `json:"hidden_concepts"`
	Guidance          string         `json:"guidance"`
	CheckQuestions    []CheckQuestion `json:"check_questions"`
}

type CheckQuestion struct {
	Question       string   `json:"question"`
	AnswerType     string   `json:"answer_type"`     // MultipleChoice, FreeText, Numeric, Visual
	Options        []string `json:"options,omitempty"`
	CorrectAnswer  string   `json:"correct_answer"`
	Explanation    string   `json:"explanation"`
	ConceptTested  string   `json:"concept_tested"`
}

// ─────────────────────────────────────────────────────────────────────────────
// MODULE 1: SEMANTIC NETWORK — Spreading Activation
// ─────────────────────────────────────────────────────────────────────────────

type SemanticNetwork struct {
	nodes       map[string]*SemanticNode
	edges       map[string][]SemanticEdge
	activation  map[string]float32
	labels      map[string]SemanticLabel
	scaffold    CognitiveScaffold
	rng         *rand.Rand
}

type SemanticNode struct {
	ID       string   `json:"id"`
	Concept  string   `json:"concept"`
	NodeType string   `json:"node_type"` // Concept, Category, Instance, Relation, Property, Process
	Weight   float32  `json:"weight"`
	Tags     []string `json:"tags"`
}

type SemanticEdge struct {
	Target       string  `json:"target"`
	Relation     string  `json:"relation"`
	Strength     float32 `json:"strength"`
	Bidirectional bool    `json:"bidirectional"`
}

func NewSemanticNetwork() *SemanticNetwork {
	sn := &SemanticNetwork{
		nodes:      make(map[string]*SemanticNode),
		edges:      make(map[string][]SemanticEdge),
		activation: make(map[string]float32),
		labels:     make(map[string]SemanticLabel),
		rng:        rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	sn.initializeCoreNetwork()
	return sn
}

func (sn *SemanticNetwork) initializeCoreNetwork() {
	concepts := []struct {
		id, concept, nodeType string
		tags                  []string
	}{
		{"working_memory", "Working Memory", "Concept", []string{"cognitive_psychology", "neuroscience"}},
		{"long_term_memory", "Long-Term Memory", "Concept", []string{"cognitive_psychology", "neuroscience"}},
		{"attention", "Attention", "Process", []string{"cognitive_psychology", "neuroscience"}},
		{"spreading_activation", "Spreading Activation", "Process", []string{"semantic_memory", "connectionism"}},
		{"semantic_priming", "Semantic Priming", "Concept", []string{"semantic_memory", "psycholinguistics"}},
		{"schema", "Schema", "Category", []string{"cognitive_psychology", "knowledge_representation"}},
		{"mental_model", "Mental Model", "Concept", []string{"cognitive_science", "reasoning"}},
		{"cognitive_load", "Cognitive Load", "Concept", []string{"educational_psychology", "instructional_design"}},
		{"dual_coding", "Dual Coding Theory", "Theorem", []string{"memory", "multimedia_learning"}},
		{"elaboration", "Elaboration", "Process", []string{"memory", "learning_strategies"}},
		{"retrieval_practice", "Retrieval Practice", "Algorithm", []string{"learning", "metacognition"}},
		{"spaced_repetition", "Spaced Repetition", "Algorithm", []string{"memory", "optimization"}},
		{"interleaving", "Interleaving", "Algorithm", []string{"learning", "optimization"}},
		{"metacognition", "Metacognition", "MetaCognition", []string{"self_regulation", "learning"}},
		{"transfer", "Transfer of Learning", "Concept", []string{"educational_psychology", "application"}},
	}

	for _, c := range concepts {
		sn.nodes[c.id] = &SemanticNode{
			ID:       c.id,
			Concept:  c.concept,
			NodeType: c.nodeType,
			Weight:   1.0,
			Tags:     c.tags,
		}
		sn.labels[c.id] = SemanticLabel{
			Text:               c.concept,
			ConceptType:        sn.conceptTypeFromNodeType(c.nodeType),
			Difficulty:         "Intermediate",
			Prerequisites:      []string{},
			LearningObjectives: []string{"Understand " + c.concept},
			CognitiveLoad:      0.5,
			SemanticTags:       c.tags,
		}
		sn.activation[c.id] = 0.0
	}

	relations := []struct {
		src, tgt, relation string
		strength           float32
		bidir              bool
	}{
		{"working_memory", "long_term_memory", "consolidation", 0.8, true},
		{"working_memory", "attention", "gatekeeping", 0.9, true},
		{"attention", "spreading_activation", "modulation", 0.7, true},
		{"spreading_activation", "semantic_priming", "mechanism", 0.85, true},
		{"schema", "mental_model", "instantiation", 0.75, true},
		{"cognitive_load", "working_memory", "capacity_limit", 0.9, true},
		{"dual_coding", "working_memory", "parallel_channels", 0.8, true},
		{"elaboration", "long_term_memory", "encoding", 0.85, true},
		{"retrieval_practice", "long_term_memory", "strengthening", 0.9, true},
		{"spaced_repetition", "retrieval_practice", "scheduling", 0.8, true},
		{"interleaving", "retrieval_practice", "variation", 0.7, true},
		{"metacognition", "retrieval_practice", "monitoring", 0.75, true},
		{"transfer", "mental_model", "application", 0.8, true},
		{"schema", "transfer", "framework", 0.7, true},
	}

	for _, r := range relations {
		sn.edges[r.src] = append(sn.edges[r.src], SemanticEdge{
			Target:        r.tgt,
			Relation:      r.relation,
			Strength:      r.strength,
			Bidirectional: r.bidir,
		})
		if r.bidir {
			sn.edges[r.tgt] = append(sn.edges[r.tgt], SemanticEdge{
				Target:        r.src,
				Relation:      r.relation,
				Strength:      r.strength,
				Bidirectional: true,
			})
		}
	}

	sn.scaffold = CognitiveScaffold{
		Level:            1,
		RevealedConcepts: []string{"working_memory", "attention"},
		HiddenConcepts:   []string{"spreading_activation", "schema", "metacognition"},
		Guidance:         "Start by exploring Working Memory and Attention — the gatekeepers of cognition.",
		CheckQuestions: []CheckQuestion{
			{
				Question:      "What limits the capacity of working memory?",
				AnswerType:    "FreeText",
				CorrectAnswer: "Cognitive load and attentional resources",
				Explanation:   "Working memory has limited capacity (~4±1 chunks) governed by cognitive load theory.",
				ConceptTested: "working_memory",
			},
		},
	}
}

func (sn *SemanticNetwork) conceptTypeFromNodeType(nt string) string {
	switch nt {
	case "Theorem":
		return "Theorem"
	case "Process":
		return "Algorithm"
	case "Category":
		return "Definition"
	default:
		return "Definition"
	}
}

// Activate spreads activation through the network
func (sn *SemanticNetwork) Activate(this js.Value, args []js.Value) interface{} {
	if len(args) < 1 {
		return nil
	}
	conceptID := args[0].String()
	initialActivation := float32(1.0)
	if len(args) > 1 {
		initialActivation = float32(args[1].Float())
	}

	// Reset activation
	for k := range sn.activation {
		sn.activation[k] = 0.0
	}
	sn.activation[conceptID] = initialActivation

	// Spreading activation (3 iterations)
	for range 3 {
		newActivation := make(map[string]float32)
		for k, v := range sn.activation {
			newActivation[k] = v
		}
		for src, edges := range sn.edges {
			srcAct := sn.activation[src]
			for _, edge := range edges {
				spread := srcAct * edge.Strength * 0.3
				current := newActivation[edge.Target]
				newActivation[edge.Target] = float32(math.Min(float64(current+spread), 1.0))
				if edge.Bidirectional {
					backSpread := sn.activation[edge.Target] * edge.Strength * 0.3
					currentSrc := newActivation[src]
					newActivation[src] = float32(math.Min(float64(currentSrc+backSpread), 1.0))
				}
			}
		}
		sn.activation = newActivation
	}

	// Return activation state as JSON
	result, _ := json.Marshal(sn.activation)
	return js.ValueOf(string(result))
}

// GetScaffold returns current cognitive scaffold
func (sn *SemanticNetwork) GetScaffold(this js.Value, args []js.Value) interface{} {
	result, _ := json.Marshal(sn.scaffold)
	return js.ValueOf(string(result))
}

// AdvanceScaffold increases progressive disclosure level
func (sn *SemanticNetwork) AdvanceScaffold(this js.Value, args []js.Value) interface{} {
	sn.scaffold.Level++
	switch sn.scaffold.Level {
	case 2:
		sn.scaffold.RevealedConcepts = append(sn.scaffold.RevealedConcepts, "spreading_activation", "schema")
		newHidden := []string{}
		for _, c := range sn.scaffold.HiddenConcepts {
			if c != "spreading_activation" && c != "schema" {
				newHidden = append(newHidden, c)
			}
		}
		sn.scaffold.HiddenConcepts = newHidden
		sn.scaffold.Guidance = "Explore Spreading Activation — how concepts prime each other through semantic links."
		sn.scaffold.CheckQuestions = append(sn.scaffold.CheckQuestions, CheckQuestion{
			Question:      "How does semantic priming demonstrate spreading activation?",
			AnswerType:    "MultipleChoice",
			Options:       []string{"Faster recognition of related words", "Slower recognition of unrelated words", "No effect on recognition", "Only affects visual processing"},
			CorrectAnswer: "Faster recognition of related words",
			Explanation:   "Priming activates related concepts, reducing retrieval time for semantically linked items.",
			ConceptTested: "semantic_priming",
		})
	case 3:
		sn.scaffold.RevealedConcepts = append(sn.scaffold.RevealedConcepts, "metacognition")
		newHidden := []string{}
		for _, c := range sn.scaffold.HiddenConcepts {
			if c != "metacognition" {
				newHidden = append(newHidden, c)
			}
		}
		sn.scaffold.HiddenConcepts = newHidden
		sn.scaffold.Guidance = "Metacognition — monitoring your own learning. Try the retrieval practice questions."
		sn.scaffold.CheckQuestions = append(sn.scaffold.CheckQuestions, CheckQuestion{
			Question:      "What is the testing effect?",
			AnswerType:    "FreeText",
			CorrectAnswer: "Retrieval practice strengthens memory more than re-reading",
			Explanation:   "Active recall creates stronger memory traces than passive review (Roediger & Karpicke, 2006).",
			ConceptTested: "retrieval_practice",
		})
	}
	result, _ := json.Marshal(sn.scaffold)
	return js.ValueOf(string(result))
}

// GetGraphData returns network structure for visualization
func (sn *SemanticNetwork) GetGraphData(this js.Value, args []js.Value) interface{} {
	type NodeData struct {
		ID             string   `json:"id"`
		Label          string   `json:"label"`
		X              float32  `json:"x"`
		Y              float32  `json:"y"`
		NodeType       string   `json:"node_type"`
		Activation     float32  `json:"activation"`
		SemanticTags   []string `json:"semantic_tags"`
	}
	type EdgeData struct {
		Source   string  `json:"source"`
		Target   string  `json:"target"`
		Relation string  `json:"relation"`
		Strength float32 `json:"strength"`
	}
	type GraphData struct {
		Nodes []NodeData `json:"nodes"`
		Edges []EdgeData `json:"edges"`
	}

	nodes := make([]NodeData, 0, len(sn.nodes))
	rng := rand.New(rand.NewSource(42))
	for id, node := range sn.nodes {
		label := sn.labels[id]
		nodes = append(nodes, NodeData{
			ID:           node.ID,
			Label:        node.Concept,
			X:            rng.Float32()*600 - 300,
			Y:            rng.Float32()*600 - 300,
			NodeType:     node.NodeType,
			Activation:   sn.activation[id],
			SemanticTags: label.SemanticTags,
		})
	}

	edges := make([]EdgeData, 0)
	for src, edgeList := range sn.edges {
		for _, edge := range edgeList {
			edges = append(edges, EdgeData{
				Source:   src,
				Target:   edge.Target,
				Relation: edge.Relation,
				Strength: edge.Strength,
			})
		}
	}

	result, _ := json.Marshal(GraphData{Nodes: nodes, Edges: edges})
	return js.ValueOf(string(result))
}

// ─────────────────────────────────────────────────────────────────────────────
// MODULE 2: NEURAL DYNAMICS — Izhikevich Neurons
// ─────────────────────────────────────────────────────────────────────────────

type NeuronType string

const (
	RegularSpiking      NeuronType = "RegularSpiking"
	FastSpiking         NeuronType = "FastSpiking"
	LowThresholdSpiking NeuronType = "LowThresholdSpiking"
	IntrinsicallyBursting NeuronType = "IntrinsicallyBursting"
	Chattering          NeuronType = "Chattering"
	Resonator           NeuronType = "Resonator"
)

type SynapseType string

const (
	Excitatory  SynapseType = "Excitatory"
	Inhibitory  SynapseType = "Inhibitory"
	Modulatory  SynapseType = "Modulatory"
)

type NeuronState struct {
	V              float32    `json:"v"`
	U              float32    `json:"u"`
	A              float32    `json:"a"`
	B              float32    `json:"b"`
	C              float32    `json:"c"`
	D              float32    `json:"d"`
	NeuronType     NeuronType `json:"neuron_type"`
	FiringHistory  []float32  `json:"firing_history"`
}

type Connection struct {
	From       int         `json:"from"`
	To         int         `json:"to"`
	Weight     float32     `json:"weight"`
	Delay      float32     `json:"delay"`
	SynapseType SynapseType `json:"synapse_type"`
}

type NeuralDynamics struct {
	neurons      []NeuronState
	connections  []Connection
	time         float32
	dt           float32
	labels       map[int]SemanticLabel
	rng          *rand.Rand
}

func NewNeuralDynamics(numNeurons int) *NeuralDynamics {
	nd := &NeuralDynamics{
		neurons:     make([]NeuronState, numNeurons),
		connections: []Connection{},
		time:        0.0,
		dt:          0.1,
		labels:      make(map[int]SemanticLabel),
		rng:         rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	nd.initializeNetwork(numNeurons)
	return nd
}

func (nd *NeuralDynamics) initializeNetwork(numNeurons int) {
	types := []NeuronType{
		RegularSpiking, FastSpiking, LowThresholdSpiking,
		IntrinsicallyBursting, Chattering, Resonator,
	}

	for i := 0; i < numNeurons; i++ {
		ntype := types[i%len(types)]
		var a, b, c, d float32
		switch ntype {
		case RegularSpiking:
			a, b, c, d = 0.02, 0.2, -65.0, 8.0
		case FastSpiking:
			a, b, c, d = 0.1, 0.2, -65.0, 2.0
		case LowThresholdSpiking:
			a, b, c, d = 0.02, 0.25, -65.0, 2.0
		case IntrinsicallyBursting:
			a, b, c, d = 0.02, 0.2, -55.0, 4.0
		case Chattering:
			a, b, c, d = 0.02, 0.2, -50.0, 2.0
		case Resonator:
			a, b, c, d = 0.1, 0.26, -65.0, 2.0
		}
		nd.neurons[i] = NeuronState{
			V: -65.0, U: b * -65.0, A: a, B: b, C: c, D: d,
			NeuronType: ntype, FiringHistory: []float32{},
		}
		nd.labels[i] = SemanticLabel{
			Text:               string(ntype)+ " Neuron #" + strconv.Itoa(i),
			ConceptType:        "Visualization",
			Difficulty:         "Advanced",
			Prerequisites:      []string{"izhikevich_model", "membrane_potential"},
			LearningObjectives: []string{"Observe spiking dynamics", "Understand neuron types"},
			CognitiveLoad:      0.7,
			SemanticTags:       []string{"computational_neuroscience", "spiking_networks", "izhikevich"},
		}
	}

	// Sparse connectivity (10%)
	for i := 0; i < numNeurons; i++ {
		for j := 0; j < numNeurons; j++ {
			if i != j && nd.rng.Float32() < 0.1 {
				synType := Excitatory
				if nd.rng.Float32() < 0.2 {
					synType = Inhibitory
				}
				nd.connections = append(nd.connections, Connection{
					From: i, To: j,
					Weight: nd.rng.Float32()*4.5 + 0.5,
					Delay: nd.rng.Float32()*4.0 + 1.0,
					SynapseType: synType,
				})
			}
		}
	}
}

// Step runs one simulation step
func (nd *NeuralDynamics) Step(this js.Value, args []js.Value) interface{} {
	var currents []float32
	if len(args) > 0 {
		// Parse input currents from JS array
		arr := args[0]
		length := arr.Length()
		currents = make([]float32, length)
		for i := 0; i < length; i++ {
			currents[i] = float32(arr.Index(i).Float())
		}
	}
	if len(currents) < len(nd.neurons) {
		currents = append(currents, make([]float32, len(nd.neurons)-len(currents))...)
	}

	for i := range nd.neurons {
		n := &nd.neurons[i]
		iSyn := currents[i]

		// Izhikevich dynamics
		n.V += nd.dt * (0.04*n.V*n.V + 5.0*n.V + 140.0 - n.U + iSyn)
		n.U += nd.dt * (n.A * (n.B*n.V - n.U))

		// Spike detection
		if n.V >= 30.0 {
			n.FiringHistory = append(n.FiringHistory, nd.time)
			n.V = n.C
			n.U += n.D

			// Propagate to connected neurons
			for _, conn := range nd.connections {
				if conn.From == i {
					target := &nd.neurons[conn.To]
					w := conn.Weight
					if conn.SynapseType == Inhibitory {
						w = -w
					} else if conn.SynapseType == Modulatory {
						w *= 0.5
					}
					target.V += w * 0.1
				}
			}
		}

		// Bound history
		if len(n.FiringHistory) > 1000 {
			n.FiringHistory = n.FiringHistory[500:]
		}
	}

	nd.time += nd.dt

	// Collect spikes
	spikes := make([]int, 0)
	potentials := make([]float32, len(nd.neurons))
	for i, n := range nd.neurons {
		potentials[i] = n.V
		if len(n.FiringHistory) > 0 && n.FiringHistory[len(n.FiringHistory)-1] == nd.time-nd.dt {
			spikes = append(spikes, i)
		}
	}

	result := map[string]interface{}{
		"potentials": potentials,
		"spikes":     spikes,
		"time":       nd.time,
	}
	jsonResult, _ := json.Marshal(result)
	return js.ValueOf(string(jsonResult))
}

// GetLabels returns neuron labels
func (nd *NeuralDynamics) GetLabels(this js.Value, args []js.Value) interface{} {
	result, _ := json.Marshal(nd.labels)
	return js.ValueOf(string(result))
}

// SetDT sets simulation time step
func (nd *NeuralDynamics) SetDT(this js.Value, args []js.Value) interface{} {
	if len(args) > 0 {
		nd.dt = float32(math.Max(0.01, math.Min(1.0, args[0].Float())))
	}
	return nil
}

// ─────────────────────────────────────────────────────────────────────────────
// MODULE 3: GRAPH ALGORITHMS — Dijkstra, SCC, MST
// ─────────────────────────────────────────────────────────────────────────────

type GraphAlgorithmVisualizer struct {
	nodes     map[string]int
	edges     map[string][]GraphEdge
	adjList   map[int][]GraphEdge
	labels    map[int]SemanticLabel
	stepHistory []AlgorithmStep
	currentStep int
	rng       *rand.Rand
}

type GraphEdge struct {
	To     int     `json:"to"`
	Weight float32 `json:"weight"`
}

type AlgorithmStep struct {
	Description      string   `json:"description"`
	HighlightedNodes []int    `json:"highlighted_nodes"`
	HighlightedEdges [][2]int `json:"highlighted_edges"`
	DataSnapshot     string   `json:"data_snapshot"`
}

func NewGraphAlgorithmVisualizer() *GraphAlgorithmVisualizer {
	gav := &GraphAlgorithmVisualizer{
		nodes:       make(map[string]int),
		edges:       make(map[string][]GraphEdge),
		adjList:     make(map[int][]GraphEdge),
		labels:      make(map[int]SemanticLabel),
		stepHistory: []AlgorithmStep{},
		currentStep: 0,
		rng:         rand.New(rand.NewSource(123)),
	}
	gav.createSampleGraph()
	return gav
}

func (gav *GraphAlgorithmVisualizer) createSampleGraph() {
	concepts := []string{
		"Perception", "Attention", "Working Memory", "Long-Term Memory",
		"Language", "Reasoning", "Decision Making", "Problem Solving",
		"Metacognition", "Learning", "Transfer", "Expertise",
	}

	for i, c := range concepts {
		gav.nodes[c] = i
		gav.adjList[i] = []GraphEdge{}
		gav.labels[i] = SemanticLabel{
			Text:               c,
			ConceptType:        "Definition",
			Difficulty:         "Intermediate",
			Prerequisites:      []string{},
			LearningObjectives: []string{"Understand " + c + " in cognitive architecture"},
			CognitiveLoad:      0.4,
			SemanticTags:       []string{"cognitive_architecture", "cognitive_science"},
		}
	}

	edgeList := []struct{ src, tgt string; w float32 }{
		{"Perception", "Attention", 0.9}, {"Attention", "Working Memory", 0.95},
		{"Working Memory", "Long-Term Memory", 0.8}, {"Working Memory", "Language", 0.7},
		{"Working Memory", "Reasoning", 0.85}, {"Working Memory", "Decision Making", 0.75},
		{"Working Memory", "Problem Solving", 0.9}, {"Long-Term Memory", "Reasoning", 0.7},
		{"Long-Term Memory", "Problem Solving", 0.65}, {"Language", "Reasoning", 0.6},
		{"Reasoning", "Decision Making", 0.8}, {"Problem Solving", "Metacognition", 0.7},
		{"Metacognition", "Learning", 0.85}, {"Learning", "Transfer", 0.75},
		{"Transfer", "Expertise", 0.8}, {"Expertise", "Problem Solving", 0.6},
	}

	for _, e := range edgeList {
		s := gav.nodes[e.src]
		t := gav.nodes[e.tgt]
		gav.adjList[s] = append(gav.adjList[s], GraphEdge{To: t, Weight: e.w})
		gav.adjList[t] = append(gav.adjList[t], GraphEdge{To: s, Weight: e.w})
		gav.edges[e.src] = append(gav.edges[e.src], GraphEdge{To: t, Weight: e.w})
	}
}

// RunDijkstra executes Dijkstra step by step
func (gav *GraphAlgorithmVisualizer) RunDijkstra(this js.Value, args []js.Value) interface{} {
	if len(args) < 1 {
		return nil
	}
	sourceLabel := args[0].String()
	var targetLabel string
	if len(args) > 1 {
		targetLabel = args[1].String()
	}

	source, ok := gav.nodes[sourceLabel]
	if !ok {
		return nil
	}
	var target int = -1
	if targetLabel != "" {
		if t, ok := gav.nodes[targetLabel]; ok {
			target = t
		}
	}

	gav.stepHistory = []AlgorithmStep{}
	gav.currentStep = 0

	n := len(gav.nodes)
	dist := make([]float32, n)
	visited := make([]bool, n)
	for i := range dist {
		dist[i] = float32(math.Inf(1))
	}
	dist[source] = 0.0

	type pqItem struct {
		dist float32
		node int
	}
	pq := []pqItem{{0.0, source}}

	gav.stepHistory = append(gav.stepHistory, AlgorithmStep{
		Description:      "Initialize: Set distance to source '" + sourceLabel + "' = 0, all others = ∞",
		HighlightedNodes: []int{source},
		HighlightedEdges: [][2]int{},
		DataSnapshot:     gav.snapshotDistances(dist, visited),
	})

	for len(pq) > 0 {
		// Pop min
		minIdx := 0
		for i := 1; i < len(pq); i++ {
			if pq[i].dist < pq[minIdx].dist {
				minIdx = i
			}
		}
		current := pq[minIdx]
		pq = append(pq[:minIdx], pq[minIdx+1:]...)

		if visited[current.node] {
			continue
		}
		visited[current.node] = true

		// Find label for current node
		currentLabel := ""
		for k, v := range gav.nodes {
			if v == current.node {
				currentLabel = k
				break
			}
		}

		gav.stepHistory = append(gav.stepHistory, AlgorithmStep{
			Description:      "Visit '" + currentLabel + "' (distance = " + strconv.FormatFloat(float64(current.dist), 'f', 2, 32) + ")",
			HighlightedNodes: []int{current.node},
			HighlightedEdges: [][2]int{},
			DataSnapshot:     gav.snapshotDistances(dist, visited),
		})

		if current.node == target {
			break
		}

		for _, edge := range gav.adjList[current.node] {
			if visited[edge.To] {
				continue
			}
			newDist := current.dist + edge.Weight
			if newDist < dist[edge.To] {
				dist[edge.To] = newDist
				pq = append(pq, pqItem{newDist, edge.To})
				
				// Find labels
				srcLabel, tgtLabel := "", ""
				for k, v := range gav.nodes {
					if v == current.node { srcLabel = k }
					if v == edge.To { tgtLabel = k }
				}

				gav.stepHistory = append(gav.stepHistory, AlgorithmStep{
					Description: "Relax edge " + srcLabel + " → " + tgtLabel + ": new distance = " + strconv.FormatFloat(float64(newDist), 'f', 2, 32),
					HighlightedNodes: []int{current.node, edge.To},
					HighlightedEdges: [][2]int{{current.node, edge.To}},
					DataSnapshot: gav.snapshotDistances(dist, visited),
				})
			}
		}
	}

	result, _ := json.Marshal(gav.stepHistory)
	return js.ValueOf(string(result))
}

func (gav *GraphAlgorithmVisualizer) snapshotDistances(dist []float32, visited []bool) string {
	type snap struct {
		Distances []float32 `json:"distances"`
		Visited   []bool    `json:"visited"`
	}
	s := snap{Distances: dist, Visited: visited}
	b, _ := json.Marshal(s)
	return string(b)
}

// NextStep returns next algorithm step
func (gav *GraphAlgorithmVisualizer) NextStep(this js.Value, args []js.Value) interface{} {
	if gav.currentStep >= len(gav.stepHistory) {
		return js.Null()
	}
	step := gav.stepHistory[gav.currentStep]
	gav.currentStep++
	result, _ := json.Marshal(step)
	return js.ValueOf(string(result))
}

// Reset resets to beginning
func (gav *GraphAlgorithmVisualizer) Reset(this js.Value, args []js.Value) interface{} {
	gav.currentStep = 0
	return nil
}

// GetGraphData returns graph for visualization
func (gav *GraphAlgorithmVisualizer) GetGraphData(this js.Value, args []js.Value) interface{} {
	type NodeData struct {
		ID    string  `json:"id"`
		Label string  `json:"label"`
		X     float32 `json:"x"`
		Y     float32 `json:"y"`
	}
	type EdgeData struct {
		Source  string  `json:"source"`
		Target  string  `json:"target"`
		Weight  float32 `json:"weight"`
	}
	type GraphData struct {
		Nodes []NodeData `json:"nodes"`
		Edges []EdgeData `json:"edges"`
	}

	nodes := make([]NodeData, 0, len(gav.nodes))
	for label, idx := range gav.nodes {
		nodes = append(nodes, NodeData{
			ID:    strconv.Itoa(idx),
			Label: label,
			X:     gav.rng.Float32()*800 - 400,
			Y:     gav.rng.Float32()*600 - 300,
		})
	}

	edges := make([]EdgeData, 0)
	for srcLabel, edgeList := range gav.edges {
		for _, edge := range edgeList {
			tgtLabel := ""
			for k, v := range gav.nodes {
				if v == edge.To {
					tgtLabel = k
					break
				}
			}
			edges = append(edges, EdgeData{
				Source: srcLabel, Target: tgtLabel, Weight: edge.Weight,
			})
		}
	}

	result, _ := json.Marshal(GraphData{Nodes: nodes, Edges: edges})
	return js.ValueOf(string(result))
}

// ─────────────────────────────────────────────────────────────────────────────
// MODULE 4: COGNITIVE LOAD SIMULATOR
// ─────────────────────────────────────────────────────────────────────────────

type CognitiveLoadSimulator struct {
	workingMemoryCapacity int
	currentLoad           float32
	intrinsicLoad         float32
	extraneousLoad        float32
	germaneLoad           float32
	schemaCount           int
	automationLevel       float32
	timePressure          float32
	labels                map[string]SemanticLabel
}

func NewCognitiveLoadSimulator() *CognitiveLoadSimulator {
	cls := &CognitiveLoadSimulator{
		workingMemoryCapacity: 4,
		labels:                make(map[string]SemanticLabel),
	}
	cls.initializeLabels()
	return cls
}

func (cls *CognitiveLoadSimulator) initializeLabels() {
	cls.labels["intrinsic_load"] = SemanticLabel{
		Text: "Intrinsic Cognitive Load", ConceptType: "Definition", Difficulty: "Intermediate",
		Prerequisites: []string{"working_memory"}, LearningObjectives: []string{"Understand element interactivity"},
		CognitiveLoad: 0.3, SemanticTags: []string{"cognitive_load_theory", "instructional_design"},
	}
	cls.labels["extraneous_load"] = SemanticLabel{
		Text: "Extraneous Cognitive Load", ConceptType: "Definition", Difficulty: "Intermediate",
		Prerequisites: []string{}, LearningObjectives: []string{"Identify poor instructional design"},
		CognitiveLoad: 0.3, SemanticTags: []string{"cognitive_load_theory", "ux_design"},
	}
	cls.labels["germane_load"] = SemanticLabel{
		Text: "Germane Cognitive Load", ConceptType: "Definition", Difficulty: "Advanced",
		Prerequisites: []string{"schema_acquisition"}, LearningObjectives: []string{"Design for schema construction"},
		CognitiveLoad: 0.4, SemanticTags: []string{"cognitive_load_theory", "learning_optimization"},
	}
	cls.labels["schema"] = SemanticLabel{
		Text: "Schema", ConceptType: "Definition", Difficulty: "Introductory",
		Prerequisites: []string{}, LearningObjectives: []string{"Understand chunking and automation"},
		CognitiveLoad: 0.2, SemanticTags: []string{"cognitive_psychology", "expertise"},
	}
}

// SimulateTask runs a cognitive load simulation
func (cls *CognitiveLoadSimulator) SimulateTask(this js.Value, args []js.Value) interface{} {
	if len(args) < 4 {
		return nil
	}
	elementCount := int(args[0].Float())
	elementInteractivity := float32(args[1].Float())
	instructionalQuality := float32(args[2].Float())
	learnerExpertise := float32(args[3].Float())

	cls.intrinsicLoad = float32(elementCount) * elementInteractivity / (1.0 + learnerExpertise*3.0)
	cls.extraneousLoad = (1.0 - instructionalQuality) * 2.0
	totalCapacity := float32(cls.workingMemoryCapacity)
	cls.germaneLoad = float32(math.Max(float64(totalCapacity-cls.intrinsicLoad-cls.extraneousLoad), 0))

	if cls.timePressure > 0.5 {
		mult := 1.0 + cls.timePressure
		cls.intrinsicLoad *= mult
		cls.extraneousLoad *= mult
	}

	cls.currentLoad = cls.intrinsicLoad + cls.extraneousLoad + cls.germaneLoad

	if cls.germaneLoad > 1.0 {
		cls.schemaCount++
		cls.automationLevel = float32(math.Min(float64(cls.schemaCount)*0.1, 0.8))
	}

	type LoadResult struct {
		Intrinsic      float32  `json:"intrinsic"`
		Extraneous     float32  `json:"extraneous"`
		Germane        float32  `json:"germane"`
		Total          float32  `json:"total"`
		Capacity       float32  `json:"capacity"`
		Overloaded     bool     `json:"overloaded"`
		SchemaCount    int      `json:"schema_count"`
		Automation     float32  `json:"automation"`
		Recommendations []string `json:"recommendations"`
	}

	recs := []string{}
	if cls.intrinsicLoad > totalCapacity*0.6 {
		recs = append(recs, "Reduce element interactivity or provide pre-training")
	}
	if cls.extraneousLoad > totalCapacity*0.3 {
		recs = append(recs, "Improve instructional design: remove split-attention, redundancy")
	}
	if cls.germaneLoad < 0.5 {
		recs = append(recs, "Increase germane load: add worked examples, self-explanation prompts")
	}
	if cls.currentLoad > totalCapacity {
		recs = append(recs, "⚠️ COGNITIVE OVERLOAD: Reduce complexity immediately")
	}

	result := LoadResult{
		Intrinsic: cls.intrinsicLoad, Extraneous: cls.extraneousLoad, Germane: cls.germaneLoad,
		Total: cls.currentLoad, Capacity: totalCapacity, Overloaded: cls.currentLoad > totalCapacity,
		SchemaCount: cls.schemaCount, Automation: cls.automationLevel, Recommendations: recs,
	}
	jsonResult, _ := json.Marshal(result)
	return js.ValueOf(string(jsonResult))
}

// SetTimePressure sets time pressure (0.0 - 1.0)
func (cls *CognitiveLoadSimulator) SetTimePressure(this js.Value, args []js.Value) interface{} {
	if len(args) > 0 {
		cls.timePressure = float32(math.Max(0.0, math.Min(1.0, args[0].Float())))
	}
	return nil
}

// GetLabels returns cognitive load labels
func (cls *CognitiveLoadSimulator) GetLabels(this js.Value, args []js.Value) interface{} {
	result, _ := json.Marshal(cls.labels)
	return js.ValueOf(string(result))
}

// ─────────────────────────────────────────────────────────────────────────────
// REGISTRY & EXPORTS
// ─────────────────────────────────────────────────────────────────────────────

func getModuleRegistry() interface{} {
	type ModuleInfo struct {
		Name              string   `json:"name"`
		Description       string   `json:"description"`
		CognitiveDomain   string   `json:"cognitive_domain"`
		Difficulty        string   `json:"difficulty"`
		KeyConcepts       []string `json:"key_concepts"`
		LearningObjectives []string `json:"learning_objectives"`
	}

	modules := []ModuleInfo{
		{
			Name: "SemanticNetwork", Description: "Interactive semantic network with spreading activation",
			CognitiveDomain: "Semantic Memory / Cognitive Psychology", Difficulty: "Intermediate",
			KeyConcepts: []string{"Spreading Activation", "Semantic Priming", "Working Memory", "Schema Theory", "Cognitive Scaffolding"},
			LearningObjectives: []string{"Understand activation spread", "Experience progressive disclosure", "Learn retrieval practice"},
		},
		{
			Name: "NeuralDynamics", Description: "Izhikevich spiking neuron simulator with 6 biologically-plausible neuron types",
			CognitiveDomain: "Computational Neuroscience", Difficulty: "Advanced",
			KeyConcepts: []string{"Action Potentials", "Neuron Types", "Synaptic Transmission", "Network Dynamics", "E/I Balance"},
			LearningObjectives: []string{"Observe spiking patterns", "Understand E/I balance", "Explore network dynamics"},
		},
		{
			Name: "GraphAlgorithmVisualizer", Description: "Step-by-step graph algorithms on cognitive architecture graphs",
			CognitiveDomain: "Graph Theory / Cognitive Architecture", Difficulty: "Intermediate",
			KeyConcepts: []string{"Shortest Path", "Strongly Connected Components", "MST", "Cognitive Process Flow"},
			LearningObjectives: []string{"Trace algorithm execution", "Apply graphs to cognitive architectures"},
		},
		{
			Name: "CognitiveLoadSimulator", Description: "Real-time cognitive load simulation based on CLT",
			CognitiveDomain: "Educational Psychology / Instructional Design", Difficulty: "Intermediate",
			KeyConcepts: []string{"Intrinsic Load", "Extraneous Load", "Germane Load", "Schema Acquisition", "Automation"},
			LearningObjectives: []string{"Calculate cognitive load", "Optimize germane load", "Reduce extraneous load"},
		},
	}
	result, _ := json.Marshal(modules)
	return js.ValueOf(string(result))
}

// ─────────────────────────────────────────────────────────────────────────────
// WASM ENTRY POINT
// ─────────────────────────────────────────────────────────────────────────────

func registerCallbacks() {
	js.Global().Set("wasmEducation", js.ValueOf(map[string]interface{}{
		"newSemanticNetwork": func() interface{} { return NewSemanticNetwork() },
		"newNeuralDynamics": func(args []js.Value) interface{} {
			n := 50
			if len(args) > 0 { n = int(args[0].Float()) }
			return NewNeuralDynamics(n)
		},
		"newGraphAlgorithmVisualizer": func() interface{} { return NewGraphAlgorithmVisualizer() },
		"newCognitiveLoadSimulator": func() interface{} { return NewCognitiveLoadSimulator() },
		"getModuleRegistry": getModuleRegistry,
	}))

	// Register methods on SemanticNetwork
	sn := NewSemanticNetwork()
	js.Global().Get("wasmEducation").Set("SemanticNetwork", js.ValueOf(map[string]interface{}{
		"activate":       js.FuncOf(sn.Activate),
		"getScaffold":    js.FuncOf(sn.GetScaffold),
		"advanceScaffold": js.FuncOf(sn.AdvanceScaffold),
		"getGraphData":   js.FuncOf(sn.GetGraphData),
	}))

	// NeuralDynamics
	nd := NewNeuralDynamics(50)
	js.Global().Get("wasmEducation").Set("NeuralDynamics", js.ValueOf(map[string]interface{}{
		"step":      js.FuncOf(nd.Step),
		"getLabels": js.FuncOf(nd.GetLabels),
		"setDT":     js.FuncOf(nd.SetDT),
	}))

	// GraphAlgorithmVisualizer
	gav := NewGraphAlgorithmVisualizer()
	js.Global().Get("wasmEducation").Set("GraphAlgorithmVisualizer", js.ValueOf(map[string]interface{}{
		"runDijkstra": js.FuncOf(gav.RunDijkstra),
		"nextStep":    js.FuncOf(gav.NextStep),
		"reset":       js.FuncOf(gav.Reset),
		"getGraphData": js.FuncOf(gav.GetGraphData),
	}))

	// CognitiveLoadSimulator
	cls := NewCognitiveLoadSimulator()
	js.Global().Get("wasmEducation").Set("CognitiveLoadSimulator", js.ValueOf(map[string]interface{}{
		"simulateTask":  js.FuncOf(cls.SimulateTask),
		"setTimePressure": js.FuncOf(cls.SetTimePressure),
		"getLabels":    js.FuncOf(cls.GetLabels),
	}))
}

func main() {
	// Register the JS API before blocking forever. Without this call the
	// module builds but exports nothing.
	registerCallbacks()

	// TinyGo WASM needs main to stay alive; the API is driven by the
	// registered callbacks above.
	select {}
}