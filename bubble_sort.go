// package main

// import (
// 	// "bufio"
// 	"fmt"
// 	// "os"
// 	// "strings"
// )

// func Swap(slice []int, j int){
// 	var temp int

// 	temp = slice[j]
// 	slice[j] = slice[j+1]
// 	slice[j+1] = temp

// }

// func BubbleSort(slice []int){

// 	n := len(slice)

// 	for i := 0; i < n-1; i++ {
// 		for j := 0; j < n-i-1; j++ {
// 			if slice[j] > slice[j+1]{
// 				Swap(slice, j)
// 			}
// 		}
// 	}
// }



// func main(){
// 	slice := make([]int, 0, 10)
// 	var num int
// 	for x:= 0; x<10; x++ {
// 		fmt.Scanln(&num)
// 		slice = append(slice, num)
// 	}

// 	BubbleSort(slice)
// 	fmt.Println(slice)
// }















package main

import (
	"fmt"
	"strconv"
	"strings"
)

func IntsToStrings(numbers []int) []string {
	result := make([]string, len(numbers))

	for i, number := range numbers {
		result[i] = strconv.Itoa(number)
	}

	return result
}

func Swap(numbers []int, i int) {
	temp := numbers[i]
	numbers[i] = numbers[i+1]
	numbers[i+1] = temp
}

func BubbleSort(numbers []int) []int {
	n := len(numbers)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				Swap(numbers, j)
			}
		}
	}

	return numbers
}

func main() {
	var number string
	var countNumbers string

	fmt.Scanln(&countNumbers)
	countNumbersInt, _ := strconv.Atoi(countNumbers)
	numbers := make([]int, countNumbersInt)

	for i := 0; i < countNumbersInt; i++ {
		fmt.Scanln(&number)

		numberInt, _ := strconv.Atoi(number)
		numbers[i] = numberInt
	}

	sorted := BubbleSort(numbers)
	sortedStrings := IntsToStrings(sorted)
	result := strings.Join(sortedStrings, ", ")

	fmt.Println(result)
}