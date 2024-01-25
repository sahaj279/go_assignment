package familytree

import "github.com/pkg/errors"

// 1 AddNode method.
func (ft *familyTree) AddNode(id int, name string, metadata map[string]string) error {
	if _, exists := ft.nodes[id]; exists {
		return errors.Errorf(NodeAlreadyExists, id)
	}

	node := NewNode(id, name, metadata)
	ft.nodes[id] = node
	return nil
}

// 2 AddEdge method.
func (ft *familyTree) AddEdge(parentID, childID int) error {
	if childID == parentID {
		return errors.New("child and parent have same id")
	}

	if _, exists := ft.nodes[childID]; !exists {
		return errors.Errorf(NodeDoesNotExist, childID)
	}

	if _, exists := ft.nodes[parentID]; !exists {
		return errors.Errorf(NodeDoesNotExist, parentID)
	}

	ancestors, _ := ft.GetAncestors(parentID)
	for id := range ancestors {
		if id == childID {
			return errors.Errorf("the edge with parent %d and child %d will cause a cycle", parentID, childID)
		}
	}

	if ft.nodes[childID].parents == nil {
		ft.nodes[childID].parents = make(map[int]*node)
	}

	ft.nodes[childID].parents[parentID] = ft.nodes[parentID]
	if ft.nodes[parentID].children == nil {
		ft.nodes[parentID].children = make(map[int]*node)
	}

	ft.nodes[parentID].children[childID] = ft.nodes[childID]

	return nil
}

// 3 GetAncestors method.
func (ft *familyTree) GetAncestors(id int) ([]*node, error) {
	if _, exists := ft.nodes[id]; !exists {
		return nil, errors.Errorf(NodeDoesNotExist, id)
	}
	return getAncestorsDfs(ft.nodes[id]), nil
}

func getAncestorsDfs(n *node) []*node {
	ancestors := make([]*node, 0)
	for id := range n.parents {
		ancestors = append(ancestors, n.parents[id])
		ancestors = append(ancestors, getAncestorsDfs(n.parents[id])...)
	}
	return ancestors
}

// 4 GetParents method.
func (ft *familyTree) GetParents(id int) ([]*node, error) {
	if _, exists := ft.nodes[id]; !exists {
		return nil, errors.Errorf(NodeDoesNotExist, id)
	}

	parents := make([]*node, 0)
	for i := range ft.nodes[id].parents {
		parents = append(parents, ft.nodes[id].parents[i])
	}
	return parents, nil
}

// 5 GetChildren method.
func (ft *familyTree) GetChildren(id int) ([]*node, error) {
	if _, exists := ft.nodes[id]; !exists {
		return nil, errors.Errorf(NodeDoesNotExist, id)
	}

	children := make([]*node, 0)
	for i := range ft.nodes[id].children {
		children = append(children, ft.nodes[id].children[i])
	}
	return children, nil
}

// 6 GetDescendants method.
func (ft *familyTree) GetDescendants(id int) ([]*node, error) {
	if _, exists := ft.nodes[id]; !exists {
		return nil, errors.Errorf(NodeDoesNotExist, id)
	}

	return getDescendantsDfs(ft.nodes[id]), nil
}

func getDescendantsDfs(n *node) []*node {
	descendants := make([]*node, 0)
	for id := range n.children {
		descendants = append(descendants, n.children[id])
		descendants = append(descendants, getDescendantsDfs(n.children[id])...)
	}
	return descendants
}

// 7 DeleteNode method.
func (ft *familyTree) DeleteNode(id int) error {
	if _, exists := ft.nodes[id]; !exists {
		return errors.Errorf(NodeDoesNotExist, id)
	}

	defer delete(ft.nodes, id)
	// removing this node from its children who have this node as a parent
	for node := range ft.nodes[id].children {
		delete(ft.nodes[node].parents, id)
	}
	// removing this node from its parents who have this node as a child
	for node := range ft.nodes[id].parents {
		delete(ft.nodes[node].children, id)
	}

	return nil
}

// 8 GetDependencies method.
func (ft *familyTree) GetEdges() []edge {
	edges := make([]edge, 0)
	for node := range ft.nodes {
		for child := range ft.nodes[node].children {
			edges = append(edges, edge{
				parentID: node,
				childID:  child,
			})
		}
	}
	return edges
}

// 9 DeleteDependency method.
func (ft *familyTree) DeleteEdge(parentID, childID int) error {
	if childID == parentID {
		return errors.New("child and parent have same id")
	}

	if _, exists := ft.nodes[childID]; !exists {
		return errors.Errorf(NodeDoesNotExist, childID)
	}

	if _, exists := ft.nodes[parentID]; !exists {
		return errors.Errorf(NodeDoesNotExist, parentID)
	}

	// confirming whether this dependency exists
	if _, exists := ft.nodes[parentID].children[childID]; !exists {
		return errors.Errorf("No dependency between parent %d and child %d", parentID, childID)
	}

	delete(ft.nodes[parentID].children, childID)
	delete(ft.nodes[childID].parents, parentID)
	return nil
}
