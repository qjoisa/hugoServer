package graph

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestNewGraph(t *testing.T) {
	tests := []struct {
		name  string
		count int
	}{
		{
			name:  "graph",
			count: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewGraph(tt.count)
			assert.Equal(t, tt.count, len(got.Nodes))
		})
	}
}

func TestGraph_ToMermaid(t *testing.T) {
	type fields struct {
		graph *Graph
	}
	tests := []struct {
		name   string
		fields fields
		pref   string
		suf    string
	}{
		{
			name:   "mermaid",
			fields: fields{graph: NewGraph(6)},
			pref:   "{{< mermaid >}}\ngraph LR\n",
			suf:    "{{< /mermaid >}}\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := tt.fields.graph.ToMermaid()

			if !strings.HasPrefix(res, tt.pref) {
				t.Error("not prefix")
			}
			if !strings.HasSuffix(res, tt.suf) {
				t.Error("not suffix")
			}
		})
	}
}
