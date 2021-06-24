// package main

// import (
// 	"fmt"
// 	"time"
// )

// func GenDisplaceFn(a float64, vo float64, so float64) func (float64) float64 {

// 	fn := func (t float64) float64 {
// 		return so + (vo*t) + (0.5*a*t*t)
// 	}
// 	return fn

// }


// func main(){
// 	fmt.Println("Enter Acceleration: ")
// 	var a float64
// 	fmt.Scanln(&a)

// 	fmt.Println("Enter Initial Velocity: ")
// 	var vo float64
// 	fmt.Scanln(&vo)

// 	fmt.Println("Enter Initial Displacement: ")
// 	var so float64
// 	fmt.Scanln(&so)

// 	fmt.Println("Enter Time: ")
// 	var t float64
// 	fmt.Scanln(&t)

// 	gen_disp := GenDisplaceFn(a, vo, so)
// 	disp := gen_disp(t)

// 	time.Sleep(2 * time.Second)
	
// 	fmt.Println(disp)
// }





package main

import "fmt"

func main() {
	var A, V0, S0 float64
	var t float64
	fmt.Printf("Enter value of Acceleration(a) \t\t\t:=>")
	fmt.Scanln(&A)

	fmt.Printf("Enter value of Initial Velocity(V0) \t\t:=>")
	fmt.Scanln(&V0)

	fmt.Printf("Enter value of Initial Displacement(S0) \t:=>")
	fmt.Scanln(&S0)

	fmt.Printf("Enter value of Time(t) \t\t\t\t:=>")
	fmt.Scanln(&t)
	fmt.Println("------------------------------------------------------")

	fn := GenDisplaceFn(A, V0, S0)
	fmt.Printf("Displacement after %v Seconds \t\t\t:=>%v\n", t, fn(t))
	fmt.Println("------------------------------------------------------")

	//fmt.Println("Displacement after 5 Seconds: \t", fn(5))
}

func GenDisplaceFn(acc, iniVelocity, iniDisplacement float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*acc*(t*t) + iniVelocity*t + iniDisplacement
	}
}