package tree

import (
	"bytes"
	"log"
	"math/rand"
	"os"
	"time"
)

func Worker() {
	t := time.NewTicker(5 * time.Second)
	var first, last int
	rand.Seed(time.Now().Unix())
	fileData, err := os.ReadFile("/app/static/tasks/binary.md")
	if err != nil {
		log.Fatal(err)
	}
	avl := GenerateTree(5)
	lines := bytes.Split(fileData, []byte("\n"))
	for i, l := range lines {
		if bytes.Contains(l, []byte("{{< mermaid >}}")) {
			first = i
		}
		if bytes.Contains(l, []byte("{{< /mermaid >}}")) {
			last = i
		}
	}

	for {
		select {
		case <-t.C:
			avl.Insert(rand.Intn(1000))
			mm := avl.ToMermaid()
			fileData = append(fileData[:first], append([]byte(mm), fileData[last:]...)...)
			log.Fatal(os.WriteFile("/app/static/tasks/binary.md", bytes.Join(lines, []byte("\n")), 644))
		}
	}
}
