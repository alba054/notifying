package notification

import (
	"alba054/kartjis-notify/internal/model/request"
	webresponse "alba054/kartjis-notify/internal/model/web"
	"alba054/kartjis-notify/internal/service/notification"
	"alba054/kartjis-notify/shared"
	"context"
	"net/http"

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
	// w.Header().Set("Content-Type", "text/event-stream")
	// w.Header().Set("Cache-Control", "no-cache")
	// w.Header().Set("Connection", "keep-alive")

	// flusher, ok := w.(http.Flusher)

	// if !ok {
	// 	panic(exception.NewCustomHttpError(500, "internal server error"))
	// }

	// topic := p.ByName("topic")

	// for {
	// 	order, err := c.orderService.GetOrderByOrderId(context.Background(), orderId)
	// 	helper.ThrowToPanicHandler(err)

	// 	data := make(map[string]string)
	// 	data["orderId"] = *helper.HandleNullableStringColumn(order.OrderId)
	// 	data["status"] = string(order.OrderStatus)
	// 	streamData, err := json.Marshal(data)
	// 	helper.ThrowToPanicHandler(err)

	// 	fmt.Fprintf(w, "data: %s\n\n", streamData)
	// 	flusher.Flush() // Flush the data to the client immediately

	// 	// Simulate a delay
	// 	time.Sleep(2 * time.Second)
	// }
}
