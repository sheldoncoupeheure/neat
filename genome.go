/*


genome.go implementation of the genome in NEAT.

@licstart   The following is the entire license notice for
the Go code in this page.

Copyright (C) 2016 jin yeom, whitewolf.studio

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.

As additional permission under GNU GPL version 3 section 7, you
may distribute non-source (e.g., minimized or compacted) forms of
that code without the copy of the GNU GPL normally required by
section 4, provided you include this license notice and a URL
through which recipients can access the Corresponding Source.

@licend    The above is the entire license notice
for the Go code in this page.


*/

package neat

// Genome is an implementation of genotype of an evolving network;
// it includes NodeGenes and ConnGenes.
type Genome struct {
	numSensors int // number of sensor nodes
	numOutputs int // number of output nodes
	numHidden  int // number of hidden nodes

	numNodes int // total number of nodes
	numConns int // total number of connections

	nodes []*NodeGene // collection of node genes
	conns []*ConnGene // collection of connection genes
}

// NewGenome creates a new genome in its initial state, it is
// only consist of fully connected sensor nodes and output nodes.
func NewGenome(numSensors, numOutputs int) *Genome {
	// number of nodes and connections including bias
	numNodes := numSensors + 1 + numOutputs
	numConns := (numSensors + 1) * numOutputs

	nodes := make([]*NodeGene, 0, numNodes)
	conns := make([]*ConnGene, 0, numConns)
	// sensor nodes
	for i := 0; i < numSensors; i++ {
		nodes = append(nodes, NewNodeGene(i, "sensor", nil))
	}
	// output nodes and connections
	nodes = append(nodes, NewNodeGene(numNodes-1, "bias", nil))
	for i := numSensors + 1; i < numNodes; i++ {
		nodes = append(nodes, NewNodeGene(i, "output", nil))
		// connect from input nodes to this node
		for j := 0; j <= numSensors; j++ {
			conns = append(conns, NewConnGene(0, j, i))
		}
	}

	return &Genome{
		numSensors: numSensors,
		numOutputs: numOutputs,
		numHidden:  0,
		numNodes:   numNodes,
		numConns:   numConns,
		nodes:      nodes,
		conns:      conns,
	}
}

// NumSensors returns the number of sensor nodes in the genome.
func (g *Genome) NumSensors() int {
	return g.numSensors
}

// NumOutputs returns the number of output nodes in the genome.
func (g *Genome) NumOutputs() int {
	return g.numOutputs
}

// NumHidden returns the number of hidden nodes in the genome.
func (g *Genome) NumHidden() int {
	return g.numHidden
}

// NumNodes returns the total number of nodes in the genome.
func (g *Genome) NumNodes() int {
	return g.numNodes
}

// NumConns returns the total number of connections in the genome.
func (g *Genome) NumConns() int {
	return g.numConns
}

// Nodes returns all nodes in the genome.
func (g *Genome) Nodes() []*NodeGene {
	return g.nodes
}

// Conns returns all connections in the genome.
func (g *Genome) Conns() []*ConnGene {
	return g.conns
}

// NodeGene is an implementation of each node within a genome.
// Each node includes a node ID (NID), a node type (NType), and
// a pointer to an activation function.
type NodeGene struct {
	nid   int             // node ID
	ntype string          // node type
	afn   *ActivationFunc // activation function
}

// NewNodeGene creates a new node gene with the given NID, node type, and
// a pointer to an activation function.
func NewNodeGene(nid int, ntype string, afn *ActivationFunc) *NodeGene {
	return &NodeGene{
		nid:   nid,
		ntype: ntype,
		afn:   afn,
	}
}

// NID returns the node's node ID (NID).
func (n *NodeGene) NID() int {
	return n.nid
}

// NType returns the node's node type (NType).
func (n *NodeGene) NType() string {
	return n.ntype
}

// Afn returns the node's activation function.
func (n *NodeGene) Afn() *ActivationFunc {
	return n.afn
}

// ConnGene is an implementation of each connection within a genome.
// It represents a connection between an in-node and an out-node;
// it contains an innovation number and nids of the in-node and the
// out-node.
type ConnGene struct {
	innov int // innovation number
	in    int // NID of in-node
	out   int // NID of out-node
}

// NewConnGene creates a new connection gene with the given innovation
// number, the in-node NID, and the out-node NID.
func NewConnGene(innov, in, out int) *ConnGene {
	return &ConnGene{
		innov: innov,
		in:    in,
		out:   out,
	}
}

// Innov returns the connection's innovation number.
func (c *ConnGene) Innov() int {
	return c.innov
}

// In returns the NID of in-node of the connection.
func (c *ConnGene) In() int {
	return c.in
}

// Out returns the NID of out-node of the connection.
func (c *ConnGene) Out() int {
	return c.out
}