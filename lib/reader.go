package lib

import (
    "os"
    "encoding/csv"
)

type NodeReader interface {
    Read() ([]Node, error)
}

type CSVNodeReader struct {
    Path string

}

func (r CSVNodeReader) Read()([]Node, error) {
    nodes := []Node{}

    fp, err := os.Open(r.Path)

    if err != nil {
        return nodes, err
    }

    defer fp.Close()
    headers := []string{}

    reader := csv.NewReader(fp)
    records, err := reader.ReadAll()

    if err != nil {
        return nodes, err
    }

    firstLine := true

    for _, record := range records {

        if firstLine {
            for _, item := range record {
                headers = append(headers, item)
            }
            firstLine = false
            continue
        }

        values := []string{}


        for _, item := range record {
            values = append(values, item)
        }

        node := MakeNode(headers, values)
        nodes = append(nodes, node)

    }


    return nodes, nil
}

func MakeCSVNodeReader(path string) CSVNodeReader {
    return CSVNodeReader{Path: path}
}




