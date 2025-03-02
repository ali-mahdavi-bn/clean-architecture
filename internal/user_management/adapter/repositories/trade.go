package repositories

import (
	"clean-hex/internal/user_management/domain/entities"
	"clean-hex/pkg/framwork/adapter"
	"gorm.io/gorm"
)

type TradeRepository interface {
	adapter.BaseRepository[*entities.Trade]
}

type tradeGormRepository struct {
	adapter.BaseRepository[*entities.Trade]
	db *gorm.DB
}

func NewTradeGormRepository(db *gorm.DB) TradeRepository {
	return &tradeGormRepository{
		BaseRepository: adapter.NewGormRepository[*entities.Trade](db),
		db:             db,
	}
}
