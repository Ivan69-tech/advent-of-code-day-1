package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	dataString := string(data)
	fmt.Println(dataString)
	taille := len(dataString)
	lignes := taille / 14
	listNumber1 := []int{}
	listNumber2 := []int{}

	for i := 0; i < lignes; i++ {
		number1 := ""
		number2 := ""
		for j := 0; j < 14; j++ {
			if j == 0|1|2|3|4 {
				number1 += string(dataString[i+j])
				fmt.Println("number1")
			} else if j == 8|9|10|11|12 {
				number2 += string(dataString[i+j])
			}
		}
		int1, _ := strconv.Atoi(number1)
		int2, _ := strconv.Atoi(number2)
		listNumber1 = append(listNumber1, int1)
		listNumber2 = append(listNumber2, int2)
	}

	fmt.Println(listNumber1)
}
