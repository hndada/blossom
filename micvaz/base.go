package micvaz

import "errors"

const (
	Even = iota
	Odd
)

const (
	Outer = Even
	Inner = Odd
)

// Define the Node structure.
type Node struct {
	ID        int64
	Neighbors []*Node
	Levels    [2]int
	Matched   *Edge
	Blossom   *Blossom
	Label     int // Used to mark the node during the algorithm
}

func (n Node) State() int {
	if n.Levels[Even] < n.Levels[Odd] {
		return Even
	}
	return Odd
}
func (n Node) Level() int      { return n.Levels[n.State()] }
func (n Node) OtherLevel() int { return n.Levels[1-n.State()] }

// Edge represents an edge in the graph.
type Edge struct {
	From    *Node
	To      *Node
	Weight  float64
	IsMatch bool
}

// Define the Blossom structure.
type Blossom struct {
	Nodes  []*Node
	Base   *Node
	Parent *Blossom
	// Additional fields as required by the algorithm...
}

// Define the Graph structure.
type Graph struct {
	Nodes []*Node
	Edges []*Edge
	// Additional fields for algorithm state, like free nodes, etc.
}

// AddNode adds a node to the graph.
func (g *Graph) AddNode(id int64) *Node {
	node := &Node{ID: id}
	g.Nodes = append(g.Nodes, node)
	return node
}

// AddEdge adds an edge to the graph.
func (g *Graph) AddEdge(from, to *Node, weight float64) *Edge {
	edge := &Edge{From: from, To: to, Weight: weight}
	from.Children = append(from.Children, to)
	to.Children = append(to.Children, from)
	g.Edges = append(g.Edges, edge)
	return edge
}

// FindNodeByID finds a node by its ID.
func (g *Graph) FindNodeByID(id int64) (*Node, error) {
	for _, node := range g.Nodes {
		if node.ID == id {
			return node, nil
		}
	}
	return nil, errors.New("node not found")
}
