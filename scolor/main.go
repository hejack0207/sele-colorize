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
	sele.top = *flag.Int("t", 1, "top line number of selection")
	sele.height = *(flag.Int("h", 1, "height of line of selection"))
	sele.left = *flag.Int("l", 0, "column number of selection's lefside")
	sele.width = *flag.Int("w", 1, "number of columns of selection")
	flag.Parse()

	fmt.Printf("%v", sele)
}
