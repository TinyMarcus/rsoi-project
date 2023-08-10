package repositories

import (
	"github.com/jinzhu/gorm"
	"loyalties/errors"
	"loyalties/objects"
)

type LoyaltyRepository interface {
	Create(loyalty *objects.Loyalty) error
	Find(username string) (*objects.Loyalty, error)
	Update(*objects.Loyalty) error
	Delete(username string) error
}

type PostgresLoyaltyRepository struct {
	db *gorm.DB
}

func NewPostgresLoyaltyRepository(db *gorm.DB) *PostgresLoyaltyRepository {
	return &PostgresLoyaltyRepository{db: db}
}

func (repository *PostgresLoyaltyRepository) Create(loyalty *objects.Loyalty) error {
	return repository.db.Create(loyalty).Error
}

func (repository *PostgresLoyaltyRepository) Find(username string) (*objects.Loyalty, error) {
	temp := new(objects.Loyalty)
	err := repository.db.
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

func (repository *PostgresLoyaltyRepository) Update(loyalty *objects.Loyalty) error {
	return repository.db.
		Save(loyalty).
		Error
}

func (repository *PostgresLoyaltyRepository) Delete(username string) error {
	record, err := repository.Find(username)
	if err != nil {
		return err
	}

	err = repository.db.
		Where(objects.Loyalty{Username: username}).
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
