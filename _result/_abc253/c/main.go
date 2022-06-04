package main

import (
	"os"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	Q := nextInt()
    m := map[string]int{}
	maxv := -1
	maxc := 0
	minv := -1
	minc := 0

	for i := 0; i < Q; i++ {
		q := nextLine()
		if strings.HasPrefix(q, "1") {
			x := strings.Split(q, " ")[1]
			_, ok := m[x]
			if ok {
				m[x]++
			} else {
				m[x] = 1
			}

			if maxv == -1 {
				maxv = atoi(x)
				maxc = 1
			}
			if minv == -1 {
				minv = atoi(x)
				minc = 1
			}
			
			if maxv < atoi(x) {
				maxv = atoi(x)
				maxc = 1
			} else if minv > atoi(x) {
				minv = atoi(x)
				minc = 1
			} else if maxv == atoi(x) {
				maxc++
				if minv == atoi(x) {
					minc++
				}
			} else if minv == atoi(x) {
				minc++
				if maxv == atoi(x) {
					maxv++
				}
			}

		}
		if strings.HasPrefix(q, "2") {
			x := strings.Split(q, " ")[1]
			c := atoi(strings.Split(q, " ")[2])

			_, ok := m[x]
			if ok {
				if m[x] <= c {
					delete(m, x)
					if atoi(x) == maxv {
						maxv = getMax(m)
						maxc = m[x]
					} 
					if atoi(x) == minv {
						minv = getMin(m)
						minc = m[x]
					} 
				} else {
					m[x] = m[x] - c
				}
			}
		}
//		fmt.Println(m)
		if strings.HasPrefix(q, "3") {
			fmt.Println(maxv - minv)
		}
	}
}

func getMax(m map[string]int) int {
	if len(m) == 0 {
		return -1
	}

	max := 0
	for k := range m {
		intk, e := strconv.Atoi(k)
		if e != nil {
			panic(e)
		}
		max = Max(max, intk)
	}
	return max
}

func getMin(m map[string]int) int {
	if len(m) == 0 {
		return -1
	}

	min := 1000000000
	for k := range m {
		intk, e := strconv.Atoi(k)
		if e != nil {
			panic(e)
		}
		min = Min(min, intk)
	}
	return min
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
