package LUHN

import (
	"testing"
)

func TestAccountNumber(t *testing.T) {
	accountNumber := "xxx"
	scaledNumber := scaleTo10Digits(accountNumber)
	if checkLength(scaledNumber) && checkType(scaledNumber) {
		if !luhn(scaledNumber) {
			t.Error("Not valid account number")
		}
	}
}
