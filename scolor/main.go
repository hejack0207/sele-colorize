package main

import (
	"flag"
	"fmt"
	//"github.com/fatih/color"
	"os"
	"strconv"
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
	flag.StringVar(&start, "start", "1,1", "start position of selection")
	flag.StringVar(&end, "end", "1,1", "end position of selection")
	flag.IntVar(&offset, "offset", 1, "offset of bytes in file")
	flag.StringVar(&pos, "pos", "1,1", "position in input file")
	flag.StringVar(&filepath, "f", "", "file to print")

	flag.Parse()
	lc := strings.Split(pos, ",")
	fmt.Printf("line:%s,column:%s\n", lc[0], lc[1])

	l, _ := strconv.Atoi(lc[0])
	c, _ := strconv.Atoi(lc[1])
	ppos := Pos{l, c}
	f, e := os.Open(filepath)
	defer f.Close()
	if e != nil {
		fmt.Fprintf(os.Stderr, "%s\n", e)
		os.Exit(1)
	}
	fi, e := f.Stat()
	if e != nil {
		fmt.Fprintf(os.Stderr, "%s\n", e)
		os.Exit(1)
	}
	fc := make([]byte, fi.Size())

	n, err := f.Read(fc)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	if int64(n) != fi.Size() {
		fmt.Fprintf(os.Stderr, "Only part of file is read")
		os.Exit(1)
	}
	sfc := string(fc)
	cPos := Pos{1, 1}
	fmt.Printf("%s\n", sfc)
	for i, c := range sfc {
		if c == '\n' || c == '\r' {
			cPos.line += 1
			cPos.column = 1
		} else {
			cPos.column += 1
		}

		if (cPos.line == ppos.line && cPos.column > ppos.column) ||
			(cPos.line > ppos.line) {
			fmt.Println("Selected:" + sfc[:i])
			break
		}
	}
}
