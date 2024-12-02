package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {

	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	dataString := string(data)

	taille := len(dataString)
	lignes := taille / 14
	listNumber1 := []int{}
	listNumber2 := []int{}

	for i := 0; i < lignes; i++ {
		number1 := string(dataString[14*i : 14*i+5])
		number2 := string(dataString[14*i+8 : 14*i+13])
		int1, _ := strconv.Atoi(number1)
		int2, _ := strconv.Atoi(number2)
		listNumber1 = append(listNumber1, int1)
		listNumber2 = append(listNumber2, int2)
	}

	sortedList1 := []int{}
	sortedList2 := []int{}
	step := 0
	listNumber1bis := listNumber1
	listNumber2bis := listNumber2

	for i := 0; i < len(listNumber1); i++ {
		min := findMin(listNumber1bis)
		sortedList1 = append(sortedList1, min)
		listNumber1bis = removeElement(listNumber1bis, min)

		min2 := findMin(listNumber2bis)
		sortedList2 = append(sortedList2, min2)
		listNumber2bis = removeElement(listNumber2bis, min2)
		step++
	}

	var somme float64

	for i, _ := range sortedList1 {
		distance := math.Abs(float64(sortedList1[i] - sortedList2[i]))
		somme += distance
	}

	similarity := 0
	for i := 0; i < len(sortedList1); i++ {
		appearance := 0
		for j := 0; j < len(sortedList2); j++ {
			if sortedList1[i] == sortedList2[j] {
				appearance++
			}
		}
		similarity += appearance * sortedList1[i]
	}

	fmt.Println(similarity)
}

func findMin(list []int) int {
	min := list[0]
	for _, j := range list {
		if j < min {
			min = j
		}
	}
	return min
}

func removeElement(list []int, element int) []int {
	for i, j := range list {
		if j == element {
			return append(list[:i], list[i+1:]...)
		}
	}
	return list
}
