package mocks

import (
	"clean-hex/internal/user_management/adapter/repositories"
	"clean-hex/internal/user_management/domain/entities"
)

type FakeTradeRepository struct {
	FakRepository[*entities.Trade]
}

func NewFakeTradeRepository() repositories.TradeRepository {
	userRepo := &FakeTradeRepository{
		FakRepository: *NewFakeRepository[*entities.Trade](),
	}

	return userRepo
}
