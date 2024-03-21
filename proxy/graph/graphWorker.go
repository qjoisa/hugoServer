package graph

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

var start = "---\nmenu:\n    after:\n        name: graph\n        weight: 1\ntitle: Построение графа\n---\n\n# Построение графа\n{{< columns >}}\n"
var end = "\n{{< /columns >}}"
var path = "/app/static/tasks/graph.md"

func Worker() {
	t := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-t.C:
			g := NewGraph(rand.Intn(26) + 5)
			s := g.ToMermaid()
			res := fmt.Sprintf("%s%s%s", start, s, end)
			err := os.WriteFile(path, []byte(res), 0644)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
