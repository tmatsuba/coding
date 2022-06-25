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
	h1 := nextInt()
	h2 := nextInt()
	h3 := nextInt()
	w1 := nextInt()
	w2 := nextInt()
	w3 := nextInt()
	ret := 0

	for a := 1 ; a <= Min(h1-2, w1-2); a++ {
		for b := 1 ; b <= Min(h1-a-1,w2-2); b++ {
			for c := 1; c <= Min(h1-a-b,w3-2); c++ {
				if a+b+c != h1 {continue}
				for d := 1; d <= Min(h2-2, w1-a-1); d++ {
					for g :=1; g <= Min(h3-2,w1-a-d); g++ {
						if a+d+g != w1 {continue}
						for e :=1; e <= Min(h2-d-1,w2-b-1); e++ {
							for h :=1; h <= Min(h3-g-1,w2-b-e); h++ {
								if b+e+h != w2 {continue}
								for f :=1; f <= Min(h2-d-e,w3-c-1); f++ {
									if d+e+f != h2 {continue}
									for i :=1; i <= Min(h3-g-h,w3-c-f); i++ {
										if a+b+c == h1 && d+e+f == h2 && g+h+i == h3 && a+d+g == w1 && b+e+h == w2 && c+f+i == w3 {ret++}
										}
								}
							}
						}
					}
				}
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
