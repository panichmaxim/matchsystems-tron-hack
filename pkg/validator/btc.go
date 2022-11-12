package validator

import (
	"regexp"
)

var btcAddressRegex = regexp.MustCompile(`^([11]|[13]|bc1)[\w+]{25,59}$`)
var btcBlockTransactionRegex = regexp.MustCompile(`^0{8}[a-fA-F\d]{56}$`)
var btcTransactionRegex = regexp.MustCompile(`^[a-fA-F\d]{64}$`)
