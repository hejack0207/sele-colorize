package main

import (
	"flag"
	"fmt"
	"strings"
)

type Selection struct {
	start Pos
	end   Pos
}

type Pos struct {
	line   int
	column int
}

func main() {
	var (
		start    string
		end      string
		offset   int
		filepath string
	)
	flag.StringVar(&start, "start", "1,0", "start position of selection")
	flag.StringVar(&end, "end", "1,0", "end position of selection")
	flag.IntVar(&offset, "offset", 0, "offset of bytes in file")
	flag.StringVar(&filepath, "f", "", "file to print")

	flag.Parse()
	strings.Split(start, ",")
	//fmt.Printf("start:%s,end:%s", start, end)
}
