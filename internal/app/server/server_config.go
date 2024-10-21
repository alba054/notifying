package serverconfig

import (
	"alba054/kartjis-notify/shared"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Server interface {
	StartServer()
}

type ServerImpl struct {
	*http.ServeMux
	*httprouter.Router
}

func (server *ServerImpl) StartServer(host string, port int) {
	server.ServeMux.Handle("/", http.StripPrefix("", server.Router))

	fmt.Printf("Server listening on %s:%d\n", host, port)
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), server.ServeMux)
	shared.ThrowError(err)
}

func New(router *httprouter.Router) *ServerImpl {
	return &ServerImpl{
		ServeMux: http.NewServeMux(),
		Router:   router,
	}
}
