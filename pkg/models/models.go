package models

func GetModels() []interface{} {
	return []interface{}{
		(*AuthToken)(nil),
		(*User)(nil),
		(*AccessRequest)(nil),
		(*Billing)(nil),
		(*BillingPacket)(nil),
		(*BillingRequest)(nil),
		(*BillingKey)(nil),
		(*Category)(nil),
	}
}
