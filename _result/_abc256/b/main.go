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
	P := 0
	m := []int{1, 0, 0, 0}


	for i := 0; i < N; i++ {
		a := nextInt()
//		fmt.Println("m: ", m, "a: ", a)
//		fmt.Println("m[len(m)-a:]: ", m[len(m)-a:], "m[:len(m)-a]: ", m[:len(m)-a])
		for _, v := range(m[len(m)-a:]) {
			P = P + v
		}

		t := []int{0, 0, 0, 0}
		for j, v := range(m){
//			fmt.Println("a: ", a, "j :", j, m)
			if v == 1 && j+a < len(m){
				t[j+a] = 1
			}
		}
//		fmt.Println("t: ", t, "P: ", P)
		m = t
		m[0] = 1
	}

	fmt.Println(P)
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
