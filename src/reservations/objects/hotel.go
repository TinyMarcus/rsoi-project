package objects

type Hotel struct {
	Id       int    `json:"id" gorm:"primary_key; index"`
	HotelUid string `json:"hotel_uid" gorm:"not null; unique"`
	Name     string `json:"name" gorm:"not null;"`
	Country  string `json:"country" gorm:"not null;""`
	City     string `json:"city" gorm:"not null;""`
	Address  string `json:"address" gorm:"not null" sql:"DEFAULT: 'BRONZE'"`
	Stars    int    `json:"stars" gorm:"not null"`
	Price    int    `json:"price" gorm:"not null"`
}

func (Hotel) TableName() string {
	return "hotel"
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

type CreateHotelRequestDto struct {
	Name    string `json:"name"`
	Country string `json:"country"`
	City    string `json:"city"`
	Address string `json:"address"`
	Stars   int    `json:"stars"`
	Price   int    `json:"price"`
}

func ToHotelInfoDto(hotel *Hotel) *HotelInfoDto {
	return &HotelInfoDto{
		HotelUid:    hotel.HotelUid,
		Name:        hotel.Name,
		FullAddress: hotel.Country + ", " + hotel.City + ", " + hotel.Address,
		Stars:       hotel.Stars,
	}
}

func (hotel *Hotel) ToHotelResponseDto() *HotelResponseDto {
	return &HotelResponseDto{
		HotelUid: hotel.HotelUid,
		Name:     hotel.Name,
		Country:  hotel.Country,
		City:     hotel.City,
		Address:  hotel.Address,
		Stars:    hotel.Stars,
		Price:    hotel.Price,
	}
}

func ToHotelResponses(hotels []Hotel) []HotelResponseDto {
	resps := make([]HotelResponseDto, len(hotels))
	for k, v := range hotels {
		resps[k] = *v.ToHotelResponseDto()
	}
	return resps
}

type HotelPaginationResponse struct {
	Page          int                `json:"page"`
	PageSize      int                `json:"pageSize"`
	TotalElements int                `json:"totalElements"`
	Items         []HotelResponseDto `json:"items"`
}
