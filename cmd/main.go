package main

import (
	"clean-hex/internal/user_management"
	"clean-hex/pkg/framwork/infrastructure/databases"
	"clean-hex/pkg/framwork/service_layer/types"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/spf13/cast"
	"os"
)

func main() {
	ctx := context.Background()
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	db, err := databases.NewDbConnection(databases.Config{
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

	g := gin.Default()
	api := g.Group("/api")

	// init modules
	RunModules(
		&user_management.UserManagementModule{
			Ctx:         ctx,
			DB:          db,
			RouterGroup: api,
		},
	)

	g.Run()
}

func RunModules(modules ...types.Modules) {
	for _, module := range modules {
		err := module.Init()
		if err != nil {
			panic(err)
		}
	}
}
