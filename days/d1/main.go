package d1

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {
	fmt.Print("day 1: increased count ")
	i := getInput()
	r := parseIncreased(i)
	fmt.Println(r)
}

func getInput() []int {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	dat, err := os.ReadFile(dir + "/days/d1/input")
	if err != nil {
		panic(err)
	}

	s := strings.Split(string(dat), "\n")
	n := []int{}
	for _, s := range s {
		i, _ := strconv.Atoi(s)
		n = append(n, i)
	}

	return n
}

func parseIncreased(measurements []int) int {
	increased := 0
	for i := range measurements {
		if i == 0 {
			continue
		}
		if measurements[i] > measurements[i-1] {
			increased += 1
		}
	}
	return increased
}
