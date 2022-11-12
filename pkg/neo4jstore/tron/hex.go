package tron

import (
	"encoding/hex"
	"github.com/blocktree/go-owcdrivers/addressEncoder"
)

func DecodeAddress(addr string) (string, error) {
	input, err := addressEncoder.AddressDecode(addr, addressEncoder.TRON_mainnetAddress)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(append(addressEncoder.TRON_mainnetAddress.Prefix, input...)), nil
}

func EncodeAddress(input string) (string, error) {
	b, err := hex.DecodeString(input)
	if err != nil {
		return "", err
	}
	if len(b) > 20 {
		b = b[1:]
	}

	return addressEncoder.AddressEncode(b, addressEncoder.TRON_mainnetAddress), nil
}

func MustEncodeAddress(address string) (string, error) {
	if len(address) == 42 {
		addressValid, err := EncodeAddress(address)
		if err != nil {
			return "", err
		}

		return addressValid, nil
	}

	return address, nil
}

func MustDecodeAddress(address string) (string, error) {
	if len(address) == 34 {
		addressValid, err := DecodeAddress(address)
		if err != nil {
			return "", err
		}

		return addressValid, nil
	}

	return address, nil
}
