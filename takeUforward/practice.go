package main

import "fmt"

func NumberToWords(num int) string {
	ones := []string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten",
		"eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
	tens := []string{"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}

	if num == 0 {
		return "zero"
	} else if num <= 19 {
		return ones[num]
	}

	res := ""
	if num >= 1000 {
		thComp := num / 1000
		if thComp <= 19 {
			res = ones[thComp] + " thousand "
		} else {
			thCompTens := thComp / 10
			thCompOnes := thComp % 10
			res = tens[thCompTens] + " " + ones[thCompOnes] + " thousand "
		}
		num = num % 1000
	}

	if num >= 100 {
		hundComp := num / 100
		res = res + ones[hundComp] + " hundred "
		num = num % 100
	}

	if num <= 19 {
		res = res + ones[num]
		return res
	}

	tensComp := num / 10
	if tensComp > 0 {
		res = res + tens[tensComp] + " "
	}

	onesComp := num % 10
	if onesComp > 0 {
		res = res + ones[onesComp]
	}

	return res
}

func main() {
	tests := []int{0, 5, 13, 20, 45, 100, 101, 110, 569, 999, 1000, 1234, 20005, 99000, 99009, 99011, 99999}
	for _, t := range tests {
		fmt.Printf("%5d -> %s\n", t, NumberToWords(t))
	}
}
