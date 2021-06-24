package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) eat() {
	fmt.Println(a.food)
}

func (a Animal) move() {
	fmt.Println(a.locomotion)
}

func (a Animal) speak() {
	fmt.Println(a.noise)
}

func main() {

	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}

	for {

		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf("> ")

		if scanner.Scan() {
			data := strings.Fields(scanner.Text())
			
			if len(data) != 2{
				fmt.Println("Invalid request")
				continue
			}

			name := data[0]
			meth := data[1]

			if name == "cow" {
				if meth == "eat" {
					cow.eat()
				} else if meth == "move" {
					cow.move()
				} else if meth == "speak" {
					cow.speak()
				} else{fmt.Println("Invalid Request")}

			} else if name == "snake" {
				if meth == "eat" {
					snake.eat()
				} else if meth == "move" {
					snake.move()
				} else if meth == "speak" {
					snake.speak()
				} else{fmt.Println("Invalid Request")}

			} else if name == "bird" {
				if meth == "eat" {
					bird.eat()
				} else if meth == "move" {
					bird.move()
				} else if meth == "speak" {
					bird.speak()
				} else{fmt.Println("Invalid Request")}
			
			} else{fmt.Println("Invalid Request")}
		}
	}
}