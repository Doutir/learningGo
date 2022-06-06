package main

import (
	"fmt"
	"strconv"
)

type evenAndOdd []int

func generateNumbers(quantity int) evenAndOdd {
	newNumbers := evenAndOdd{}

	for i := 0; i < quantity; i++ {
		fmt.Println(i)
		newNumbers = append(newNumbers, i+1)
	}

	return newNumbers
}

func (eo evenAndOdd) printIsOddOrEven() {
	for index, number := range eo {
		if number%2 == 0 {
			fmt.Println(strconv.Itoa(eo[index]) + " - is even")
		} else {
			fmt.Println(strconv.Itoa(eo[index]) + " - is odd")
		}
	}
}
