package familytree

import (
	"errors"
	"testing"
)

func TestAddNode(t *testing.T) {
	ft := NewFamilyTree()
	tests := []struct {
		scenario string
		id       int
		name     string
		metadata map[string]string
		err      error
	}{
		{
			scenario: "node gets added successfully",
			id:       1,
			name:     "A",
			metadata: nil,
			err:      nil,
		},
		{
			scenario: "node do not get added as it is repeated",
			id:       1,
			name:     "A",
			metadata: nil,
			err:      errors.New("node already exists"),
		},
	}

	for _, tc := range tests {
		err := ft.AddNode(tc.id, tc.name, tc.metadata)
		if err != nil && tc.err == nil {
			t.Errorf("for scenario: %s, expected %v, got %v", tc.scenario, tc.err, err)
		}
		if err == nil && tc.err != nil {
			t.Errorf("for scenario: %s, expected %v, got %v", tc.scenario, tc.err, err)
		}
	}
}

func TestAddEdge(t *testing.T) {
	ft := NewFamilyTree()
	ft.AddNode(1, "A", nil)
	ft.AddNode(2, "B", nil)
	ft.AddNode(3, "C", nil)
	ft.AddNode(4, "D", nil)
	ft.AddNode(5, "E", nil)
	ft.AddEdge(1, 2)
	ft.AddEdge(5, 2)
	ft.AddEdge(2, 3)
	ft.AddEdge(4, 3)

	tests := []struct {
		scenario string
		childID  int
		parentID int
		err      error
	}{
		{
			scenario: "a valid edge",
			childID:  2,
			parentID: 1,
			err:      nil,
		},
		{
			scenario: "a valid edge adding again",
			childID:  2,
			parentID: 1,
			err:      nil,
		},
		{
			scenario: "invalid edge parent and child same",
			childID:  2,
			parentID: 2,
			err:      errors.New("parent and child have same id"),
		},
		{
			scenario: "cyclic case: child trying to become a parent",
			childID:  1,
			parentID: 2,
			err:      errors.New("cyclic case"),
		},
		{
			scenario: "parent id not exist",
			childID:  1,
			parentID: 7,
			err:      errors.New("parent id not exist"),
		},
		{
			scenario: "child id not exist",
			childID:  9,
			parentID: 7,
			err:      errors.New("parent and child id not exist"),
		},
	}
	for _, tc := range tests {
		err := ft.AddEdge(tc.parentID, tc.childID)
		if tc.err != nil && err == nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		} else if tc.err == nil && err != nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		}
	}
}

func TestGetAncestors(t *testing.T) {
	ft := NewFamilyTree()
	ft.AddNode(1, "A", nil)
	ft.AddNode(2, "B", nil)
	ft.AddNode(3, "C", nil)
	ft.AddNode(4, "D", nil)
	ft.AddNode(5, "E", nil)
	ft.AddEdge(1, 2)
	ft.AddEdge(5, 2)
	ft.AddEdge(2, 3)
	ft.AddEdge(4, 3)

	tests := []struct {
		scenario  string
		nodeID    int
		ancestors []int
		err       error
	}{
		{
			scenario:  "List node ancestors when node exist",
			nodeID:    3,
			ancestors: []int{1, 5, 2, 4},
			err:       nil,
		}, {
			scenario:  "List node ancestors when node doesn't exist",
			nodeID:    12,
			ancestors: []int{},
			err:       errors.New("node does not exist"),
		}, {
			scenario:  "List node ancestors when node exist",
			nodeID:    1,
			ancestors: []int{},
			err:       nil,
		},
	}
	for _, tc := range tests {
		nodes, err := ft.GetAncestors(tc.nodeID)
		if tc.err != nil && err == nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		} else if tc.err == nil && err != nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		}

		if len(tc.ancestors) != len(nodes) {
			t.Errorf("Scenario: %s \n got no of ancestors: %v, expected no of ancestors: %v", tc.scenario, len(nodes), len(tc.ancestors))
		}

		for _, id := range tc.ancestors {
			if !isNodePresent(nodes, id) {
				t.Errorf("Scenario: %s \n %d got: false, expected: true", tc.scenario, id)
			}
		}
	}
}

func isNodePresent(slice []*node, target int) bool {
	for _, num := range slice {
		if num.id == target {
			return true
		}
	}
	return false
}

func TestGetParents(t *testing.T) {
	ft := NewFamilyTree()
	ft.AddNode(1, "A", nil)
	ft.AddNode(2, "B", nil)
	ft.AddNode(3, "C", nil)
	ft.AddNode(4, "D", nil)
	ft.AddNode(5, "E", nil)
	ft.AddEdge(1, 2)
	ft.AddEdge(5, 2)
	ft.AddEdge(2, 3)
	ft.AddEdge(4, 3)

	tests := []struct {
		scenario string
		nodeID   int
		parents  []int
		err      error
	}{
		{
			scenario: "List node parents when node exist",
			nodeID:   3,
			parents:  []int{2, 4},
			err:      nil,
		}, {
			scenario: "List node parents when node doesn't exist",
			nodeID:   12,
			parents:  []int{},
			err:      errors.New("node does not exist"),
		}, {
			scenario: "List node parents when node exist",
			nodeID:   1,
			parents:  []int{},
			err:      nil,
		},
	}
	for _, tc := range tests {
		nodes, err := ft.GetParents(tc.nodeID)
		if tc.err != nil && err == nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		} else if tc.err == nil && err != nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		}

		if len(tc.parents) != len(nodes) {
			t.Errorf("Scenario: %s \n got no of parents: %v, expected no of parents: %v", tc.scenario, len(nodes), len(tc.parents))
		}

		for _, id := range tc.parents {
			if !isNodePresent(nodes, id) {
				t.Errorf("Scenario: %s \n %d got: false, expected: true", tc.scenario, id)
			}
		}
	}
}

func TestGetChildren(t *testing.T) {
	ft := NewFamilyTree()
	ft.AddNode(1, "A", nil)
	ft.AddNode(2, "B", nil)
	ft.AddNode(3, "C", nil)
	ft.AddNode(4, "D", nil)
	ft.AddNode(5, "E", nil)
	ft.AddEdge(1, 2)
	ft.AddEdge(5, 2)
	ft.AddEdge(2, 3)
	ft.AddEdge(4, 3)

	tests := []struct {
		scenario string
		nodeID   int
		children []int
		err      error
	}{
		{
			scenario: "List node children when node exist",
			nodeID:   3,
			children: []int{},
			err:      nil,
		}, {
			scenario: "List node children when node doesn't exist",
			nodeID:   12,
			children: []int{},
			err:      errors.New("node does not exist"),
		}, {
			scenario: "List node children when node exist",
			nodeID:   1,
			children: []int{2},
			err:      nil,
		},
	}
	for _, tc := range tests {
		nodes, err := ft.GetChildren(tc.nodeID)
		if tc.err != nil && err == nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		} else if tc.err == nil && err != nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		}

		if len(tc.children) != len(nodes) {
			t.Errorf("Scenario: %s \n got no of children: %v, expected no of children: %v", tc.scenario, len(nodes), len(tc.children))
		}

		for _, id := range tc.children {
			if !isNodePresent(nodes, id) {
				t.Errorf("Scenario: %s \n %d got: false, expected: true", tc.scenario, id)
			}
		}
	}
}

func TestGetDescendants(t *testing.T) {
	ft := NewFamilyTree()
	ft.AddNode(1, "A", nil)
	ft.AddNode(2, "B", nil)
	ft.AddNode(3, "C", nil)
	ft.AddNode(4, "D", nil)
	ft.AddNode(5, "E", nil)
	ft.AddEdge(1, 2)
	ft.AddEdge(5, 2)
	ft.AddEdge(2, 3)
	ft.AddEdge(4, 3)

	tests := []struct {
		scenario    string
		nodeID      int
		descendants []int
		err         error
	}{
		{
			scenario:    "List node descendants when node exist",
			nodeID:      3,
			descendants: []int{},
			err:         nil,
		}, {
			scenario:    "List node descendants when node doesn't exist",
			nodeID:      12,
			descendants: []int{},
			err:         errors.New("node does not exist"),
		}, {
			scenario:    "List node descendants when node exist",
			nodeID:      1,
			descendants: []int{2, 3},
			err:         nil,
		},
	}
	for _, tc := range tests {
		nodes, err := ft.GetDescendants(tc.nodeID)
		if tc.err != nil && err == nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		} else if tc.err == nil && err != nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		}

		if len(tc.descendants) != len(nodes) {
			t.Errorf("Scenario: %s \n got no of descendants: %v, expected no of descendants: %v", tc.scenario, len(nodes), len(tc.descendants))
		}

		for _, id := range tc.descendants {
			if !isNodePresent(nodes, id) {
				t.Errorf("Scenario: %s \n %d got: false, expected: true", tc.scenario, id)
			}
		}
	}
}

func TestDeleteNode(t *testing.T) {
	ft := NewFamilyTree()
	ft.AddNode(1, "A", nil)
	ft.AddNode(2, "B", nil)
	ft.AddNode(3, "C", nil)
	ft.AddNode(4, "D", nil)
	ft.AddNode(5, "E", nil)
	ft.AddEdge(1, 2)
	ft.AddEdge(5, 2)
	ft.AddEdge(2, 3)
	ft.AddEdge(4, 3)
	ft.AddEdge(1, 3)

	tests := []struct {
		scenario string
		nodeID   int
		edges    []edge
		err      error
	}{
		{
			scenario: "Deleting an invalid node",
			nodeID:   9,
			edges:    []edge{},
			err:      errors.New("node does not exists"),
		}, {
			scenario: "Deleting a valid node",
			nodeID:   2,
			edges: []edge{
				{
					parentID: 1,
					childID:  3,
				}, {
					parentID: 4,
					childID:  3,
				},
			},
			err: nil,
		}, {
			scenario: "Deleting valid node which removes all edges ",
			nodeID:   3,
			edges:    []edge{},
			err:      nil,
		},
	}
	for _, tc := range tests {
		err := ft.DeleteNode(tc.nodeID)
		if tc.err != nil && err == nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		} else if tc.err == nil && err != nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		}

		if tc.err == nil && err == nil {
			edges := ft.GetEdges()
			if len(edges) != len(tc.edges) {
				t.Errorf("Scenario: %s \n got no of dependencies: %v, expected no of dependencies: %v", tc.scenario, len(edges), len(tc.edges))
			}
		}
	}
}

func TestDeleteEdge(t *testing.T) {
	ft := NewFamilyTree()
	ft.AddNode(1, "A", nil)
	ft.AddNode(2, "B", nil)
	ft.AddNode(3, "C", nil)
	ft.AddNode(4, "D", nil)
	ft.AddNode(5, "E", nil)
	ft.AddEdge(1, 2)
	ft.AddEdge(5, 2)
	ft.AddEdge(2, 3)
	ft.AddEdge(4, 3)
	ft.AddEdge(1, 3)

	tests := []struct {
		scenario       string
		childID        int
		parentID       int
		err            error
		remainingEdges int
	}{
		{
			scenario: "an invalid edge parent invalid",
			childID:  2,
			parentID: 6,
			err:      errors.New("parent is invalid"),
		},
		{
			scenario: "an invalid edge child is invalid",
			childID:  6,
			parentID: 1,
			err:      errors.New("child is invalid"),
		},
		{
			scenario: "invalid edge parent and child same",
			childID:  2,
			parentID: 2,
			err:      errors.New("parent and child have same id"),
		},
		{
			scenario: "no edge exists between two nodes",
			childID:  1,
			parentID: 5,
			err:      errors.New("no edge from parent to child"),
		},
		{
			scenario:       "valid edge case 1",
			childID:        2,
			parentID:       1,
			err:            nil,
			remainingEdges: 4,
		},
		{
			scenario:       "valid edge case 2",
			childID:        3,
			parentID:       2,
			err:            nil,
			remainingEdges: 3,
		},
	}
	for _, tc := range tests {
		err := ft.DeleteEdge(tc.parentID, tc.childID)
		if tc.err != nil && err == nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		} else if tc.err == nil && err != nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		}

		if tc.err == nil && err == nil {
			edges := ft.GetEdges()
			if len(edges) != tc.remainingEdges {
				t.Errorf("Scenario: %s \n got no of dependencies: %v, expected no of dependencies: %v", tc.scenario, len(edges), tc.remainingEdges)
			}
		}
	}
}
