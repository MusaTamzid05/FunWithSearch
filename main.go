package main

import (
    "fun_with_search/lib"
    "fmt"
)

type SizeData struct {
    MinSize int
    MaxSize int
}

func main() {
    nodeReader := lib.MakeCSVNodeReader("/home/musa/datasets/amazon_dataset/amz_ca_total_products_data_processed.csv")
    nodes, err := nodeReader.Read()

    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("Total nodes ", len(nodes))

    sizeDataList := []SizeData {
        SizeData{MinSize:0, MaxSize:500},
        SizeData{MinSize:500, MaxSize:1000},
        SizeData{MinSize:1000, MaxSize:1000000},
    }


    var selectors []lib.Selector
    selectors = make([]lib.Selector, len(sizeDataList))

    for index, sizeData := range sizeDataList {
        selectors[index] =  lib.MakeSelectorByTextSize(sizeData.MinSize,
                                                     sizeData.MaxSize, 
                                                     "title")
    }



    nodeToIndexGenerator := lib.MakeNodeToIndexGenerator(selectors)
    nodeToIndexGenerator.Generate(nodes)

    fmt.Println("total index ", len(nodeToIndexGenerator.IndexList))

    for i, indexData := range nodeToIndexGenerator.IndexList {
        fmt.Println(i , " => ", len(indexData.Nodes))
    }




}
