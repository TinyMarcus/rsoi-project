package objects

type PaymentDto struct {
	Status string `json:"status"`
	Price  int    `json:"price"`
}

type PaymentDtoWithId struct {
	PaymentUid string `json:"payment_uid"`
	Status     string `json:"status"`
	Price      int    `json:"price"`
}
