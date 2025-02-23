package repositories

import (
	"clean-hex/internal/user_management/domain/entities"
	"clean-hex/pkg/framwork/adapter"
	"context"
	"gorm.io/gorm"
)

type UserRepository interface {
	adapter.BaseRepository[*entities.User]
	ByUserName(ctx context.Context, username string) (*entities.User, error)
}

type UserGormRepository struct {
	adapter.BaseRepository[*entities.User]
}

func (u *UserGormRepository) ByUserName(ctx context.Context, username string) (*entities.User, error) {
	return u.ByField(ctx, "user_name", username)
}

func NewUserGormRepository(db *gorm.DB) UserRepository {
	return &UserGormRepository{
		BaseRepository: adapter.NewGormRepository[*entities.User](db),
	}
}
