package lib

type Index struct {
    Nodes []Node
}

func (i *Index) Add(node Node) {
    i.Nodes = append(i.Nodes, node)
}
