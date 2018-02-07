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
		pos      string
		offset   int
		filepath string
	)
	flag.StringVar(&start, "start", "1,0", "start position of selection")
	flag.StringVar(&end, "end", "1,0", "end position of selection")
	flag.IntVar(&offset, "offset", 0, "offset of bytes in file")
	flag.StringVar(&pos, "pos", "1,0", "position in input file")
	flag.StringVar(&filepath, "f", "", "file to print")

	flag.Parse()
	lc := strings.Split(pos, ",")
	fmt.Printf("line:%s,column:%s", lc[0], lc[1])
}
