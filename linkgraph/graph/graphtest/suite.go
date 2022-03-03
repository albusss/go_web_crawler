package graphtest

import "github.com/albus/srch/linkgraph/graph"

// SuiteBase defines a re-usable set of graph-related tests that can
// be executed against any type that implements graph.Graph.
type SuiteBase struct {
	g graph.Graph
}

// SetGraph configures the test-suite to run all tests against g.
func (s *SuiteBase) SetGraph(g graph.Graph) {
	s.g = g
}
