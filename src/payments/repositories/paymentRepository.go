package repositories

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"payments/errors"
	"payments/objects"
)

type PaymentRepository interface {
	Create(payment *objects.Payment) error
	Find(paymentUid string) (*objects.Payment, error)
	Update(payment *objects.Payment) error
	Delete(paymentUid string) error
}

type PostgresPaymentRepository struct {
	db *gorm.DB
}

func NewPostgresPaymentRepository(db *gorm.DB) *PostgresPaymentRepository {
	return &PostgresPaymentRepository{db: db}
}

func (repository *PostgresPaymentRepository) Create(payment *objects.Payment) error {
	payment.PaymentUid = uuid.New().String()
	payment.Status = "PAID"

	return repository.db.Create(payment).Error
}

func (repository *PostgresPaymentRepository) Find(paymentUid string) (*objects.Payment, error) {
	temp := new(objects.Payment)
	err := repository.db.
		First(temp, "payment_uid = ?", paymentUid).
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

func (repository *PostgresPaymentRepository) Update(payment *objects.Payment) error {
	return repository.db.
		Save(payment).
		Error
}

func (repository *PostgresPaymentRepository) Delete(paymentUid string) error {
	return repository.db.
		Model(&objects.Payment{}).
		Where("payment_uid = ?", paymentUid).
		Update("status", "CANCELED").
		Error
}
