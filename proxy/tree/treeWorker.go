package tree

import (
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

var path = "/app/static/tasks/binary.md"

func Worker() {
	t := time.NewTicker(5 * time.Second)
	rand.Seed(time.Now().Unix())
	fileData, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	avl := GenerateTree(5)
	for {
		select {
		case <-t.C:
			if avl.nCount >= 100 {
				GenerateTree(5)
			} else {
				avl.Insert(rand.Intn(1000))
			}
			mm := avl.ToMermaid()
			lines, first, last := where(fileData)
			lines = append(lines[:first], append([]string{mm}, lines[last+1:]...)...)
			log.Println(mm)
			err := os.WriteFile("/app/static/tasks/binary.md", []byte(strings.Join(lines, "\n")), 644)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func where(fileData []byte) ([]string, int, int) {
	var first, last int
	lines := strings.Split(string(fileData), "\n")
	for i, l := range lines {
		if strings.Contains(l, "{{< mermaid >}}") {
			first = i
		}
		if strings.Contains(l, "{{< /mermaid >}}") {
			last = i
		}
	}
	return lines, first, last
}
