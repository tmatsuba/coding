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
	K := nextInt()
	var rarr []int
	for i := 0; i < K; i ++ {
		rarr = append(rarr, nextInt())
	}
	var arr [][]int
	for i := 0; i < N; i ++ {
		var b []int
		b = append(b, nextInt())
		b = append(b, nextInt())
		arr = append(arr, b)
	}

	max := 0.0
	for _, a := range(arr) {
		min1 := math.MaxFloat64
		for _, r := range(rarr) {
			rbase := arr[r-1]
		 	v := math.Sqrt(math.Pow(float64(a[0]-rbase[0]),2) + math.Pow(float64(a[1]-rbase[1]),2))			
			if v < min1 {
				min1 = v
			}
		}
		if max < min1 {
			max = min1
		}
	}
	fmt.Println(max)
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
