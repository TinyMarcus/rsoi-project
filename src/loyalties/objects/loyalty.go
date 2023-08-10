package objects

type Loyalty struct {
	Id               int    `json:"id" gorm:"primary_key; index"`
	Username         string `json:"username" gorm:"not null; unique"`
	ReservationCount int    `json:"reservation_count" gorm:"not null" sql:"DEFAULT: 0"`
	Status           string `json:"status" gorm:"not null" sql:"DEFAULT: 'BRONZE'"`
	Discount         int    `json:"discount" gorm:"not null"`
}

func (Loyalty) TableName() string {
	return "loyalty"
}

type LoyaltyDto struct {
	ReservationCount int    `json:"reservation_count"`
	Status           string `json:"status"`
	Discount         int    `json:"discount"`
}

func ToLoyaltyDtoResponse(loyalty *Loyalty) *LoyaltyDto {
	return &LoyaltyDto{
		ReservationCount: loyalty.ReservationCount,
		Status:           loyalty.Status,
		Discount:         loyalty.Discount,
	}
}
