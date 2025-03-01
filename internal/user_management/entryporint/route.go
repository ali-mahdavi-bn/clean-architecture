package entryporint

import (
	"clean-hex/pkg/framwork/service_layer/cache"
	"clean-hex/pkg/framwork/service_layer/messagebus"
	"github.com/gin-gonic/gin"
)

var Bus *messagebus.MessageBus
var RedisStore cache.Store

// https://github.com/IBM/sarama
func RegisterV1Routers(bus *messagebus.MessageBus, routerGroup *gin.RouterGroup, redisStore cache.Store) {
	Bus = bus
	RedisStore = redisStore
	userRoute := routerGroup.Group("/v1/user")
	{
		userRoute.POST("", CreateUserHandler)
		userRoute.GET("/:userId", GetUserHandler)
		userRoute.GET("", ViewUserHandler)
		userRoute.PUT("/:userId", UpdateUserHandler)
		userRoute.DELETE("/:userId", DeleteUserHandler)
	}
	tradeRoute := userRoute.Group("/:userId/trade")
	{
		tradeRoute.POST("", CreateTradeHandler)
		//tradeRoute.GET("/:id", GetUserHandler)
		tradeRoute.GET("", ViewTradeHandler)
		//tradeRoute.PUT("/:id", UpdateUserHandler)
		//tradeRoute.DELETE("/:id", DeleteUserHandler)
	}

}
