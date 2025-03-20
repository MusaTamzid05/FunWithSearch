package lib

import (
    "fmt"
)

type Node struct {
    Data map[string]string
}

func (n Node) Show() {
    for key, value := range n.Data {
        fmt.Println(key, " => ", value)

    }

}


func MakeNode(keys, values []string) Node {
    dataMap := make(map[string]string)

    for index, key := range keys {
        dataMap[key] = values[index]
    }

    return Node{Data: dataMap}
}
