package main

import (
	"fmt"
)

func infixToPostfix(infix string) string {
	//Assign special characters int values
	special := map[rune]int{"*": 10 , ".": 9 , "|": 8}


	//Empty array of runes
	pFix := []rune{}

	stack := []rune{}

	//Loop over string and convert to postfix
	for _, r := range infix {
		//range is used to convert type string into type rune

	}

	return string(pFix)
}

func main() {
	//Input a.b.c
	fmt.Println("Infix: a.b.c")
	fmt.Println("Postfis: ", infixToPostfix("a.b.c"))
	//Should return ab.c*
}