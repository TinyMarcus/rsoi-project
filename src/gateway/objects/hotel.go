package objects

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

type CreateHotelRequestDto struct {
	Name    string `json:"name"`
	Country string `json:"country"`
	City    string `json:"city"`
	Address string `json:"address"`
	Stars   int    `json:"stars"`
	Price   int    `json:"price"`
}

type HotelPaginationResponse struct {
	Page          int                `json:"page"`
	PageSize      int                `json:"pageSize"`
	TotalElements int                `json:"totalElements"`
	Items         []HotelResponseDto `json:"items"`
}
