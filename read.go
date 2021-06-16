package main

import (
	"fmt"
	"bufio"
	"strings"
	"os"
)

type name struct {
	fname string
	lname string
}

func main(){

	var filename string

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter name of file: ")

	if scanner.Scan() {
		filename = scanner.Text()
	}
	file, _ := os.Open(filename)
    defer file.Close()

	s := make([]name, 0, 20)
    scanner2 := bufio.NewScanner(file)
	
    for scanner2.Scan() {
        fullname := strings.Fields(scanner2.Text())
		var n name
		n.fname = fullname[0]
		n.lname = fullname[1]
		s = append(s, n)
    }

	for _, nam := range s {

        fmt.Println("First Name:", nam.fname, "\nLast Name: ", nam.lname, "\n")
    }
}