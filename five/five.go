package five

import (
	"fmt"
)

func main() {
	fmt.Println("test")
}

func longestPalindrome(s string) string {
	charArray := make([]rune, 0)
	index := 0
	longest := 0
	openForThree := true
	for i, singleChar := range s {
		if singleChar == charArray[index] {
			if index > 0 {
				charArray = charArray[:index-1]
			} else {
				charArray = make([]rune, 0)
			}

			longest = longest + 2

		}
	}
}
