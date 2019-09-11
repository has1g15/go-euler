package LUHN

import (
	"strconv"
	"strings"
)

//A 10 digit card account number that has 2 leading digits ‘91’ (the mask), and a check digit ‘C’ in the last position, e.g. 91abcdefgC
//The letters abcdefg represent a unique account numbers allowing 10 million -1 card holders.
//The account numbers must be numeric in the range 0 to 9.
//
//Example valid card account number: 9134567891
//
//LUHN algorithm:
//1. Take copy of the original number.
//2. Take every 2nd digit from the right hand side of the copy.
//3. Double it.
//4. If the result if > 9 then subtract 9.
//5. Overwrite the digit in the copy with the result.
//6. Sum all of the digits.
//7. Multiple the total by 9.
//8. Apply MOD 10 and the result becomes the check digit.
//9. Compare the check digit against the original number check digit.
//10. If they match, this is a valid card account number.
//
//Prior to running the LUHN algorithm on a test number the following checks need to be satisfied:
//1. The length is exactly 10 digits.
//2. All of the digits are integers.
//3. There are no alphanumeric characters in the number.
//4. The first two digits are 91
func luhn(accountNumber string) bool{
	originalCheckDigit, _ := strconv.ParseInt(accountNumber[len(accountNumber)-1:], 0, 64)
	numCopy := accountNumber

	for x := len(numCopy) - 2; x > 0; x-= 2 {
		toInt, _ := strconv.ParseInt(string(numCopy[x]), 0, 64)
		doubled := toInt * 2
		if doubled > 9 {
			doubled = doubled - 9
		}
		numCopy = strings.Replace(numCopy, string(numCopy[x]), string(doubled), 1)
	}

	var digitTotal int64 = 0
	for y := 0; y < len(numCopy); y++ {
		toInt, _ := strconv.ParseInt(string(numCopy[y]), 0, 64)
		digitTotal += toInt
	}

	totalBy9 := digitTotal * 9
	checkDigit := totalBy9 % 10

	return checkDigit == originalCheckDigit
}

func checkLength(accountNumber string) bool {
	return len(accountNumber) > 9 && len(accountNumber) < 17
}

func checkType(accountNumber string) bool {
	_, err := strconv.ParseFloat(accountNumber, 64)
	return err == nil
}

func scaleTo10Digits(accountNumber string) string{
	return accountNumber[len(accountNumber)-10:]
}