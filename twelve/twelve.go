package main

import "fmt"

func main() {
	fmt.Println(intToRoman(41))
}

// I             1
// IV			 4
// V             5
// IX			 9
// X             10
// XL			 40
// L             50
// XC			 90
// C             100
// CD			 400
// D             500
// CM			 900
// M             1000

func intToRoman(num int) string {
	romanNumber := make([]rune, 0)
	for num >= 1000 {
		num = num - 1000
		romanNumber = append(romanNumber, 'M')
	}
	if num >= 900 {
		num = num - 900
		romanNumber = append(romanNumber, 'C')
		romanNumber = append(romanNumber, 'M')
	}

	if num >= 500 {
		num = num - 500
		romanNumber = append(romanNumber, 'D')
	}

	if num >= 400 {
		num = num - 400
		romanNumber = append(romanNumber, 'C')
		romanNumber = append(romanNumber, 'D')
	}
	for num >= 100 {
		num = num - 100
		romanNumber = append(romanNumber, 'C')
	}
	if num >= 90 {
		num = num - 90
		romanNumber = append(romanNumber, 'X')
		romanNumber = append(romanNumber, 'C')
	}
	if num >= 50 {
		num = num - 50
		romanNumber = append(romanNumber, 'L')
	}
	if num >= 40 {
		num = num - 40
		romanNumber = append(romanNumber, 'X')
		romanNumber = append(romanNumber, 'L')
	}
	for num >= 10 {
		num = num - 10
		romanNumber = append(romanNumber, 'X')
	}
	if num >= 9 {
		num = num - 9
		romanNumber = append(romanNumber, 'I')
		romanNumber = append(romanNumber, 'X')
	}

	if num >= 5 {
		num = num - 5
		romanNumber = append(romanNumber, 'V')
	}

	if num >= 4 {
		num = num - 4
		romanNumber = append(romanNumber, 'I')
		romanNumber = append(romanNumber, 'V')
	}

	for num >= 1 {
		num = num - 1
		romanNumber = append(romanNumber, 'I')
	}

	return string(romanNumber)
}
