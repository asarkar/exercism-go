package tree

import (
	"errors"
	"fmt"
	"slices"
)

type Record struct {
	ID     int
	Parent int
	// feel free to add fields as you see fit
}

type Node struct {
	ID       int
	Children []*Node
	// feel free to add fields as you see fit
}

func Build(records []Record) (*Node, error) {
	//nolint:nilnil // ok to return both values as nil
	if len(records) == 0 {
		return nil, nil
	}

	adjList, err := makeAdjList(records)
	if err != nil {
		return nil, err
	}

	explored := make(map[int]bool)
	return dfs(0, adjList, explored)
}

func makeAdjList(records []Record) (map[int][]int, error) {
	adjList := make(map[int][]int)
	numRecords := len(records)
	// Tests expect the children to be sorted.
	slices.SortFunc(records, func(r1, r2 Record) int {
		if r1.ID < r2.ID {
			return -1
		}
		if r1.ID > r2.ID {
			return 1
		}
		return 0
	})

	for i := range numRecords {
		record := records[i]
		if err := validate(record, numRecords); err != nil {
			return nil, err
		}
		if i > 0 && records[i-1].ID == record.ID {
			return nil, fmt.Errorf("duplicate node [%d]", record.ID)
		}
		if _, exists := adjList[record.ID]; !exists {
			adjList[record.ID] = make([]int, 0)
		}
		// root should not be its own child.
		if record.ID > 0 {
			adjList[record.Parent] = append(adjList[record.Parent], record.ID)
		}
	}
	if len(adjList) != len(records) {
		return nil, errors.New("missing record(s)")
	}
	return adjList, nil
}

func validate(record Record, numRecords int) error {
	if record.ID < 0 || record.ID >= numRecords {
		return fmt.Errorf("ID out of range [%d] with length %d", record.ID, numRecords)
	}
	if record.Parent < 0 || record.Parent >= numRecords {
		return fmt.Errorf("parent ID out of range [%d] with length %d", record.Parent, numRecords)
	}
	if record.ID == 0 && record.Parent > 0 {
		return fmt.Errorf("root with non-zero parent ID %d", record.Parent)
	}
	if record.ID != 0 && record.ID <= record.Parent {
		return fmt.Errorf("ID [%d] less than or equal to parent ID %d", record.ID, record.Parent)
	}
	return nil
}

func dfs(id int, adjList map[int][]int, explored map[int]bool) (*Node, error) {
	// explored[id] states:
	//   absent -> unvisited
	//   false  -> visiting
	//   true   -> done

	if val, exists := explored[id]; exists && !val {
		return nil, fmt.Errorf("cycle detected at node [%d]", id)
	}

	explored[id] = false // mark visiting

	var children []*Node
	for _, childID := range adjList[id] {
		child, err := dfs(childID, adjList, explored)
		if err != nil {
			return nil, err
		}
		children = append(children, child)
	}

	explored[id] = true // mark done
	return &Node{ID: id, Children: children}, nil
}
