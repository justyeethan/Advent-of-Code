package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to advent of code!")
	f, err := os.Open("./day1PuzzleInput.txt")
	if err != nil {
		log.Fatal()
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	list1 := []int{}
	list2 := []int{}
	for scanner.Scan() {
		splitLine := strings.Split(scanner.Text(), "   ")
		val1, err := strconv.Atoi(splitLine[0])
		if err != nil {
			panic(err)
		}
		list1 = append(list1, val1)
		val2, err := strconv.Atoi(splitLine[1])
		if err != nil {
			panic(err)
		}
		list2 = append(list2, val2)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	// fmt.Println(listComparison(list1, list2)) // Part 1
	fmt.Println(similarityScore(list1, list2)) // Part 2
}

// Part 1
func listComparison(list1 []int, list2 []int) int64 {
	sort.Slice(list1, func(i, j int) bool {
		return list1[i] > list1[j]
	})
	sort.Slice(list2, func(i, j int) bool {
		return list2[i] > list2[j]
	})
	differences := []int{}
	for i := range list1 {
		var res float64 = float64(list2[i] - list1[i])
		differences = append(differences, int(math.Abs(res)))
	}
	var res int64
	for i := range differences {
		res += int64(differences[i])
	}
	return res
}

// Part 2
func similarityScore(list1 []int, list2 []int) int64 {
	var res int64
	seenInList2 := map[int]int{}
	for _, v := range list2 {
		fmt.Println(v)
		if _, ok := seenInList2[v]; ok {
			fmt.Printf("Value: %v", v)
			seenInList2[v] += 1
		} else {
			seenInList2[v] = 1
		}
	}
	fmt.Println(seenInList2)
	for _, list1Value := range list1 {
		if _, ok := seenInList2[list1Value]; ok {
			res += int64(list1Value) * int64(seenInList2[list1Value])
		}
	}
	return res
}
