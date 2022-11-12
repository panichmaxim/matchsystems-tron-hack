package models

type BillingStatistics struct {
	Btc  *BillingStatisticsBlockchain `json:"btc"`
	Eth  *BillingStatisticsBlockchain `json:"eth"`
	Tron *BillingStatisticsBlockchain `json:"tron"`
}

type BillingStatisticsBlockchain struct {
	Total      int64          `json:"total"`
	Categories map[string]int `json:"categories"`
}

type BillingStatisticsRisk struct {
	Name  string
	Total int64
}

type BillingStatisticsCategory struct {
	Name string
	Risk float64
}

type StatisticsSummary struct {
	Network string `json:"network"`
	Total   int    `json:"total"`
}
