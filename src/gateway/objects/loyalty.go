package objects

type LoyaltyDto struct {
	ReservationCount int    `json:"reservation_count"`
	Status           string `json:"status"`
	Discount         int    `json:"discount"`
}

type UserInfoResponse struct {
	Loyalties    LoyaltyDto                     `json:"loyalties"`
	Reservations ReservationsPaginationResponse `json:"reservations"`
}
