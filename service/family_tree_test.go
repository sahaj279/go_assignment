package familytree

import (
	"errors"
	"testing"
)

func TestAddNode(t *testing.T) {
	tests := []struct {
		scenario string
		ft       familyTree
		id       int
		name     string
		metadata map[string]string
		err      error
	}{
		{
			scenario: "node gets added successfully",
			ft: familyTree{
				nodes: map[int]*node{},
			},
			id:       1,
			name:     "A",
			metadata: nil,
			err:      nil,
		},
		{
			scenario: "node do not get added as it is repeated",
			ft: familyTree{
				nodes: map[int]*node{
					1: {
						id: 1,
					},
				},
			},
			id:       1,
			name:     "A",
			metadata: nil,
			err:      errors.New("node already exists"),
		},
	}

	for _, tc := range tests {
		err := tc.ft.AddNode(tc.id, tc.name, tc.metadata)
		if err != nil && tc.err == nil {
			t.Errorf("for scenario: %s, expected %v, got %v", tc.scenario, tc.err, err)
		}
		if err == nil && tc.err != nil {
			t.Errorf("for scenario: %s, expected %v, got %v", tc.scenario, tc.err, err)
		}
	}
}

func TestAddEdge(t *testing.T) {
	tests := []struct {
		scenario string
		ft       familyTree
		childID  int
		parentID int
		err      error
	}{
		{
			scenario: "a valid edge",
			ft: familyTree{
				nodes: map[int]*node{
					1: {
						id: 1,
						children: map[int]*node{
							2: {
								id: 2,
							},
						},
					},
					2: {
						id: 2,
						parents: map[int]*node{
							1: {
								id: 1,
							},
						},
						children: map[int]*node{
							3: {
								id: 3,
							},
						},
					},
					3: {
						id: 3,
						parents: map[int]*node{
							2: {
								id: 2,
							},
						},
					},
				},
			},
			childID:  2,
			parentID: 1,
			err:      nil,
		},
		{
			scenario: "a valid edge adding again",
			ft: familyTree{
				nodes: map[int]*node{
					1: {
						id: 1,
						children: map[int]*node{
							2: {
								id: 2,
							},
						},
					},
					2: {
						id: 2,
						parents: map[int]*node{
							1: {
								id: 1,
							},
						},
					},
					3: {
						id: 3,
					},
				},
			},
			childID:  3,
			parentID: 2,
			err:      nil,
		},
		{
			scenario: "invalid edge parent and child same",
			ft: familyTree{
				nodes: map[int]*node{
					1: {
						id: 1,
					},
				},
			},
			childID:  1,
			parentID: 1,
			err:      errors.New("parent and child have same id"),
		},
		{
			scenario: "cyclic case: child trying to become a parent",
			ft: familyTree{
				nodes: map[int]*node{
					1: {
						id: 1,
						children: map[int]*node{
							2: {
								id: 2,
							},
						},
					},
					2: {
						id: 2,
						parents: map[int]*node{
							1: {
								id: 1,
							},
						},
						children: map[int]*node{
							3: {
								id: 3,
							},
						},
					},
					3: {
						id: 3,
						parents: map[int]*node{
							2: {
								id: 2,
							},
						},
					},
				},
			},
			childID:  1,
			parentID: 2,
			err:      errors.New("cyclic case"),
		},
		{
			scenario: "parent id not exist",
			ft: familyTree{
				nodes: map[int]*node{
					1: {
						id: 1,
						children: map[int]*node{
							2: {
								id: 2,
							},
						},
					},
					2: {
						id: 2,
						parents: map[int]*node{
							1: {
								id: 1,
							},
						},
						children: map[int]*node{
							3: {
								id: 3,
							},
						},
					},
					3: {
						id: 3,
						parents: map[int]*node{
							2: {
								id: 2,
							},
						},
					},
				},
			},
			childID:  1,
			parentID: 7,
			err:      errors.New("parent id not exist"),
		},
		{
			scenario: "child id not exist",
			childID:  9,
			parentID: 1,
			err:      errors.New("parent and child id not exist"),
		},
	}
	for _, tc := range tests {
		err := tc.ft.AddEdge(tc.parentID, tc.childID)
		if tc.err != nil && err == nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		} else if tc.err == nil && err != nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		}
	}
}

func TestGetAncestors(t *testing.T) {
	tests := []struct {
		scenario  string
		ft        familyTree
		nodeID    int
		ancestors []int
		err       error
	}{
		{
			scenario: "List node ancestors when node exist",
			ft: familyTree{
				nodes: map[int]*node{
					1: {
						id: 1,
						children: map[int]*node{
							2: {
								id: 2,
							},
						},
					},
					2: {
						id: 2,
						parents: map[int]*node{
							1: {
								id: 1,
							},
						},
						children: map[int]*node{
							3: {
								id: 3,
							},
						},
					},
					3: {
						id: 3,
						parents: map[int]*node{
							2: {
								id: 2,
							},
						},
					},
				},
			},
			nodeID:    2,
			ancestors: []int{1},
			err:       nil,
		}, {
			scenario: "List node ancestors when node doesn't exist",
			ft: familyTree{
				nodes: map[int]*node{
					1: {
						id: 1,
						children: map[int]*node{
							2: {
								id: 2,
							},
						},
					},
					2: {
						id: 2,
						parents: map[int]*node{
							1: {
								id: 1,
							},
						},
						children: map[int]*node{
							3: {
								id: 3,
							},
						},
					},
					3: {
						id: 3,
						parents: map[int]*node{
							2: {
								id: 2,
							},
						},
					},
				},
			},
			nodeID:    12,
			ancestors: []int{},
			err:       errors.New("node does not exist"),
		}, {
			scenario: "List node ancestors when node exist",
			ft: familyTree{
				nodes: map[int]*node{
					1: {
						id: 1,
						children: map[int]*node{
							2: {
								id: 2,
							},
						},
					},
					2: {
						id: 2,
						parents: map[int]*node{
							1: {
								id: 1,
							},
						},
						children: map[int]*node{
							3: {
								id: 3,
							},
						},
					},
					3: {
						id: 3,
						parents: map[int]*node{
							2: {
								id: 2,
							},
						},
					},
				},
			},
			nodeID:    1,
			ancestors: []int{},
			err:       nil,
		},
	}
	for _, tc := range tests {
		nodes, err := tc.ft.GetAncestors(tc.nodeID)
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
	tests := []struct {
		scenario string
		nodeID   int
		ft       familyTree
		parents  []int
		err      error
	}{
		{
			scenario: "List node parents when node exist",
			ft: familyTree{
				nodes: map[int]*node{
					1: {
						id: 1,
						children: map[int]*node{
							2: {
								id: 2,
							},
						},
					},
					2: {
						id: 2,
						parents: map[int]*node{
							1: {
								id: 1,
							},
						},
						children: map[int]*node{
							3: {
								id: 3,
							},
						},
					},
					3: {
						id: 3,
						parents: map[int]*node{
							2: {
								id: 2,
							},
						},
					},
				},
			},
			nodeID:  2,
			parents: []int{1},
			err:     nil,
		}, {
			scenario: "List node parents when node doesn't exist",
			ft: familyTree{
				nodes: map[int]*node{
					1: {
						id: 1,
						children: map[int]*node{
							2: {
								id: 2,
							},
						},
					},
					2: {
						id: 2,
						parents: map[int]*node{
							1: {
								id: 1,
							},
						},
						children: map[int]*node{
							3: {
								id: 3,
							},
						},
					},
					3: {
						id: 3,
						parents: map[int]*node{
							2: {
								id: 2,
							},
						},
					},
				},
			},
			nodeID:  12,
			parents: []int{},
			err:     errors.New("node does not exist"),
		}, {
			scenario: "List node parents when node exist",
			ft: familyTree{
				nodes: map[int]*node{
					1: {
						id: 1,
						children: map[int]*node{
							2: {
								id: 2,
							},
						},
					},
					2: {
						id: 2,
						parents: map[int]*node{
							1: {
								id: 1,
							},
						},
						children: map[int]*node{
							3: {
								id: 3,
							},
						},
					},
					3: {
						id: 3,
						parents: map[int]*node{
							2: {
								id: 2,
							},
						},
					},
				},
			},
			nodeID:  1,
			parents: []int{},
			err:     nil,
		},
	}
	for _, tc := range tests {
		nodes, err := tc.ft.GetParents(tc.nodeID)
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
	tests := []struct {
		scenario string
		nodeID   int
		ft       familyTree
		children []int
		err      error
	}{
		{
			scenario: "List node children when node exist",
			ft: familyTree{
				nodes: map[int]*node{
					1: {
						id: 1,
						children: map[int]*node{
							2: {
								id: 2,
							},
						},
					},
					2: {
						id: 2,
						parents: map[int]*node{
							1: {
								id: 1,
							},
						},
						children: map[int]*node{
							3: {
								id: 3,
							},
						},
					},
					3: {
						id: 3,
						parents: map[int]*node{
							2: {
								id: 2,
							},
						},
					},
				},
			},
			nodeID:   2,
			children: []int{3},
			err:      nil,
		}, {
			scenario: "List node children when node doesn't exist",
			ft: familyTree{
				nodes: map[int]*node{
					1: {
						id: 1,
						children: map[int]*node{
							2: {
								id: 2,
							},
						},
					},
					2: {
						id: 2,
						parents: map[int]*node{
							1: {
								id: 1,
							},
						},
						children: map[int]*node{
							3: {
								id: 3,
							},
						},
					},
					3: {
						id: 3,
						parents: map[int]*node{
							2: {
								id: 2,
							},
						},
					},
				},
			},
			nodeID:   12,
			children: []int{},
			err:      errors.New("node does not exist"),
		}, {
			scenario: "List node children when node exist",
			ft: familyTree{
				nodes: map[int]*node{
					1: {
						id: 1,
						children: map[int]*node{
							2: {
								id: 2,
							},
						},
					},
					2: {
						id: 2,
						parents: map[int]*node{
							1: {
								id: 1,
							},
						},
						children: map[int]*node{
							3: {
								id: 3,
							},
						},
					},
					3: {
						id: 3,
						parents: map[int]*node{
							2: {
								id: 2,
							},
						},
					},
				},
			},
			nodeID:   1,
			children: []int{2},
			err:      nil,
		},
	}
	for _, tc := range tests {
		nodes, err := tc.ft.GetChildren(tc.nodeID)
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
	tests := []struct {
		scenario    string
		ft          familyTree
		nodeID      int
		descendants []int
		err         error
	}{
		{
			scenario: "List node descendants when node exist",
			ft: familyTree{
				nodes: map[int]*node{
					1: {
						id: 1,
						children: map[int]*node{
							2: {
								id: 2,
							},
						},
					},
					2: {
						id: 2,
						parents: map[int]*node{
							1: {
								id: 1,
							},
						},
						children: map[int]*node{
							3: {
								id: 3,
							},
						},
					},
					3: {
						id: 3,
						parents: map[int]*node{
							2: {
								id: 2,
							},
						},
					},
				},
			},
			nodeID:      2,
			descendants: []int{3},
			err:         nil,
		}, {
			scenario: "List node descendants when node doesn't exist",
			ft: familyTree{
				nodes: map[int]*node{
					1: {
						id: 1,
						children: map[int]*node{
							2: {
								id: 2,
							},
						},
					},
					2: {
						id: 2,
						parents: map[int]*node{
							1: {
								id: 1,
							},
						},
						children: map[int]*node{
							3: {
								id: 3,
							},
						},
					},
					3: {
						id: 3,
						parents: map[int]*node{
							2: {
								id: 2,
							},
						},
					},
				},
			},
			nodeID:      12,
			descendants: []int{},
			err:         errors.New("node does not exist"),
		}, {
			scenario: "List node descendants when node exist",
			ft: familyTree{
				nodes: map[int]*node{
					1: {
						id: 1,
						children: map[int]*node{
							2: {
								id: 2, children: map[int]*node{
									3: {
										id: 3,
									},
								},
							},
						},
					},
					2: {
						id: 2,
						parents: map[int]*node{
							1: {
								id: 1,
							},
						},
						children: map[int]*node{
							3: {
								id: 3,
							},
						},
					},
					3: {
						id: 3,
						parents: map[int]*node{
							2: {
								id: 2,
							},
						},
					},
				},
			},
			nodeID:      1,
			descendants: []int{2, 3},
			err:         nil,
		},
	}
	for _, tc := range tests {
		nodes, err := tc.ft.GetDescendants(tc.nodeID)
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
				t.Errorf("Scenario: %s \n id not present %d got: false, expected: true", tc.scenario, id)
			}
		}
	}
}

func TestDeleteNode(t *testing.T) {
	tests := []struct {
		scenario  string
		nodeID    int
		edgeCount int
		ft        familyTree
		err       error
	}{
		{
			scenario: "Deleting an invalid node",
			ft: familyTree{
				nodes: map[int]*node{
					1: {
						id: 1,
						children: map[int]*node{
							2: {
								id: 2,
							},
						},
					},
					2: {
						id: 2,
						parents: map[int]*node{
							1: {
								id: 1,
							},
						},
						children: map[int]*node{
							3: {
								id: 3,
							},
						},
					},
					3: {
						id: 3,
						parents: map[int]*node{
							2: {
								id: 2,
							},
						},
					},
				},
			},
			nodeID:    9,
			edgeCount: 0,
			err:       errors.New("node does not exists"),
		}, {
			scenario: "Deleting a valid node",
			nodeID:   2,
			ft: familyTree{
				nodes: map[int]*node{
					1: {
						id: 1,
						children: map[int]*node{
							2: {
								id: 2,
							},
						},
					},
					2: {
						id: 2,
						parents: map[int]*node{
							1: {
								id: 1,
							},
						},
						children: map[int]*node{
							3: {
								id: 3,
							},
						},
					},
					3: {
						id: 3,
						parents: map[int]*node{
							2: {
								id: 2,
							},
						},
					},
				},
			},
			edgeCount: 0,
			err:       nil,
		}, {
			scenario: "Deleting valid node which removes all edges ",
			ft: familyTree{
				nodes: map[int]*node{
					1: {
						id: 1,
						children: map[int]*node{
							2: {
								id: 2,
							},
						},
					},
					2: {
						id: 2,
						parents: map[int]*node{
							1: {
								id: 1,
							},
						},
						children: map[int]*node{
							3: {
								id: 3,
							},
						},
					},
					3: {
						id: 3,
						parents: map[int]*node{
							2: {
								id: 2,
							},
						},
					},
				},
			},
			nodeID:    2,
			edgeCount: 0,
			err:       nil,
		},
	}
	for _, tc := range tests {
		err := tc.ft.DeleteNode(tc.nodeID)
		if tc.err != nil && err == nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		} else if tc.err == nil && err != nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		}

		if tc.err == nil && err == nil {
			edgeCount := tc.ft.GetEdgeCount()
			if edgeCount != tc.edgeCount {
				t.Errorf("Scenario: %s \n got no of dependencies: %v, expected no of dependencies: %v", tc.scenario, edgeCount, tc.edgeCount)
			}
		}
	}
}

func TestDeleteEdge(t *testing.T) {
	tests := []struct {
		scenario       string
		childID        int
		parentID       int
		err            error
		remainingEdges int
		ft             familyTree
	}{
		{
			scenario: "an invalid edge parent invalid",
			ft: familyTree{
				nodes: map[int]*node{
					1: {
						id: 1,
						children: map[int]*node{
							2: {
								id: 2,
							},
						},
					},
					2: {
						id: 2,
						parents: map[int]*node{
							1: {
								id: 1,
							},
						},
						children: map[int]*node{
							3: {
								id: 3,
							},
						},
					},
					3: {
						id: 3,
						parents: map[int]*node{
							2: {
								id: 2,
							},
						},
					},
				},
			},
			childID:  2,
			parentID: 6,
			err:      errors.New("parent is invalid"),
		},
		{
			scenario: "an invalid edge child is invalid",
			ft: familyTree{
				nodes: map[int]*node{
					1: {
						id: 1,
						children: map[int]*node{
							2: {
								id: 2,
							},
						},
					},
					2: {
						id: 2,
						parents: map[int]*node{
							1: {
								id: 1,
							},
						},
						children: map[int]*node{
							3: {
								id: 3,
							},
						},
					},
					3: {
						id: 3,
						parents: map[int]*node{
							2: {
								id: 2,
							},
						},
					},
				},
			},
			childID:  6,
			parentID: 1,
			err:      errors.New("child is invalid"),
		},
		{
			scenario: "invalid edge parent and child same",
			childID:  2,
			parentID: 2,
			ft: familyTree{
				nodes: map[int]*node{
					1: {
						id: 1,
						children: map[int]*node{
							2: {
								id: 2,
							},
						},
					},
					2: {
						id: 2,
						parents: map[int]*node{
							1: {
								id: 1,
							},
						},
						children: map[int]*node{
							3: {
								id: 3,
							},
						},
					},
					3: {
						id: 3,
						parents: map[int]*node{
							2: {
								id: 2,
							},
						},
					},
				},
			},
			err: errors.New("parent and child have same id"),
		},
		{
			scenario: "no edge exists between two nodes",
			childID:  1,
			parentID: 3,
			ft: familyTree{
				nodes: map[int]*node{
					1: {
						id: 1,
						children: map[int]*node{
							2: {
								id: 2,
							},
						},
					},
					2: {
						id: 2,
						parents: map[int]*node{
							1: {
								id: 1,
							},
						},
						children: map[int]*node{
							3: {
								id: 3,
							},
						},
					},
					3: {
						id: 3,
						parents: map[int]*node{
							2: {
								id: 2,
							},
						},
					},
				},
			},
			err: errors.New("no edge from parent to child"),
		},
		{
			scenario: "valid edge case 1",
			parentID: 1,
			childID:  2,
			ft: familyTree{
				nodes: map[int]*node{
					1: {
						id: 1,
						children: map[int]*node{
							2: {
								id: 2,
							},
						},
					},
					2: {
						id: 2,
						parents: map[int]*node{
							1: {
								id: 1,
							},
						},
						children: map[int]*node{
							3: {
								id: 3,
							},
						},
					},
					3: {
						id: 3,
						parents: map[int]*node{
							2: {
								id: 2,
							},
						},
					},
				},
			},
			err:            nil,
			remainingEdges: 1,
		},
		{
			scenario: "valid edge case 2",
			ft: familyTree{
				nodes: map[int]*node{
					1: {
						id: 1,
						children: map[int]*node{
							2: {
								id: 2,
							},
						},
					},
					2: {
						id: 2,
						parents: map[int]*node{
							1: {
								id: 1,
							},
						},
						children: map[int]*node{
							3: {
								id: 3,
							},
						},
					},
					3: {
						id: 3,
						parents: map[int]*node{
							2: {
								id: 2,
							},
						},
					},
				},
			},
			childID:        3,
			parentID:       2,
			err:            nil,
			remainingEdges: 1,
		},
	}
	for _, tc := range tests {
		err := tc.ft.DeleteEdge(tc.parentID, tc.childID)
		if tc.err != nil && err == nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		} else if tc.err == nil && err != nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		}

		if tc.err == nil && err == nil {
			edgeCount := tc.ft.GetEdgeCount()
			if edgeCount != tc.remainingEdges {
				t.Errorf("Scenario: %s \n got no of dependencies: %v, expected no of dependencies: %v", tc.scenario, edgeCount, tc.remainingEdges)
			}
		}
	}
}
