package main

import "fmt"

func defangIPaddr(address string) string {
	result := ""

	for _, symbol := range address {
		symbol := string(symbol)

		if symbol != "." {
			result += symbol
		} else {
			result += "[.]"
		}
	}

	return result
}

func main() {
	fmt.Println(defangIPaddr("1.1.1.1"))
	fmt.Println(defangIPaddr("255.100.50.0"))
}
