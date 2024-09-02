package wsapp

import (
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/websocket"
	"log"
)

var upgrader = websocket.Upgrader{}

type Application struct {
}

func New() *Application {
	return &Application{}
}

func (a *Application) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("mt: %d, recv: %s", mt, message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
