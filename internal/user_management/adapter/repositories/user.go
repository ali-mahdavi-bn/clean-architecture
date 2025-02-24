package repositories

import (
	"clean-hex/internal/user_management/domain/entities"
	"clean-hex/pkg/framwork/adapter"
	"context"
	"gorm.io/gorm"
)

type UserRepository interface {
	adapter.BaseRepository[*entities.User]
	FindByUserName(ctx context.Context, username string) (*entities.User, error)
	FindByUsernameExcludingID(ctx context.Context, username string, Id uint) (*entities.User, error)
}

type UserGormRepository struct {
	adapter.BaseRepository[*entities.User]
	db *gorm.DB
}

func (u *UserGormRepository) FindByUsernameExcludingID(ctx context.Context, username string, id uint) (*entities.User, error) {
	var user = new(entities.User)
	err := u.db.WithContext(ctx).Where("user_name = ? and id != ? and deleted_at is null", username, id).First(&user).Error
	return user, err
}

func (u *UserGormRepository) FindByUserName(ctx context.Context, username string) (*entities.User, error) {
	return u.FindByField(ctx, "user_name", username)

}

func NewUserGormRepository(db *gorm.DB) UserRepository {
	return &UserGormRepository{
		BaseRepository: adapter.NewGormRepository[*entities.User](db),
		db:             db,
	}
}
