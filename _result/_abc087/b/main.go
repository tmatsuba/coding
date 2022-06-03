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
	A := nextInt()
	B := nextInt()
	C := nextInt()
	X := nextInt()
	ret := 0
	for a := 0; a <= A; a++ {
		for b := 0; b <= B; b++ {
			for c := 0; c <= C; c++ {
				if (a * 500 + b * 100 + c * 50  == X) {
					ret++
				}
				
			}
		}
	}

	fmt.Println(ret)
}


func Max(x, y int) int {
    if x < y {
        return y
    }
    return x
}

func Min(x, y int) int {
    if x > y {
        return y
    }
    return x
}
