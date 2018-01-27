package main

import (
	"flag"
	"fmt"
)

type Selection struct {
	top    int
	height int
	left   int
	width  int
}

func main() {
	sele := Selection{}
	flag.IntVar(&sele.top, "t", 1, "top line number of selection")
	flag.IntVar(&sele.height, "h", 1, "height of line of selection")
	flag.IntVar(&sele.left, "l", 0, "column number of selection's lefside")
	flag.IntVar(&sele.width, "w", 1, "number of columns of selection")

	flag.Parse()

	fmt.Printf("%v", sele)
}
