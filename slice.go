package main

import (
	"fmt"
	"sort"
)

func main() {

	var num int
	sli := make([]int, 0, 20)

	for {

		fmt.Print("Enter integer: ")
		_, err := fmt.Scan(&num)

		if err == nil {
			sli = append(sli, num)
			sort.Ints(sli)
			fmt.Println(sli)

		} else {
			break
		}
	}
}




// package main

// import (
// 	"fmt"
// 	"sort"
// 	"strconv"
// )

// func main() {
// 	userIntegers := make([]int, 0, 3)

// 	continueLoop := true
// 	for continueLoop {
// 		var userInput string
// 		fmt.Println("Please enter an integer or 'x' to exit.")
// 		fmt.Scan(&userInput)
// 		if userInput == "x" {
// 			continueLoop = false
// 		} else {
// 			userInteger, _ := strconv.Atoi(userInput)
// 			userIntegers = append(userIntegers, userInteger)
// 			sort.Ints(userIntegers)
// 			fmt.Println("Sorted entered integers: ", userIntegers)
// 		}
// 	}

// }
