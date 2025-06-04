package graph_test

import (
	"testing"

	"github.com/droxer/algo/graph"
)

func TestAddEdge(t *testing.T) {
	g := graph.New()
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)

	if len(g.GetNeighbors(1)) != 1 || g.GetNeighbors(1)[0] != 2 {
		t.Errorf("Expected edge 1->2, got %v", g.GetNeighbors(1))
	}
	if len(g.GetNeighbors(2)) != 2 {
		t.Errorf("Expected 2 edges for vertex 2, got %d", len(g.GetNeighbors(2)))
	}
}

func TestBFS(t *testing.T) {
	g := graph.New()
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	g.AddEdge(1, 4)

	result := g.BFS(1)
	expected := []int{1, 2, 4, 3}

	if len(result) != len(expected) {
		t.Errorf("Expected BFS result length %d, got %d", len(expected), len(result))
	}

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("Expected BFS result[%d] = %d, got %d", i, expected[i], result[i])
		}
	}
}

func TestDFS(t *testing.T) {
	g := graph.New()
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	g.AddEdge(1, 4)

	result := g.DFS(1)
	expected := []int{1, 2, 3, 4}

	if len(result) != len(expected) {
		t.Errorf("Expected DFS result length %d, got %d", len(expected), len(result))
	}

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("Expected DFS result[%d] = %d, got %d", i, expected[i], result[i])
		}
	}
}

func TestHasCycle(t *testing.T) {
	g := graph.New()

	// Test graph without cycle
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)

	if g.HasCycle() {
		t.Error("Expected no cycle in linear graph")
	}

	// Test graph with cycle
	g2 := graph.New()
	g2.AddEdge(1, 2)
	g2.AddEdge(2, 3)
	g2.AddEdge(3, 1)

	if !g2.HasCycle() {
		t.Error("Expected cycle in cyclic graph")
	}
}
