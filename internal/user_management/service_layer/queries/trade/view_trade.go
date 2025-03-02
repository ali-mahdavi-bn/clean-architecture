package trade

import (
	"clean-hex/internal"
	"clean-hex/internal/user_management/domain/entities"
	"clean-hex/pkg/framwork/errors"
	"clean-hex/pkg/framwork/service_layer/cache"
	"clean-hex/pkg/ginx"
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

func ViewTrade(ctx context.Context, userId uint, uow internal.UnitOfWorkImp, cache cache.Store, param *ginx.PaginationResult) (*ginx.ResponseResult, error) {
	result := &ginx.ResponseResult{
		Success: false,
	}
	key := cache.CreateKey("user", userId, "trade", "order", param.OrderBy.ToSQL(), "limit", param.Limit, "skip", param.Skip)
	err := cache.Cache(ctx, key, result, time.Second*2, func(ctx context.Context) (any, error) {
		fmt.Println(".......db hit.........")
		return uow.Do(ctx, func(ctx context.Context, tx *gorm.DB) (any, error) {
			trades := new([]entities.Trade)
			if uow.Trade().Model(ctx).Where("user_id = ?", userId).Limit(int(param.Limit)).Offset(int(param.Skip)).Order(param.OrderBy.ToSQL()).Find(trades).Count(&result.Total).Error != nil {
				return nil, errors.BadRequest("Operation.CanNot")
			}

			result.Pages, result.Page = ginx.CalculatePagination(result.Total, param.Limit, param.Skip)
			result.Data = trades
			result.Success = true
			return result, nil
		})
	})
	if err != nil {
		return nil, err
	}

	return result, nil

}
