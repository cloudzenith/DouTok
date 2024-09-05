package imapp

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/gopkgs/components/consulx"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/transport/http"
	nethttp "net/http"
)

type Handler struct {
	consul *consul.Registry
}

func New() *Handler {
	registry := consul.New(consulx.GetClient(context.Background()))
	return &Handler{
		consul: registry,
	}
}

func (h *Handler) ImWebsocketHandler(w http.ResponseWriter, r *http.Request) {
	r.URL.Scheme = "ws"
	instances, err := h.consul.GetService(context.Background(), "im-service")
	if err != nil {
		panic(err)
	}

	if len(instances) == 0 {
		panic("no instances available")
	}

	endpoint := instances[0].Endpoints[0]
	r.URL.Host = endpoint
	r.RequestURI = ""
	nethttp.DefaultClient.Do(r)
}
