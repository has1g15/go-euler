package LUHN

import (
	"testing"
)

func TestAccountNumber(t *testing.T) {
	accountNumber := "374209134567891"
	scaledNumber := scaleTo10Digits(accountNumber)
	if checkLength(scaledNumber) && checkType(scaledNumber) {
		if !luhn(scaledNumber) {
			t.Error("Not valid account number")
		}
	}
}

//374209134567891
//376000188331009
//374200462621008
