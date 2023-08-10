package repositories

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"reservations/errors"
	"reservations/objects"
)

type ReservationRepository interface {
	Create(reservation *objects.Reservation) error
	Find(username string, page int, page_size int) []objects.Reservation
	FindByReservationUid(username, reservationUid string) (*objects.Reservation, error)
	Update(reservation *objects.Reservation) error
	Delete(reservationUid string) error
}

type PostgresReservationRepository struct {
	db *gorm.DB
}

func NewPostgresReservationRepository(db *gorm.DB) *PostgresReservationRepository {
	return &PostgresReservationRepository{db: db}
}

func (repository *PostgresReservationRepository) Create(reservation *objects.Reservation) error {
	reservation.ReservationUid = uuid.New().String()

	return repository.db.Create(reservation).Error
}

func paginateReservations(page int, pageSize int) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func (repository *PostgresReservationRepository) Find(username string, page int, pageSize int) []objects.Reservation {
	temp := []objects.Reservation{}
	repository.db.
		Scopes(paginate(page, pageSize)).
		Model(&objects.Reservation{}).
		Where("username = ?", username).
		Find(&temp)

	return temp
}

func (repository *PostgresReservationRepository) FindByReservationUid(username, reservationUid string) (*objects.Reservation, error) {
	temp := new(objects.Reservation)
	err := repository.db.
		First(temp, "reservation_uid = ?", reservationUid).
		First(temp, "username = ?", username).
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

func (repository *PostgresReservationRepository) Update(reservation *objects.Reservation) error {
	return repository.db.
		Save(reservation).
		Error
}

func (repository *PostgresReservationRepository) Delete(reservationUid string) error {
	return repository.db.
		Model(&objects.Reservation{}).
		Where("reservation_uid = ?", reservationUid).
		Update("status", "CANCELED").
		Error
}
