package main

import "fmt"

func main() {
	fmt.Printf("Largest Product Palindrome: %d", findLargestPalindrome())
}

func findLargestPalindrome() int {
	palindromes := getProductPalindromes()
	max := palindromes[0]
	for  _, palindrome := range palindromes {
		if palindrome> max {
			max = palindrome
		}
	}
	return max
}

func getProductPalindromes() []int {
	product := 0
	var palindromes[]int
	for i := 999; i >= 100; i-- {
		for j := 999; j >= 100; j-- {
			product = i * j
			if checkPalindrome(product) {
				palindromes = append(palindromes, product)
			}
		}
	}
	return palindromes
}

func checkPalindrome(product int) bool{
	reverse := 0
	original := product

	for {
		//Get last digit of number
		remainder := product % 10

		//Multiply by 10 to increment number of digits of reverse number
		//Add last digit of original number to reverse number
		reverse = reverse * 10 + remainder

		//Shave off last digit of original number
		product /= 10

		//Break when original number no longer contains digits
		if product == 0 {
			break
		}
	}

	if original == reverse {
		return true
	} else {
		return false
	}
}
