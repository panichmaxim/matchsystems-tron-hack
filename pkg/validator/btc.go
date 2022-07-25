package validator

import (
	"regexp"
)

var btcAddressRegex = regexp.MustCompile(`^([13]|bc1)[\w+]{27,59}$`)
var btcBlockTransactionRegex = regexp.MustCompile(`^0{8}[a-fA-F\d]{56}$`)
var btcTransactionRegex = regexp.MustCompile(`^[a-fA-F\d]{64}$`)

func isValidBtcTransaction(txid string) bool {
	matched := btcTransactionRegex.MatchString(txid)
	if !matched {
		return matched
	}

	return !btcBlockTransactionRegex.MatchString(txid)
}
