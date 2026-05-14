// Package compose provides the core graph composition primitives for eino.
// It allows building complex AI pipelines by connecting nodes in a directed graph.
package compose

import (
	"context"
	"fmt"
	"sync"
)

// NodeType represents the type of a node in the graph.
type NodeType string

const (
	// NodeTypeLLM represents a large language model node.
	NodeTypeLLM NodeType = "llm"
	// NodeTypeTool represents a tool/function call node.
	NodeTypeTool NodeType = "tool"
	// NodeTypeRetriever represents a retriever node for RAG pipelines.
	NodeTypeRetriever NodeType = "retriever"
	// NodeTypeTransform represents a data transformation node.
	NodeTypeTransform NodeType = "transform"
)

// Node represents a single processing unit in the graph.
type Node struct {
	ID       string
	Type     NodeType
	runnable Runnable
}

// Runnable is the interface that all graph nodes must implement.
type Runnable interface {
	Run(ctx context.Context, input any) (any, error)
}

// Edge represents a directed connection between two nodes.
type Edge struct {
	From string
	To   string
}

// Graph is a directed acyclic graph of processing nodes.
type Graph struct {
	mu    sync.RWMutex
	nodes map[string]*Node
	edges []*Edge
	// adjacency list: nodeID -> list of successor nodeIDs
	adj map[string][]string
}

// NewGraph creates a new empty Graph.
func NewGraph() *Graph {
	return &Graph{
		nodes: make(map[string]*Node),
		adj:   make(map[string][]string),
	}
}

// AddNode registers a node with the given id and runnable into the graph.
func (g *Graph) AddNode(id string, nodeType NodeType, r Runnable) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	if _, exists := g.nodes[id]; exists {
		return fmt.Errorf("compose: node %q already exists in graph", id)
	}
	g.nodes[id] = &Node{
		ID:       id,
		Type:     nodeType,
		runnable: r,
	}
	g.adj[id] = []string{}
	return nil
}

// AddEdge adds a directed edge from node `from` to node `to`.
func (g *Graph) AddEdge(from, to string) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	if _, ok := g.nodes[from]; !ok {
		return fmt.Errorf("compose: source node %q not found", from)
	}
	if _, ok := g.nodes[to]; !ok {
		return fmt.Errorf("compose: destination node %q not found", to)
	}
	// Prevent self-loops, which would always form a cycle.
	if from == to {
		return fmt.Errorf("compose: self-loop on node %q is not allowed", from)
	}
	// Check for duplicate edges to avoid inflating in-degree counts during
	// topological sort, which would cause valid graphs to be rejected.
	for _, e := range g.edges {
		if e.From == from && e.To == to {
			return fmt.Errorf("compose: edge from %q to %q already exists", from, to)
		}
	}
	g.edges = append(g.edges, &Edge{From: from, To: to})
	g.adj[from] = append(g.adj[from], to)
	return nil
}

// topologicalOrder returns nodes in topological order using Kahn's algorithm.
// Returns an error if the graph contains a cycle.
func (g *Graph) topologicalOrder() ([]string, error) {
	inDegree := make(map[string]int, len(g.nodes))
	for id := range g.nodes {
		inDegree[id] = 0
	}
	for _, e := range g.edges {
		inDegree[e.To]++
	}

	queue := []string{}
	for id, deg := range inDegree {
		if deg == 0 {
			queue = append(queue, id)
		}
	}

	order := make([]string, 0, len(g.nodes))
	for len(queue) > 0 {
		cur := queue[0]
		qu