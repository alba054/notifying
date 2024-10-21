package notification

import (
	"fmt"

	"github.com/julienschmidt/httprouter"
)

type NotificationRouter struct {
	*httprouter.Router
	handler NotificationHandle
}

func NewRouter(router *httprouter.Router, handler NotificationHandle) *NotificationRouter {
	path := ""

	router.POST(fmt.Sprintf("%s/:topic", path), handler.PostMessageNotification)
	// stream using sse
	router.GET(fmt.Sprintf("%s/:topic", path), handler.GetMessageNotification)
	return &NotificationRouter{Router: router, handler: handler}
}
