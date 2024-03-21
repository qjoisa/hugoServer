package graph

import (
	"fmt"
	"math/rand"
	"slices"
)

type Node struct {
	ID    int
	Name  string
	Form  string // "circle", "rect", "square", "ellipse", "round-rect", "rhombus"
	Links []*Node
}

type Graph struct {
	Nodes []*Node
}

func NewGraph(count int) *Graph {
	g := &Graph{Nodes: make([]*Node, count)}
	for i := 0; i < count; i++ {
		node := &Node{ID: i + 1, Name: fmt.Sprintf("id%d", i+1)}
		node.randomForm()
		g.Nodes[i] = node
	}
	g.links()
	return g
}

func (g *Graph) ToMermaid() string {
	var graph string
	for _, node := range g.Nodes {
		for _, linked := range node.Links {
			graph += fmt.Sprintf("%s%s --> %s%s\n", node.Name, node.Form, linked.Name, linked.Form)
		}
	}
	return "{{< mermaid >}}\ngraph LR\n" + graph + "{{< /mermaid >}}\n"
}

func (g *Graph) links() {
	count := len(g.Nodes) - 1
	for i := 0; i < count-1; i++ {
		linksCount := rand.Intn(count)
		l := make([]*Node, 0, linksCount)
		for j := 0; j < linksCount; {
			node := g.Nodes[rand.Intn(count)]
			if node == g.Nodes[i] || slices.Contains(l, node) {
				continue
			}
			if !slices.Contains(node.Links, g.Nodes[i]) {
				l = append(l, node)
			}
			j++
		}
		g.Nodes[i].Links = l
	}
}

func (n *Node) randomForm() {
	m := map[int]string{
		0: "((circle))",
		1: "[rect]",
		2: "([ellipse])",
		3: "(round-rect)",
		4: "{rhombus}",
	}
	n.Form = m[rand.Intn(len(m))]
}
