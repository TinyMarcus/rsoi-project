package repositories

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"reservations/errors"
	"reservations/objects"
)

type HotelRepository interface {
	Create(hotel *objects.Hotel) error
	Find(hotelUid string) (*objects.Hotel, error)
	FindHotels(page int, pageSize int) []objects.Hotel
	Update(hotel *objects.Hotel) error
	Delete(hotelUid string) error
}

type PostgresHotelRepository struct {
	db *gorm.DB
}

func NewPostgresHotelRepository(db *gorm.DB) *PostgresHotelRepository {
	return &PostgresHotelRepository{db: db}
}

func (repository *PostgresHotelRepository) Create(hotel *objects.Hotel) error {
	hotel.HotelUid = uuid.New().String()

	return repository.db.Create(hotel).Error
}

func (repository *PostgresHotelRepository) Find(hotelUid string) (*objects.Hotel, error) {
	temp := new(objects.Hotel)
	err := repository.db.
		First(temp, "hotel_uid = ?", hotelUid).
		Error

	switch err {
	case nil:
		return temp, err
	case gorm.ErrRecordNotFound:
		return nil, errors.RecordNotFound
	default:
		return nil, errors.UnknownError
	}
}

func paginate(page int, pageSize int) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func (repository *PostgresHotelRepository) FindHotels(page int, pageSize int) []objects.Hotel {
	temp := []objects.Hotel{}
	repository.db.
		Scopes(paginate(page, pageSize)).
		Model(&objects.Hotel{}).
		Find(&temp)

	return temp
}

func (repository *PostgresHotelRepository) Update(hotel *objects.Hotel) error {
	return repository.db.
		Save(hotel).
		Error
}

func (repository *PostgresHotelRepository) Delete(hotelUid string) error {
	record, err := repository.Find(hotelUid)
	if err != nil {
		return err
	}

	err = repository.db.
		Where(objects.Hotel{HotelUid: hotelUid}).
		Delete(record).
		Error

	switch err {
	case nil:
		return err
	case gorm.ErrRecordNotFound:
		return errors.RecordNotFound
	default:
		return errors.UnknownError
	}
}
