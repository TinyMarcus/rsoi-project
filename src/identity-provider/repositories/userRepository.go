package repositories

import (
	"github.com/jinzhu/gorm"
	"identity-provider/errors"
	"identity-provider/objects"
)

type UserRepository interface {
	Create(loyalty *objects.User) (*objects.User, error)
	Find(username string) (*objects.User, error)
	CheckCredentials(username, password string) (bool, error)
	Update(*objects.User) error
	Delete(username string) error
}

type PostgresUserRepository struct {
	db *gorm.DB
}

func NewPostgresUserRepository(db *gorm.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (repository *PostgresUserRepository) Create(user *objects.User) (*objects.User, error) {
	return user, repository.db.Create(user).Error
}

func (repository *PostgresUserRepository) Find(username string) (*objects.User, error) {
	temp := new(objects.User)
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

func (repository *PostgresUserRepository) CheckCredentials(username, password string) (bool, error) {
	temp := new(objects.User)
	err := repository.db.
		First(temp, "username = ?", username).
		First(temp, "password = ?", password).
		Error

	switch err {
	case nil:
		return true, err
	case gorm.ErrRecordNotFound:
		return false, errors.RecordNotFound
	default:
		return false, errors.UnknownError
	}
}

func (repository *PostgresUserRepository) Update(loyalty *objects.User) error {
	return repository.db.
		Save(loyalty).
		Error
}

func (repository *PostgresUserRepository) Delete(username string) error {
	record, err := repository.Find(username)
	if err != nil {
		return err
	}

	err = repository.db.
		Where(objects.User{Username: username}).
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
