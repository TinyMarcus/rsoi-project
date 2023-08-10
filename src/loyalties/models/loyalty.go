package models

import (
	"loyalties/errors"
	"loyalties/objects"
	"loyalties/repositories"
)

type LoyaltyModel struct {
	repository      repositories.LoyaltyRepository
	bonusPercentage float32
}

func NewLoyaltyModel(repository repositories.LoyaltyRepository) *LoyaltyModel {
	return &LoyaltyModel{
		repository:      repository,
		bonusPercentage: 5,
	}
}

func (model *LoyaltyModel) GetLoyaltyForUser(username string) (*objects.Loyalty, error) {
	loyalty, err := model.repository.Find(username)
	if err != nil {
		return nil, err
	}

	return loyalty, nil
}

func (model *LoyaltyModel) IncreaseLoyalty(username string) error {
	loyalty, err := model.repository.Find(username)
	switch err {
	case nil:
		loyalty.ReservationCount += 1

		if loyalty.ReservationCount >= 10 && loyalty.ReservationCount < 20 {
			loyalty.Status = "SILVER"
			loyalty.Discount = 7
		}

		if loyalty.ReservationCount >= 20 {
			loyalty.Status = "GOLD"
			loyalty.Discount = 10
		}

		err = model.repository.Update(loyalty)
		if err != nil {
			return err
		}
	case errors.RecordNotFound:
		loyalty = &objects.Loyalty{
			Username:         username,
			ReservationCount: 0,
			Status:           "BRONZE",
			Discount:         5,
		}

		err = model.repository.Create(loyalty)
		if err != nil {
			return err
		}
	default:
		return err
	}

	return nil
}

func (model *LoyaltyModel) DecreaseLoyalty(username string) error {
	loyalty, err := model.repository.Find(username)
	if err != nil {
		return err
	}

	loyalty.ReservationCount -= 1

	if loyalty.ReservationCount < 10 {
		loyalty.Status = "BRONZE"
		loyalty.Discount = 5
	}

	if loyalty.ReservationCount >= 10 && loyalty.ReservationCount < 20 {
		loyalty.Status = "SILVER"
		loyalty.Discount = 7
	}

	if loyalty.ReservationCount >= 20 {
		loyalty.Status = "GOLD"
		loyalty.Discount = 10
	}

	err = model.repository.Update(loyalty)
	if err != nil {
		return err
	}

	return nil
}
