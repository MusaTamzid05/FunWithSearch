package main

import (
    "fun_with_search/lib"
    "fmt"
)

func main() {
    nodeReader := lib.MakeCSVNodeReader("/home/musa/datasets/amazon_dataset/amz_ca_total_products_data_processed.csv")
    nodes, err := nodeReader.Read()

    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("Total nodes ", len(nodes))

    for i := 0; i < 3; i += 1 {
        nodes[i].Show()
        fmt.Println("---")
    }
}
