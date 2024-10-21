package main

import (
	notificationapi "alba054/kartjis-notify/internal/api/notification"
	"alba054/kartjis-notify/internal/app/database"
	serverconfig "alba054/kartjis-notify/internal/app/server"
	"alba054/kartjis-notify/internal/config"
	errorhandler "alba054/kartjis-notify/internal/exception/handler"
	messagerepository "alba054/kartjis-notify/internal/repository/message"
	topicrepository "alba054/kartjis-notify/internal/repository/topic"
	notificationservice "alba054/kartjis-notify/internal/service/notification"
	"alba054/kartjis-notify/shared/constants"

	"github.com/julienschmidt/httprouter"
)

func main() {
	config := config.LoadConfig()
	router := httprouter.New()
	db := database.NewDB(config.DatabaseUrl)
	// * define repositories
	topicRepository := topicrepository.New(constants.TopicTableName)
	messageRepository := messagerepository.New(constants.MessageTableName)
	// * define services
	notificationService := notificationservice.New(topicRepository, messageRepository, db)
	// * define controllers
	notificationHandler := notificationapi.NewHandler(notificationService)
	// * define routers
	notificationapi.NewRouter(router, notificationHandler)
	// * define error handler
	errorhandler.UseErrorHandler(router)

	server := serverconfig.New(router)

	server.StartServer(config.Host, config.Port)
}
