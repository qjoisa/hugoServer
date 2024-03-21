package graph

import (
	"fmt"
	"math/rand"
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
		g.Nodes[i] = node
	}
	g.links(count)
	return g
}

func (g *Graph) ToMermaid() string {
	var graph string
	for _, node := range g.Nodes {
		for _, linked := range node.Links {
			if linked.Form == "square" {
				graph += fmt.Sprintf("%s%s --> %s\n", node.Name, node.Form, linked.Form)
			} else if linked.Form == "square" && node.Form == "square" {
				graph += fmt.Sprintf("%s --> %s\n", node.Form, linked.Form)
			} else if node.Form == "square" {
				graph += fmt.Sprintf("%s --> %s%s\n", node.Form, linked.Name, linked.Form)
			} else {
				graph += fmt.Sprintf("%s%s --> %s%s\n", node.Name, node.Form, linked.Name, linked.Form)
			}
		}
	}
	return "{{< mermaid >}}\n" + graph + "{{< /mermaid >}}\n"
}

func (g *Graph) links(count int) {
	for i := 0; i < count; i++ {
		linksCount := rand.Intn(count)
		l := make([]*Node, linksCount)
		for j := 0; j < linksCount; {
			node := g.Nodes[rand.Intn(count)]
			if node == g.Nodes[i] {
				continue
			}
			l[j] = node
			j++
		}
		g.Nodes[i].Links = l
	}
}

func (n *Node) randomForm() {
	m := map[int]string{
		0: "((circle))",
		1: "[rect]",
		2: "square",
		3: "([ellipse])",
		5: "(round-rect)",
		6: "{rhombus}",
	}
	n.Form = m[rand.Intn(len(m))]
}
