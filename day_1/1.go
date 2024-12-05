package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var left_slice []int
	var right_slice []int
	for scanner.Scan() {
		var line = strings.Split(scanner.Text(), "   ")
		var left, _ = strconv.Atoi(line[0])
		var right, _ = strconv.Atoi(line[1])

		left_slice = append(left_slice, left)
		right_slice = append(right_slice, right)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	slices.Sort(left_slice)
	slices.Sort(right_slice)
	var res int
	for i, _ := range left_slice {
		res += diff(left_slice[i], right_slice[i])
	}
	fmt.Println(res)
}

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
