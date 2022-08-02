package inits

import (
	"github.com/Crossbell-Box/OperatorSync/app/server/config"
	"github.com/Crossbell-Box/OperatorSync/app/server/consts"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	"log"
	"os"
	"strings"
)

func Config() error {
	// Get from environment variables

	var exist bool // If non-exist use default

	if config.Config.DBConnString, exist = os.LookupEnv("DATABASE_CONNECTION_STRING"); !exist {
		config.Config.DBConnString = consts.CONFIG_DEFAULT_DATABASE_CONNECTION_STRING
	}
	if config.Config.RedisConnString, exist = os.LookupEnv("REDIS_CONNECTION_STRING"); !exist {
		config.Config.RedisConnString = consts.CONFIG_DEFAULT_REDIS_CONNECTION_STRING
	}
	if config.Config.MQConnString, exist = os.LookupEnv("MQ_CONNECTION_STRING"); !exist {
		config.Config.MQConnString = commonConsts.CONFIG_DEFAULT_MQ_CONNECTION_STRING
	}
	config.Config.IsMainServer = strings.Contains(strings.ToLower(os.Getenv("MAIN_SERVER")), "t")

	config.Config.DevelopmentMode = !strings.Contains(strings.ToLower(os.Getenv("MODE")), "prod")

	if config.Config.DevelopmentMode {
		log.Println("Configurations: ", config.Config)
	}

	return nil
}
