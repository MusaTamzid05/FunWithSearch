package main

import (
    "fun_with_search/lib"
    "fmt"
)

type SizeData struct {
    MinSize int
    MaxSize int
}

// @TODO : Make all the nodes a pointer type
// otherwises we are just making copies !!

func main() {
    nodeReader := lib.MakeCSVNodeReader("/home/musa/datasets/amazon_dataset/amz_ca_total_products_data_processed.csv")
    nodes, err := nodeReader.Read()

    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("Total nodes ", len(nodes))
    stats := lib.MakeTextStats("title")
    stats.Show(nodes)

    sizeDataList := []SizeData {
        SizeData{MinSize:0, MaxSize:25},
        SizeData{MinSize:25, MaxSize:80},
        SizeData{MinSize:80, MaxSize:100},
        SizeData{MinSize:100, MaxSize:150},
        SizeData{MinSize:150, MaxSize:200},
        SizeData{MinSize:200, MaxSize:300},
        SizeData{MinSize:300, MaxSize:500},
        SizeData{MinSize:500, MaxSize:1000},
        SizeData{MinSize:1000, MaxSize:2000},
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
