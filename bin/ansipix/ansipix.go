package main

import (
	"ansipix"
	"fmt"
	_ "image/gif"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Use: ansipix filename [columns int]\neg. ansipix mario.jpg 80\ncolumns is optional, default 80.")
		return
	}

	w := 80
	var err error
	if len(os.Args) > 2 {
		w, err = strconv.Atoi(os.Args[2])
		if err != nil {
			panic(err)
		}
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer f.Close()

	str, err := ansipix.Blockify(f, w)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", str)
}
