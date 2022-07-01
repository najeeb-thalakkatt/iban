// All service level methods go here

/*
Check that the total IBAN length is correct as per the country. If not, the IBAN is invalid
Move the four initial characters to the end of the string
Replace each letter in the string with two digits, thereby expanding the string, where A = 10, B = 11, ..., Z = 35
Interpret the string as a decimal integer and compute the remainder of that number on division by 97
If the remainder is 1, the check digit test is passed and the IBAN might be valid.

Example (fictitious United Kingdom bank, sort code 12-34-56, account number 98765432):

• IBAN:		GB82 WEST 1234 5698 7654 32
• Rearrange:		W E S T12345698765432 G B82
• Convert to integer:		3214282912345698765432161182
• Compute remainder:		3214282912345698765432161182	mod 97 = 1

*/

package service

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

func CheckIBAN(iban string) (bool, error) {
	//1. Check that the total IBAN length is correct as per the country. If not, the IBAN is invalid
	country := getCountryCode(iban)
	if _, ok := CountryLength[country]; !ok {
		return false, fmt.Errorf("country code %s not supported", country)
	}
	iban = strings.Replace(iban, " ", "", -1)
	if len(iban) != CountryLength[country] {
		return false, fmt.Errorf("IBAN length for %s is %d", country, CountryLength[country])
	}
	swappedIBAN := theSwap(iban)    //Move the four initial characters to the end of the string
	intIBAN := Replace(swappedIBAN) //Replace each letter in the string with two digits, thereby expanding the string, where A = 10, B = 11, ..., Z = 35
	mod := Mod97(intIBAN)           //Interpret the string as a decimal integer and compute the remainder of that number on division by 97
	if mod != 1 {                   //If the remainder is 1, the check digit test is passed and the IBAN might be valid.
		return false, fmt.Errorf("IBAN not valid")
	}
	return true, nil

}

func intMapper() map[string]string {
	strToIntMapper := map[string]string{}
	val := 0
	for i := 0; i <= 9; i++ {
		strToIntMapper[strconv.Itoa(i)] = strconv.Itoa(val)
		val += 1
	}
	for r := 'A'; r <= 'Z'; r++ {
		strToIntMapper[string(r)] = strconv.Itoa(val)
		val += 1
	}
	return strToIntMapper
}

func getCountryCode(iban string) string {
	return iban[0:2]
}

func theSwap(iban string) string {
	return iban[4:] + iban[:4]
}

func Replace(iban string) string {
	intMap := intMapper()
	for s_from, s_to := range intMap {
		iban = strings.Replace(iban, s_from, s_to, -1)
	}
	return iban
}

func Mod97(iban string) int64 {
	bban := new(big.Int)
	bban, _ = bban.SetString(iban, 10)
	ninetyseven := new(big.Int)
	ninetyseven, _ = ninetyseven.SetString("97", 10)
	mod97 := new(big.Int)
	return mod97.Mod(bban, ninetyseven).Int64()

}
