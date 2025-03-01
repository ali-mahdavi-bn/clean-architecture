package workers

import (
	unit_of_work "clean-hex/internal"
	"clean-hex/internal/user_management/domain/entities"
	"clean-hex/pkg/framwork/errors"
	kafka "clean-hex/pkg/framwork/infrastructure/kafak"
	"clean-hex/pkg/framwork/service_layer/cache"
	"clean-hex/pkg/ginx"
	"context"
	"encoding/json"
	"github.com/IBM/sarama"
	"gorm.io/gorm"
	"log"
	"time"
)

func CacheTrade(uow unit_of_work.UnitOfWorkImp, cache cache.Store) {

	kafka.Service.ConsumeMessages(kafka.KAFKA_TOPIC_UPDATE_TRADE, func(pc sarama.PartitionConsumer) {
	kafkaBreak:
		for {
			select {
			case msg := <-pc.Messages():
				trade := new(entities.Trade)
				ctx := context.Background()
				if err := json.Unmarshal(msg.Value, trade); err != nil {
					break kafkaBreak
				}
				users := new([]entities.Trade)
				key := cache.CreateKey("user", trade.UserID, "trade", "order", "", "limit", 10, "skip", 0)
				result := &ginx.ResponseResult{
					Success: true,
				}
				err := cache.Cache(ctx, key, users, time.Minute, func(ctx context.Context) (any, error) {
					return uow.Do(ctx, func(ctx context.Context, tx *gorm.DB) (any, error) {
						user := new([]entities.Trade)
						if uow.Trade().Model(ctx).Where("user_id = ?", trade.UserID).Limit(10).Offset(0).Order("id").Find(user).Count(&result.Total).Error != nil {
							return nil, errors.BadRequest("Operation.CanNot")
						}
						result.Pages, result.Page = ginx.CalculatePagination(result.Total, 10, 0)
						result.Data = user

						return result, nil
					})
				})
				if err != nil {
					break kafkaBreak
				}
			case err := <-pc.Errors():
				log.Printf("Error: %v\n", err)
			}
		}
	})
}
