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
	n := nextInt()
	a := make([]int, n)
	i := 0

	for n > i {
		a[i] = nextInt()
		i++
	}

	ret := 0
	for {
		i = 0
		odd := false
		for n > i {
			if a[i] % 2 != 0 {
				odd = true
				break
			}
			a[i] = a[i] / 2
			i++
		}
		if odd {
			break
		}
		ret++
	}
	fmt.Println(ret)
}
