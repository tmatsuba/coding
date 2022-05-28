package main

import (
	"os"
	"bufio"
	"fmt"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
    sc.Scan()
    return sc.Text()
}

func nextInt() int {
    sc.Scan()
    i, e := strconv.Atoi(sc.Text())
    if e != nil {
        panic(e)
    }
    return i
}

func main() {
	sc.Split(bufio.ScanWords)
	a := nextInt()
	b := nextInt()

	if a * b % 2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
	
}
