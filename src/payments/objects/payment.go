package objects

type Payment struct {
	Id         int    `json:"id" gorm:"primary_key; index"`
	PaymentUid string `json:"payment_uid" gorm:"not null; unique"`
	Status     string `json:"status" gorm:"not null" sql:"DEFAULT: 'PAID'"`
	Price      int    `json:"price" gorm:"not null"`
}

func (Payment) TableName() string {
	return "payment"
}

type HotelInfoDto struct {
	HotelUid    string `json:"hotel_uid"`
	Name        string `json:"name"`
	FullAddress string `json:"full_address"`
	Stars       int    `json:"stars"`
}

type HotelResponseDto struct {
	HotelUid string `json:"hotel_uid"`
	Name     string `json:"name"`
	Country  string `json:"country"`
	City     string `json:"city"`
	Address  string `json:"address"`
	Stars    int    `json:"stars"`
	Price    int    `json:"price"`
}

type LoyaltyDto struct {
	ReservationCount int    `json:"reservation_count"`
	Status           string `json:"status"`
	Discount         int    `json:"discount"`
}

type PaymentDto struct {
	Status string `json:"status"`
	Price  int    `json:"price"`
}

type PaymentDtoWithId struct {
	PaymentUid string `json:"payment_uid"`
	Status     string `json:"status"`
	Price      int    `json:"price"`
}

type ReservationResponseDto struct {
	ReservationUid string           `json:"reservation_uid"`
	Hotel          HotelInfoDto     `json:"hotel"`
	StartDate      string           `json:"start_date"`
	EndDate        string           `json:"end_date"`
	Status         string           `json:"status"`
	Payment        PaymentDtoWithId `json:"payment"`
}

func ToPaymentDtoResponse(payment *Payment) *PaymentDtoWithId {
	return &PaymentDtoWithId{
		PaymentUid: payment.PaymentUid,
		Status:     payment.Status,
		Price:      payment.Price,
	}
}
