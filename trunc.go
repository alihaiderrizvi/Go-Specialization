package main

import ("fmt")

func main(){
	fmt.Println("Enter Floating Number: ")
	var float float32
	fmt.Scanln(&float)
	var i int = int(float)
	fmt.Println("Truncated Number: ")
	fmt.Println(i)
}



// package main

// import "fmt"

// func main() {
// 	fmt.Print("Please enter a floating point number: ")

// 	var num float64

// 	_, err := fmt.Scan(&num)
// 	if err != nil {
// 		fmt.Println("sorry you didnt enter a valid floating point number.")
// 	}

// 	fmt.Printf("truncated: %d\n", int(num))
// }