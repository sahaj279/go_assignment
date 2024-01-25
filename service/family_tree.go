package familytree

const (
	NodeAlreadyExists = "node already exists with id %d"
	NodeDoesNotExist  = "node with id %d does not exist"
)

type node struct {
	id       int
	name     string
	metadata map[string]string
	children map[int]*node
	parents  map[int]*node
}

type familyTree struct {
	nodes map[int]*node
}

type edge struct {
	parentID int
	childID  int
}

type Svc interface {
	AddNode(id int, name string, metadata map[string]string) error
	AddEdge(childID, parentID int) error
	GetAncestors(id int) ([]*node, error)
	GetChildren(id int) ([]*node, error)
	GetParents(id int) ([]*node, error)
	GetDescendants(id int) ([]*node, error)
	DeleteNode(id int) error
	GetEdges() []edge
	DeleteEdge(parentID, childID int) error
}

func NewFamilyTree() *familyTree {
	return &familyTree{
		nodes: make(map[int]*node),
	}
}

func NewNode(id int, name string, metadata map[string]string) *node {
	return &node{
		id:       id,
		name:     name,
		metadata: metadata,
	}
}
