package objects

type Reservation struct {
	Id             int    `json:"id" gorm:"primary_key; index"`
	ReservationUid string `json:"reservation_uid" gorm:"not null; unique"`
	Username       string `json:"username" gorm:"not null;"`
	PaymentUid     string `json:"payment_uid" gorm:"not null;""`
	HotelUid       string `json:"hotel_uid" gorm:"not null;""`
	Status         string `json:"status" gorm:"not null" sql:"DEFAULT: 'BRONZE'"`
	StartDate      string `json:"start_date" gorm:"not null"`
	EndDate        string `json:"end_date" gorm:"not null"`
}

func (Reservation) TableName() string {
	return "reservation"
}

type PaymentDtoWithId struct {
	PaymentUid string `json:"payment_uid"`
	Status     string `json:"status"`
	Price      int    `json:"price"`
}

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

type ReservationDeletionResponseDto struct {
	PaymentUid string `json:"payment_uid"`
}

func ToReservationResponseDto(reservation *Reservation, hotel *HotelInfoDto, payment *PaymentDtoWithId) *ReservationResponseDto {
	return &ReservationResponseDto{
		ReservationUid: reservation.ReservationUid,
		Hotel:          *hotel,
		StartDate:      reservation.StartDate,
		EndDate:        reservation.EndDate,
		Status:         reservation.Status,
		Payment:        *payment,
	}
}

func ToReservationResponses(reservations []Reservation, hotels []HotelInfoDto, payments []PaymentDtoWithId) []ReservationResponseDto {
	resps := make([]ReservationResponseDto, len(reservations))
	for k, v := range reservations {
		resps[k] = *ToReservationResponseDto(&v, &hotels[k], &payments[k])
	}
	return resps
}

type ReservationsPaginationResponse struct {
	Page          int                      `json:"page"`
	PageSize      int                      `json:"pageSize"`
	TotalElements int                      `json:"totalElements"`
	Items         []ReservationResponseDto `json:"items"`
}
