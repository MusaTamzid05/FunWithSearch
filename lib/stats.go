package lib

import (
    "fmt"
)

type Stats interface {
    Show(nodes []Node)
}

type TextStats struct {
    Key string
}

func (t TextStats) Show(nodes []Node) {
    maxSize := -1
    maxSizeIndex := -1

    minSize := 10000000
    minSizeIndex := -1

    total := 0

    for index, node := range nodes {
        value := node.Data[t.Key]
        length := len(value)

        if length > maxSize {
            maxSize = length
            maxSizeIndex = index
        }


        if length <  minSize {
            minSize = length
            minSizeIndex = index
        }

        total += length
    }

    fmt.Println("Info for attribute ", t.Key)
    fmt.Println("Min Size ",  minSize)
    fmt.Println("Min Size Index ",  minSizeIndex)
    fmt.Println("Min Text ",  nodes[minSizeIndex].Data[t.Key])

    fmt.Println("Max Size ",  maxSize)
    fmt.Println("Max Size Index ",  maxSizeIndex)
    fmt.Println("Min Text ",  nodes[maxSizeIndex].Data[t.Key])


    fmt.Println("Average length ",  total / len(nodes))

}


func MakeTextStats(key string) TextStats {
    return TextStats{Key: key}
}
