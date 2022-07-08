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
	D := nextInt()
	X := [][]int{}
	ret := 0

	for i := 0; i < N; i++ {
		l := []int{}
		for j := 0; j < D; j++ {
			l = append(l, nextInt())
		}
		X = append(X, l)
	}

	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			var sum float64 = 0
			for k := 0; k < D; k++ {
				sum += float64(pow((X[i][k]-X[j][k]), 2))
			}
			judge := math.Sqrt(sum)

			if math.Floor(judge) -  judge == 0 {
				ret++
			}
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
