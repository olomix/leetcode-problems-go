package p150

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	for len(tokens) > 1 {
		fmt.Println("was", tokens)
		for i := 0; i <= len(tokens); i++ {
			switch tokens[i] {
			case "+":
				x := iToStr(tokens[i-2])
				y := iToStr(tokens[i-1])
				tokens[i-2] = strconv.Itoa(x + y)
			case "-":
				x := iToStr(tokens[i-2])
				y := iToStr(tokens[i-1])
				tokens[i-2] = strconv.Itoa(x - y)
			case "*":
				x := iToStr(tokens[i-2])
				y := iToStr(tokens[i-1])
				tokens[i-2] = strconv.Itoa(x * y)
			case "/":
				x := iToStr(tokens[i-2])
				y := iToStr(tokens[i-1])
				tokens[i-2] = strconv.Itoa(x / y)
			default:
				continue
			}
			tokens = append(tokens[:i-1], tokens[i+1:]...)
			break
		}
		fmt.Println("now", tokens)
	}
	return iToStr(tokens[0])
}

func iToStr(in string) int {
	i, err := strconv.ParseInt(in, 10, 32)
	if err != nil {
		panic(err)
	}
	return int(i)
}
