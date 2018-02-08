package main

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
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
	flag.StringVar(&start, "start", "1,0", "start position of selection")
	flag.StringVar(&end, "end", "1,0", "end position of selection")
	flag.IntVar(&offset, "offset", 0, "offset of bytes in file")
	flag.StringVar(&pos, "pos", "1,0", "position in input file")
	flag.StringVar(&filepath, "f", "", "file to print")

	flag.Parse()
	lc := strings.Split(pos, ",")
	fmt.Printf("line:%s,column:%s\n", lc[0], lc[1])

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
	fcstr := fmt.Sprintf("%s", fc)
	//fmt.Println(fcstr)
	lines := strings.Split(fcstr, "\n")
	c := color.New(color.FgGreen, color.BgWhite, color.Underline)
	line_no, _ := strconv.Atoi(lc[0])
	if line_no < 1 || line_no > len(lines) {
		fmt.Print("line:" + lc[0] + " no do not exist")
	} else {
		c.Println("matched:")
		c.Printf("%s\n", lines[line_no-1])
	}
}
