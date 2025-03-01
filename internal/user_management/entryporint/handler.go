package entryporint

import (
	"clean-hex/internal/user_management/domain"
	"clean-hex/internal/user_management/service_layer/queries/trade"
	"clean-hex/internal/user_management/service_layer/queries/user"
	"clean-hex/pkg/ginx"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
)

func CreateUserHandler(c *gin.Context) {
	ctx := c.Request.Context()
	cmd := new(domain.CreateUserCommand)
	if err := ginx.ParseJSON(c, cmd); err != nil {
		ginx.ResError(c, err)
		return
	}

	result, err := Bus.Handle(ctx, cmd)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, result)
}

func GetUserHandler(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("userId")
	result, err := user.GetUser(ctx, Bus.Uow, cast.ToUint(id), RedisStore)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, result)
}

func ViewUserHandler(c *gin.Context) {
	ctx := c.Request.Context()
	params := new(ginx.PaginationResult)
	if err := ginx.ParsePaginationQueryParam(c, params); err != nil {
		ginx.ResError(c, err)
		return
	}

	result, err := user.ViewUser(ctx, Bus.Uow, RedisStore, params)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResPage(c, result, params)
}

func UpdateUserHandler(c *gin.Context) {
	id := c.Param("userId")
	ctx := c.Request.Context()
	cmd := new(domain.UpdateUserCommand)
	if err := ginx.ParseJSON(c, cmd); err != nil {
		ginx.ResError(c, err)
		return
	}

	cmd.UserId = cast.ToUint(id)
	result, err := Bus.Handle(ctx, cmd)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, result)
}

func DeleteUserHandler(c *gin.Context) {
	id := c.Param("userId")
	ctx := c.Request.Context()
	cmd := new(domain.DeleteUserCommand)
	cmd.UserId = cast.ToUint(id)
	result, err := Bus.Handle(ctx, cmd)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, result)
}

// trade
func CreateTradeHandler(c *gin.Context) {
	ctx := c.Request.Context()
	cmd := new(domain.CreateTradeCommand)
	if err := ginx.ParseJSON(c, cmd); err != nil {
		ginx.ResError(c, err)
		return
	}
	cmd.UserId = cast.ToUint(c.Param("userId"))

	result, err := Bus.Handle(ctx, cmd)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, result)
}

func ViewTradeHandler(c *gin.Context) {
	ctx := c.Request.Context()
	params := new(ginx.PaginationResult)
	if err := ginx.ParsePaginationQueryParam(c, params); err != nil {
		ginx.ResError(c, err)
		return
	}

	result, err := trade.ViewTrade(ctx, cast.ToUint(c.Param("userId")), Bus.Uow, RedisStore, params)
	if err != nil {
		fmt.Println(err)
		ginx.ResError(c, err)
		return
	}
	ginx.ResJSON(c, http.StatusOK, result)
}
