package main

import (
	"os"
	"bufio"
	"fmt"
	"strconv"
	"strings"	
	"math"
)

func main() {
	l := strings.Split(nextLine(), " ")
	H ,e:=  strconv.Atoi(l[0])
	if e!= nil {
		panic(e)
	}

	h1 := -1
	h2 := -1
	w1 := -1 
	w2 := -1
	
	for i := 0; i < H; i++ {
		str := nextLine()
		if strings.Index(str, "o") != -1 {
			if h1 == -1 {
				h1 = i
				w1 = strings.Index(str, "o")
				if strings.Index(str[w1 + 1:], "o") != -1 {
					h2 = i
					w2 = w1 + strings.Index(str[w1 + 1:], "o") + 1
					break
				}
			} else {
				h2 = i
				w2 = strings.Index(str, "o")
			}
		}
	}
	// fmt.Println(h1)
	// fmt.Println(w1)
	// fmt.Println(h2)
	// fmt.Println(w2)

	fmt.Println(math.Abs(float64(h1) - float64(h2)) + math.Abs(float64(w2) - float64(w1)))
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
