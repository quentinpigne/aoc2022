package utils

type TreeNode struct {
	name     string
	size     int
	nodeType string
	parent   *TreeNode
	children []*TreeNode
}

// Getters

func (t *TreeNode) GetName() string {
	return t.name
}

func (t *TreeNode) GetSize() int {
	return t.size
}

func (t *TreeNode) GetType() string {
	return t.nodeType
}

func (t *TreeNode) GetParent() *TreeNode {
	return t.parent
}

func (t *TreeNode) GetChildren() []*TreeNode {
	return t.children
}

func (t *TreeNode) FindChild(name string) *TreeNode {
	var foundTreeNode *TreeNode
	for _, treeNode := range t.children {
		if treeNode.name == name {
			foundTreeNode = treeNode
		}
	}
	return foundTreeNode
}

// Manipulation

func (t *TreeNode) AddDirectory(name string) {
	newDir := new(TreeNode)
	newDir.name = name
	newDir.nodeType = "dir"
	newDir.parent = t
	t.children = append(t.children, newDir)
}

func (t *TreeNode) AddFile(name string, size int) {
	newFile := new(TreeNode)
	newFile.name = name
	newFile.size = size
	newFile.nodeType = "file"
	newFile.parent = t
	t.UpdateSize(size)
	t.children = append(t.children, newFile)
}

func (t *TreeNode) UpdateSize(size int) {
	t.size += size
	if t.parent != nil {
		t.parent.UpdateSize(size)
	}
}
