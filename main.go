package main

import (
	"flag"
	"fmt"
	"strconv"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 3 {
		fmt.Println("need 3 arguments")
		return
	}
	from, err := strconv.ParseInt(args[1], 10, 0)
	if err != nil {
		fmt.Println(err)
	}

	target, err := strconv.ParseInt(args[0], int(from), 0)
	if err != nil {
		fmt.Println(err)
	}

	to, err := strconv.ParseInt(args[2], 10, 0)
	if err != nil {
		fmt.Println(err)
	}

	switch to {
	case 2:
		fmt.Printf("%b\n", target)
	case 8:
		fmt.Printf("%o\n", target)
	case 10:
		fmt.Printf("%d\n", target)
	case 16:
		fmt.Printf("%x\n", target)
	default:
		fmt.Println("can not convert")
	}
}
