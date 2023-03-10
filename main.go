package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// cli parse cli arguments
func cli() (num int) {
	flag.Usage = func() {
		w := flag.CommandLine.Output()
		fmt.Fprintf(w, "Usage:\n\t%s 100\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintf(w, "\n")
	}
	flag.Parse()
	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}
	num, err := strconv.Atoi(flag.Arg(0))
	if err != nil && len(os.Args) > 1 {
		fmt.Println(err)
		os.Exit(1)
	}
	return

}

// getswitches filter switches that stays on implementation of this problem https://www.youtube.com/watch?v=-UBDRX6bk-A
func getswitches(switches *[]int, num int) {
	for i := 1; i <= num; i++ {
		if multiple := i * i; multiple <= num {
			*switches = append(*switches, multiple)
		}
	}

}

// main func
func main() {
	var (
		num      int = cli()
		switches     = make([]int, 0)
	)
	getswitches(&switches, num)
	fmt.Println(strings.Trim(fmt.Sprint(switches), "[]"))
}
