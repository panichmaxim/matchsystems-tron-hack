package validator

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBitcoinHashTransactionValidation(t *testing.T) {
	matched := btcTransactionRegex.MatchString("00000000000000001e8d6829a8a21adc5d38d0a473b144b6765798e61f98bd1d")
	require.True(t, matched)

	matched = btcTransactionRegex.MatchString("e4199152beb77e9d9a6b6da2e5b821c67c979ea5a819bb5a8239f2827df59489")
	require.True(t, matched)
}

func TestBitcoinHashBlockValidation(t *testing.T) {
	matched := btcBlockTransactionRegex.MatchString("000000000bf0032fafa9771feb7d686671e4eef88c5ee8c19c8bbea5b65918b5")
	require.True(t, matched)

	matched = btcBlockTransactionRegex.MatchString("e4199152beb77e9d9a6b6da2e5b821c67c979ea5a819bb5a8239f2827df59489")
	require.False(t, matched)
}

func TestBitcoinAddressValidation(t *testing.T) {
	matched := btcAddressRegex.MatchString("bc1qm34lsc65zpw79lxes69zkqmk6ee3ewf0j77s3h")
	require.True(t, matched)

	matched = btcAddressRegex.MatchString("35bdfMb4LwToWsZH36fwZZkFTMHES8kzFg")
	require.True(t, matched)

	matched = btcAddressRegex.MatchString("e4199152beb77e9d9a6b6da2e5b821c67c979ea5a819bb5a8239f2827df59489")
	require.False(t, matched)

	matched = btcAddressRegex.MatchString("bc1q7ug4w4as2sefar89q057hnmxkakp58a25535ttlmurn6cncs8tms4e7gp2")
	require.True(t, matched)
}
