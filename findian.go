package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("Enter String: ")
	var str string
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		str = scanner.Text()
	}

	res := strings.ToLower(str)
	i := strings.Index(res, "i")
	n := strings.Index(res, "n")

	if (i == 0) && (n == len(res)-1) && strings.Contains(res, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
