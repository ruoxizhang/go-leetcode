package main

import "fmt"

func main() {
	fmt.Println(longestValidParenthesesNew(")()())"))
}

func longestValidParentheses(s string) int {
	if len(s) < 2 {
		return 0
	}
	pStack := make([]rune, len(s))
	resultStack := make([]int, len(s))
	resultStack[0] = 0
	lastIndex := 0

	result := 0
	for _, value := range s {
		if value == '(' {
			resultStack = append(resultStack, 0)
			lastIndex++
			pStack = append(pStack, value)
		} else if value == ')' {
			if len(pStack)-1 >= 0 && pStack[len(pStack)-1] == '(' {
				resultStack[lastIndex] = resultStack[lastIndex] + 2
				resultStack[lastIndex-1] = resultStack[lastIndex-1] + resultStack[lastIndex]
				resultStack = resultStack[:lastIndex]
				lastIndex--
				pStack = pStack[:len(pStack)-1]
			} else {
				for _, value := range resultStack {
					if value > result {
						result = value
					}
				}
				lastIndex = 0
				resultStack = make([]int, len(s))
				resultStack[0] = 0
			}
		} else {
			return 0
		}
	}
	for _, value := range resultStack {
		if value > result {
			result = value
		}
	}

	return result
}

func longestValidParenthesesNew(s string) int {
	sLeng := len(s)
	if sLeng < 2 {
		return 0
	}
	dpResult := make([]int, sLeng+1)
	dpResult[0] = 0
	for i, value := range s {
		if value == '(' {
			dpResult[i+1] = 0
		} else if value == ')' {
			if i > 0 {
				if s[i-1] == '(' {
					dpResult[i+1] = dpResult[i-1] + 2
				} else if s[i-1] == ')' {
					if i-1-dpResult[i] >= 0 {
						if s[i-1-dpResult[i]] == '(' {
							dpResult[i+1] = dpResult[i] + 2 + dpResult[i-1-dpResult[i]]
						}
					} else {
						dpResult[i+1] = 0
					}
				} else {
					return 0
				}
			} else {
				dpResult[i+1] = 0
			}
		} else {
			return 0
		}
	}
	result := 0
	for _, value := range dpResult {
		if value > result {
			result = value
		}
	}
	return result
}
