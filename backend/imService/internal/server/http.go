package server

import (
	"github.com/cloudzenith/DouTok/backend/imService/internal/applications/recordapp"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/mux"
)

func NewHttpServer() *http.Server {
	var opts = []http.ServerOption{
		http.Timeout(-1),
		http.Address(":23000"),
	}

	record := recordapp.New()

	wsRouter := mux.NewRouter()
	wsRouter.HandleFunc("/ws", record.WSHandler)

	srv := http.NewServer(opts...)
	srv.HandlePrefix("/", wsRouter)
	return srv
}
