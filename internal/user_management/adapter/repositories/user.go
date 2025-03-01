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

type userGormRepository struct {
	adapter.BaseRepository[*entities.User]
	db *gorm.DB
}

func NewUserGormRepository(db *gorm.DB) UserRepository {
	return &userGormRepository{
		BaseRepository: adapter.NewGormRepository[*entities.User](db),
		db:             db,
	}
}

func (u *userGormRepository) FindByUsernameExcludingID(ctx context.Context, username string, id uint) (*entities.User, error) {
	var user = new(entities.User)
	err := u.Model(ctx).Where("user_name = ? and id != ? and deleted_at is null", username, id).First(&user).Error
	return user, err
}

func (u *userGormRepository) FindByUserName(ctx context.Context, username string) (*entities.User, error) {
	return u.FindByField(ctx, "user_name", username)

}
