package main

import (
	notificationapi "alba054/kartjis-notify/internal/api/notification"
	"alba054/kartjis-notify/internal/app/database"
	serverconfig "alba054/kartjis-notify/internal/app/server"
	"alba054/kartjis-notify/internal/config"
	errorhandler "alba054/kartjis-notify/internal/exception/handler"
	"alba054/kartjis-notify/internal/model"
	messagerepository "alba054/kartjis-notify/internal/repository/message"
	topicrepository "alba054/kartjis-notify/internal/repository/topic"
	notificationservice "alba054/kartjis-notify/internal/service/notification"
	"alba054/kartjis-notify/shared/constants"
	"database/sql"
	"log"
	"os"

	"github.com/julienschmidt/httprouter"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("provide mode as argument [nodb | db]")
	}

	mode := os.Args[1]

	config := config.LoadConfig()
	var db *sql.DB
	var topicRepository topicrepository.TopicRepository
	var messageRepository messagerepository.MessageRepository

	if mode == "nodb" {
		db = nil
		// * define repositories
		topicRepository = topicrepository.NewNil()
		messageRepository = messagerepository.NewNil()
	} else if mode == "db" {
		db = database.NewDB(config.DatabaseUrl)
		// * define repositories
		topicRepository = topicrepository.New(constants.TopicTableName)
		messageRepository = messagerepository.New(constants.MessageTableName)
	} else {
		log.Fatal("mode should be [nodb | db]")
	}

	router := httprouter.New()
	// * local storage
	messageStorage := model.New()
	// * define services
	notificationService := notificationservice.New(topicRepository, messageRepository, db, messageStorage)
	// * define controllers
	notificationHandler := notificationapi.NewHandler(notificationService)
	// * define routers
	notificationapi.NewRouter(router, notificationHandler)
	// * define error handler
	errorhandler.UseErrorHandler(router)

	server := serverconfig.New(router)

	server.StartServer(config.Host, config.Port)
}
