package neo4jstore

type Node struct {
	ID     int64
	Labels []string
	Props  map[string]interface{}
}

type TypedNode[T any] struct {
	ID   int64
	Data *T
}

type FindAddressByHashNode struct {
	Address string  `json:"Address"`
	Total   float64 `json:"Totalin"`
}

type FindReportedRiskNode struct {
	Risk     float64 `json:"Score"`
	Category int     `json:"Category"`
}

type RiskData struct {
	Category int     `json:"category"`
	Risk     float64 `json:"risk"`
}

type Risk struct {
	Risk       *float64        `json:"risk"`
	Reported   *RiskData       `json:"reported"`
	Wallet     *RiskData       `json:"wallet"`
	Calculated *CalculatedRisk `json:"calculated"`
}
