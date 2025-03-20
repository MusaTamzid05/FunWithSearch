package lib


type NodeToIndexGenerator struct {
    selectors []Selector
    IndexList []Index
}

func MakeNodeToIndexGenerator(selectors []Selector) NodeToIndexGenerator {
    return NodeToIndexGenerator{selectors: selectors}
}


func (n *NodeToIndexGenerator) Generate(nodes []Node)  {
    n.IndexList = make([]Index, len(n.selectors))

    for _, node := range nodes {
        for index, selector := range n.selectors {
            if selector.IsSelected(node) {
                n.IndexList[index].Add(node)
            }

        }

    }

}
