package tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var root = &Node{Key: 3, Height: 1}
var nodeR = &Node{Key: 7, Height: 0}
var nodeL = &Node{Key: 1, Height: 0}

func TestGenerateTree(t *testing.T) {
	type args struct {
		count int
	}
	tests := []struct {
		name       string
		args       args
		wantStruct *AVLTree
		want       int
	}{
		{
			name:       "good gen",
			args:       args{count: 3},
			wantStruct: &AVLTree{},
			want:       3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateTree(3)
			assert.IsTypef(t, tt.wantStruct, got, "types not *AVLtree")
			assert.Equal(t, got.nCount, tt.want)
		})
	}
}

func TestAVLTree_Insert(t *testing.T) {
	root.Right = nodeR
	root.Left = nodeL
	type want struct {
		tree *AVLTree
	}
	type args struct {
		key []int
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "good insert",
			args: args{key: []int{1, 3, 7}},
			want: want{
				tree: &AVLTree{Root: root, nCount: 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := &AVLTree{}
			for _, v := range tt.args.key {
				res.Insert(v)
			}
			assert.Equal(t, tt.want.tree.nCount, res.nCount)
			assert.ObjectsAreEqualValues(tt.want, res)
		})
	}
}

func TestAVLTree_ToMermaid(t *testing.T) {
	root.Right = nodeR
	root.Left = nodeL
	type fields struct {
		tree *AVLTree
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Test",
			fields: fields{tree: &AVLTree{Root: root}},
			want:   "{{< mermaid >}}\ngraph TD\n3--> 1\n3--> 7\n\n{{< /mermaid >}}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := tt.fields.tree.ToMermaid()
			assert.Equal(t, tt.want, res)

		})
	}
}
