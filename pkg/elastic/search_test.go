package elastic

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestElasticImpl_Information(t *testing.T) {
	es, err := NewElastic(testElasticHosts)
	require.NoError(t, err)

	r, _, err := es.Search(ctx, "bc1qxy2kgdygjrsqtzq2n0yrf2493p83kkfjhx0wlh", 0, 1)
	require.NoError(t, err)
	require.Len(t, r, 1)

	first := r[0]
	payload, err := json.Marshal(first.Data)
	require.NoError(t, err)
	require.Equal(t, `{"accepted":"2","aml_risk":"{BTC: {total: {risk: 1}, high: {illegal: 0, darkmarket: 0, darkservice: 0, fraud: 0, gambling: 0, mixer: 0, ransom: 0, scam: 1, stolen: 0}, moderate: {atm: 0, ml_high: 0, ml_mod: 0, ml_very_high: 0, ptp_high: 0}, min: {market: 0, miner: 0, wallet: 0, ptp_risk_low: 0, payment: 0, risk_low: 0}}, ETH: {total: {risk: 0.297}, high: {illegal: 0, darkmarket: 0, darkservice: 0, fraud: 0, gambling: 0.001, mixer: 0, ransom: 0, scam: 0, stolen: 0}, moderate: {atm: 0, ml_high: 0.699, ml_mod: 0.064, ml_very_high: 0, ptp_high: 0.209}, min: {market: 0.013, miner: 0, wallet: 0.001, ptp_risk_low: 0.012, payment: 0, risk_low: 0}}}","chat_id":"5039424408","country":"Украина украинский","descr":"Wts coinlist accounts","forum_lang":"английский","forum_ref":"https://t.me/mediasocialmarket","has_txid":"t","name":"Roman","refs":"https://t.me/mediasocialmarket","text":"German, [23.02.2022 9:22]helloGerman, [23.02.2022 9:22]need 2 acGerman, [23.02.2022 9:22]I will pay firstRoman, [23.02.2022 9:23]Hey, You wanna do with escrow?German, [23.02.2022 9:23]noGerman, [23.02.2022 9:23]i pay firstRoman, [23.02.2022 9:24]How do you wanna pay ? With crypto ?German, [23.02.2022 9:24]yesGerman, [23.02.2022 9:24]Send btc eth and usdt I will look at the commission and payGerman, [23.02.2022 9:24]ua  country?Roman, [23.02.2022 9:24][В ответ на German]yepRoman, [23.02.2022 9:26][В ответ на German]Ok, I will setup my pc cause I was moving to new apartment and send you in 2~ hours ok ?German, [23.02.2022 9:27]can you send me wallet addresses now?German, [23.02.2022 9:27]or noRoman, [23.02.2022 9:27]I guess yesGerman, [23.02.2022 9:27]goodRoman, [23.02.2022 9:28]Usdt in what network ?Roman, [23.02.2022 9:29]TRC20 or solana ?German, [23.02.2022 9:29]does not matterRoman, [23.02.2022 9:29]OkGerman, [23.02.2022 9:29]trc20Roman, [23.02.2022 9:30]TJ8c27wnyc7dLReMpTzLAhM2iSgiHATwkM  Usdt trc200x26c06c3a9f250093ad24a26ec528a00ee2cb529aEthGerman, [23.02.2022 9:30]and btcRoman, [23.02.2022 9:30][В ответ на Roman]Usdt Trc has to be lowest commisionRoman, [23.02.2022 9:31]I wouldLike to not get in btcRoman, [23.02.2022 9:32]It will also be big commisionGerman, [23.02.2022 9:32]ok, but I'm just wondering what my commission will beRoman, [23.02.2022 9:32][В ответ на German]Trc is always lowestGerman, [23.02.2022 9:32]I will not pay for btcRoman, [23.02.2022 9:32]Same as solanaGerman, [23.02.2022 9:32]just to knowRoman, [23.02.2022 9:33]bc1qxy2kgdygjrsqtzq2n0yrf2493p83kkfjhx0wlhThat’s not mine just exampleRoman, [23.02.2022 9:33]Doesn’t matter whose","uid":"4a1fc0a9-978b-11ec-8590-ecb1d77beb80","username":""}`, string(payload))

	r, _, err = es.Search(ctx, "bc1qxy2kgdygjrsqtzq2n0yrf2493p83kkfjhx0wlh", 1, 1)
	require.NoError(t, err)
	require.Len(t, r, 1)

	r, _, err = es.Search(ctx, "bc1qxy2kgdygjrsqtzq2n0yrf2493p83kkfjhx0wlh", 2, 1)
	require.NoError(t, err)
	require.Len(t, r, 1)

	r, _, err = es.Search(ctx, "bc1qxy2kgdygjrsqtzq2n0yrf2493p83kkfjhx0wlh", 30, 1)
	require.NoError(t, err)
	require.Len(t, r, 0)
}
