package main

import (
	"clean-hex/internal/user_management"
	"clean-hex/pkg/framwork/infrastructure/databases"
	"clean-hex/pkg/framwork/infrastructure/redisx"
	"clean-hex/pkg/framwork/service_layer/cache"
	"clean-hex/pkg/framwork/service_layer/types"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/cast"
	"os"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}
	ctx := context.Background()

	db, err := databases.New(databases.Config{
		Debug:        os.Getenv("DEBUG") == "true",
		DBType:       os.Getenv("DB_TYPE"),
		DSN:          os.Getenv("DATABASE_URL"),
		MaxLifetime:  cast.ToInt(os.Getenv("MAX_LIFETIME")),
		MaxIdleTime:  cast.ToInt(os.Getenv("MAX_IDLETIME")),
		MaxIdleConns: cast.ToInt(os.Getenv("MAX_IDLE_CONNS")),
		MaxOpenConns: cast.ToInt(os.Getenv("MAX_OPEN_CONNS")),
		TablePrefix:  os.Getenv("TABLE_PREFIX"),
	})

	if err != nil {
		panic(err)
	}
	redisConnection, err := redisx.NewRedisConnection(ctx, &redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		Username: os.Getenv("REDIS_USER"),
		DB:       0,
	})
	if err != nil {

		panic(err)
	}
	redisStore := cache.NewRedisStore(redisConnection)

	ginServer := gin.Default()
	api := ginServer.Group("/api")

	// init modules
	RunModules(
		&user_management.UserManagementModule{
			Ctx:         ctx,
			DB:          db,
			RedisStore:  redisStore,
			RouterGroup: api,
		},
	)

	err = ginServer.Run()
	if err != nil {

		panic(err)
	}

}

func RunModules(modules ...types.Modules) {
	for _, module := range modules {
		err := module.Init()
		if err != nil {
			panic(err)
		}
	}
}
