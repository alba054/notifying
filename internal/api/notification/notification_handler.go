package notification

import (
	"alba054/kartjis-notify/internal/exception"
	"alba054/kartjis-notify/internal/model/request"
	webresponse "alba054/kartjis-notify/internal/model/web"
	"alba054/kartjis-notify/internal/service/notification"
	"alba054/kartjis-notify/shared"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

type NotificationHandler struct {
	notificationService notification.NotificationService
}

func NewHandler(notificationService notification.NotificationService) *NotificationHandler {
	return &NotificationHandler{notificationService: notificationService}
}

func (h *NotificationHandler) PostMessageNotification(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	topic := p.ByName("topic")
	notificationId, err := uuid.NewRandom()
	shared.ThrowError(err)

	payload := request.PostNotificationMessagePayload{}
	shared.ReadRequestBody(r.Body, &payload)

	payload.Id = notificationId.String()
	payload.Topic = topic

	err = h.notificationService.AddMessageToTopic(context.Background(), payload)
	shared.ThrowError(err)

	shared.WriteApiResponse(w, http.StatusCreated, webresponse.Success, "successfully publish message to topic")
}

func (h *NotificationHandler) GetMessageNotification(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	flusher, ok := w.(http.Flusher)

	if !ok {
		panic(exception.NewCustomHttpError(500, "internal server error"))
	}

	ticker := time.NewTicker(time.Second * 1)
	reqCtx := r.Context()
	topic := p.ByName("topic")
	subId := r.Header.Get("X-Sub-Id")
	if subId == "" {
		subId = r.RemoteAddr
	}

	err := h.notificationService.ActivateSubscriber(reqCtx, topic, subId)

	if err != nil {
		panic(err)
	}

	// Send an initial message immediately after connection is established
	fmt.Fprintf(w, "data: Connection established\n\n")
	flusher.Flush() // Ensure the message is sent right away

	for {
		select {
		case <-reqCtx.Done():
			fmt.Println("connection's closed")
			h.notificationService.DeactivateSubscriber(reqCtx, topic, subId)
			ticker.Stop()
			return
		case <-ticker.C:
			message, _ := h.notificationService.GetMessageNotification(reqCtx, topic, subId)
			if message == "" {
				continue
			}
			fmt.Fprintf(w, "data: %s\n", message)
			flusher.Flush()
			fmt.Printf("[X] %s Sent...\n", message)
		}
	}
}
