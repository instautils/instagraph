package graph

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"sort"
)

type Graph struct {
	nodes     map[string]int
	edges     []Edge
	edgeCache map[string]bool
}

func New() *Graph {
	return &Graph{
		edges:     make([]Edge, 0),
		nodes:     make(map[string]int),
		edgeCache: make(map[string]bool),
	}
}

func generateID(personA, personB string) string {
	slice := []string{personA, personB}
	sort.Strings(slice)
	personA, personB = slice[0], slice[1]
	return fmt.Sprintf("%s%s", personA, personB)
}

func (g *Graph) AddConnection(personA, personB string) {
	g.nodes[personA]++
	g.nodes[personB]++

	id := generateID(personA, personB)
	if _, ok := g.edgeCache[id]; !ok {
		g.edges = append(g.edges, Edge{
			ID:     id,
			Source: personA,
			Target: personB,
		})
		g.edgeCache[id] = true
	}
}

func random(a int) int {
	return rand.Intn(a) - a/2
}

func (g *Graph) Marshall() []byte {
	nodes := make([]Node, 0)

	iteration, offset := 1, 100
	for user, size := range g.nodes {
		nodes = append(nodes, Node{
			ID:    user,
			Label: user,
			Size:  size * 2,
			X:     size + random(iteration+offset),
			Y:     size + random(iteration+offset),
		})
		iteration++
	}

	bytes, err := json.Marshal(map[string]interface{}{
		"edges": g.edges,
		"nodes": nodes,
	})
	if err != nil {
		log.Println(err)
	}
	return bytes
}
