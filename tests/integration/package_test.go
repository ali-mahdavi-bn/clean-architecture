package integration

import (
	"clean-hex/internal/user_management"
	"clean-hex/pkg/framwork/infrastructure/databases"
	"clean-hex/pkg/framwork/infrastructure/redisx"
	"clean-hex/pkg/framwork/service_layer/cache"
	"clean-hex/pkg/framwork/service_layer/messagebus"
	"clean-hex/tests/mocks"
	testutil "clean-hex/tests/testutility"
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"os"
	"testing"
)

var Bus *messagebus.MessageBus
var RedisStore cache.Store

func TestMain(m *testing.M) {
	ctx := context.Background()
	testutil.IsIntegration()

	// Start the sqlite test server
	db, err := databases.New(databases.Config{
		Debug:        true,
		DBType:       "sqlite3",
		DSN:          "file::memory:?cache=shared",
		MaxLifetime:  1,
		MaxIdleTime:  1,
		MaxIdleConns: 1,
		MaxOpenConns: 1,
		TablePrefix:  "",
	})
	if err != nil {
		logrus.WithError(err).Fatal("Failed to start Sqlite test server")
		os.Exit(1)
	}
	redisConnection, err := redisx.NewRedisConnection(ctx, &redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		Username: "ali",
		DB:       0,
	})
	if err != nil {
		panic(err)
	}
	RedisStore = cache.NewRedisStore(redisConnection)

	// Migration
	userManagementModule := user_management.UserManagementModule{
		Ctx:         ctx,
		DB:          db,
		RouterGroup: nil,
	}
	err = userManagementModule.AutoMigration()
	if err != nil {
		logrus.WithError(err).Fatal("Failed to AutoMigration UserManagementModule")
		os.Exit(1)
	}
	Bus = mocks.SqliteUserManagementBootstrapTestApp(db)

	// Run the tests
	code := m.Run()

	// Teardown

	os.Exit(code)
}
