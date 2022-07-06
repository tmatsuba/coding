package main

import (
	"os"
	"bufio"
	"fmt"
	"strconv"
	"math"
)

func main() {
	sc.Split(bufio.ScanWords)
	N := nextInt()
	M := nextInt()
	slist := []uint{}
	ret := 0

	for i := 0; i < M; i++ {
		k := nextInt()
		var s uint = 0
		for j := 0; j < k; j++ {
			s += uint(pow(2, nextInt() - 1))
		}
		slist = append(slist, s)
	}

	p := []int{}
	for l := 0; l < M; l++ {
		p = append(p, nextInt())
	}

	for bits:=0;bits<(1<<uint64(N));bits++ {
		on := true
		for i, s := range(slist) {
			sum := 0
			for j:=0;j<N;j++ {

				if ((uint(bits) & s)>>uint64(j))&1 == 1 {
					sum += 1
				}
			}
			if sum % 2 != p[i] {
				on = false
				break
			}
		}
		if on {
			ret++
		}

	}

	fmt.Println(ret)
}


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

func atoi(s string) int {
	ret, e := strconv.Atoi(s)
	if e != nil {
		panic(e)
	}
	return ret
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}
