package notification

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type NotificationHandle interface {
	PostMessageNotification(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
	GetMessageNotification(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
}
