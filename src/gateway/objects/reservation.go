package objects

type CreateReservationRequestDto struct {
	PaymentUid string `json:"payment_uid"`
	HotelUid   string `json:"hotel_uid"`
	Status     string `json:"status"`
	StartDate  string `json:"start_date"`
	EndDate    string `json:"end_date"`
}

type ReservationRequestDto struct {
	HotelUid  string `json:"hotel_uid"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type ReservationResponseDto struct {
	ReservationUid string           `json:"reservation_uid"`
	Hotel          HotelInfoDto     `json:"hotel"`
	StartDate      string           `json:"start_date"`
	EndDate        string           `json:"end_date"`
	Status         string           `json:"status"`
	Payment        PaymentDtoWithId `json:"payment"`
}

type ReservationDto struct {
	ReservationUid string           `json:"reservation_uid"`
	Hotel          HotelInfoDto     `json:"hotel"`
	StartDate      string           `json:"start_date"`
	EndDate        string           `json:"end_date"`
	Status         string           `json:"status"`
	Payment        PaymentDtoWithId `json:"payment"`
}

type ReservationsPaginationResponse struct {
	Page          int                      `json:"page"`
	PageSize      int                      `json:"pageSize"`
	TotalElements int                      `json:"totalElements"`
	Items         []ReservationResponseDto `json:"items"`
}

type ReservationDeletionResponseDto struct {
	PaymentUid string `json:"payment_uid"`
}
