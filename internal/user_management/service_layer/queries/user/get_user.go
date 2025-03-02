package user

import (
	"clean-hex/internal"
	"clean-hex/internal/user_management/domain/entities"
	"clean-hex/pkg/framwork/errors"
	"clean-hex/pkg/framwork/service_layer/cache"
	"context"
	"gorm.io/gorm"
	"time"
)

func GetUser(ctx context.Context, uow internal.UnitOfWorkImp, id uint, cache cache.Store) (*entities.User, error) {
	user := new(entities.User)
	key := cache.CreateKey("user", id)
	err := cache.Cache(ctx, key, user, time.Second*5, func(ctx context.Context) (any, error) {
		return uow.Do(ctx, func(ctx context.Context, tx *gorm.DB) (any, error) {
			if uow.User().Model(ctx).Preload("Trades").First(user, id).Error != nil {
				return nil, errors.BadRequest("Operation.CanNot")
			}
			return user, nil
		})

	})
	return user, err

}
