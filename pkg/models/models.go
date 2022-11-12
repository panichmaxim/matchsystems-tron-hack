package models

func GetModels() []interface{} {
	return []interface{}{
		(*AuthToken)(nil),
		(*User)(nil),
		(*AccessRequest)(nil),
		(*Billing)(nil),
		(*BillingPacket)(nil),
		(*BillingRequest)(nil),
		(*BillingRisk)(nil),
		(*BillingKey)(nil),
		(*Category)(nil),
	}
}
