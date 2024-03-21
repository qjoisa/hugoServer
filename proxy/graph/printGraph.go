package graph

import "math/rand"

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
	var name string
	for i := 0; i < count; i++ {
		if 65+i <= 90 {
			name = string(byte(65 + i))
		} else {
			name = string([]byte{byte(65 + i), byte(i - 24)})
		}
		node := &Node{ID: i + 1, Name: name}
		g.Nodes[i] = node
	}
	g.links(count)
	return g
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
		0: "circle",
		1: "rect",
		2: "square",
		3: "ellipse",
		5: "round-rect",
		6: "rhombus",
	}
	n.Form = m[rand.Intn(len(m))]
}
