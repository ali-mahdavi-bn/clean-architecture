package entryporint

import (
	"clean-hex/internal/user_management/domain"
	queries "clean-hex/internal/user_management/service_layer/queries/user"
	"clean-hex/pkg/framwork/service_layer/messagebus"
	"clean-hex/pkg/ginx"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func RegisterV1Routers(bus *messagebus.MessageBus, g *gin.RouterGroup) {
	r := g.Group("/v1/user")

	r.POST("", func(c *gin.Context) {
		ctx := c.Request.Context()
		item := new(domain.CreateUserCommand)
		if err := ginx.ParseJSON(c, item); err != nil {
			ginx.ResError(c, err)
			return
		}
		result, err := bus.Handle(ctx, *item)
		if err != nil {
			ginx.ResError(c, err)
			return
		}
		ginx.ResSuccess(c, result)
	})

	r.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		result, err := queries.ViewGetUser(bus.DB, cast.ToUint(id))

		if err != nil {
			ginx.ResError(c, err)
			return
		}
		ginx.ResSuccess(c, result)
	})

}
