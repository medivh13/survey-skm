package main

import (
	"context"
	"database/sql"

	usecases "auth-skm/src/app/use_cases"

	"auth-skm/src/infra/config"
	postgres "auth-skm/src/infra/persistence/postgress"

	"auth-skm/src/interface/rest"

	ms_log "auth-skm/src/infra/log"

	loginUc "auth-skm/src/app/use_cases/login"
	userUc "auth-skm/src/app/use_cases/user"

	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
)

//reupdate by Jody 24 Jan 2022
func main() {
	// init context
	ctx := context.Background()

	// read the server environment variables
	conf := config.Make()

	// check is in production mode
	isProd := false
	if conf.App.Environment == "PRODUCTION" {
		isProd = true
	}

	// logger setup
	m := make(map[string]interface{})
	m["env"] = conf.App.Environment
	m["service"] = conf.App.Name
	logger := ms_log.NewLogInstance(
		ms_log.LogName(conf.Log.Name),
		ms_log.IsProduction(isProd),
		ms_log.LogAdditionalFields(m))

	postgresdb := postgres.New(conf.SqlDb, logger)

	// gracefully close connection to persistence storage
	defer func(l *logrus.Logger, sqlDB *sql.DB, dbName string) {
		err := sqlDB.Close()
		if err != nil {
			l.Errorf("error closing sql database %s: %s", dbName, err)
		} else {
			l.Printf("sql database %s successfuly closed.", dbName)
		}
	}(logger, postgresdb.SqlDB, postgresdb.DB.Name())

	userRepository := postgres.NewUsersRepository(postgresdb.DB)

	httpServer, err := rest.New(
		conf.Http,
		isProd,
		logger,
		usecases.AllUseCases{

			UserUseCase:  userUc.NewUserUseCase(userRepository),
			LoginUseCase: loginUc.NewUserUseCase(userRepository),
		},
	)
	if err != nil {
		panic(err)
	}
	httpServer.Start(ctx)

}
