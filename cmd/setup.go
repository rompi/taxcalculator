package main

import (
	"database/sql"
	"fmt"

	"github.com/rompi/tax-calc/app/handler"
	"github.com/rompi/tax-calc/app/logic"
	"github.com/sirupsen/logrus"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

// logging
func setupLogger() *logrus.Logger {
	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}
	logger.Out = &lumberjack.Logger{
		Filename:   conf.GetString("logger.filename"), //log filename
		MaxSize:    conf.GetInt("logger.max_size"),    //megabytes
		MaxBackups: conf.GetInt("logger.max_backups"), //number
		MaxAge:     conf.GetInt("logger.max_age"),     //days
	}
	return logger
}

// database connection
func connectDB() (*sql.DB, error) {
	dbHost := conf.GetString("database.host")         //conf.GetString(env + ".database.HOST")
	dbPort := conf.GetInt("database.port")            //conf.GetInt(env + ".database.PORT")
	dbUser := conf.GetString("database.user")         //conf.GetString(env + ".database.USER")
	dbPassword := conf.GetString("database.password") //conf.GetString(env + ".database.PASSWORD")
	dbName := conf.GetString("database.name")         //conf.GetString(env + ".database.DATABASE_NAME")

	dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s connect_timeout=3 sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		fmt.Println("Can't connect to Database")
		return nil, err
	}

	return db, err
}

// setup request handler
func setupHandler(db *sql.DB, logger *logrus.Logger) handler.Handler {
	return handler.Handler{
		Logic: logic.Billing{
			Log:      logger,
			Database: &logic.DB{db},
		},
		Log: logger,
	}
}
