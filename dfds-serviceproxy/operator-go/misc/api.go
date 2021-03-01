package misc

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"k8s.io/apimachinery/pkg/util/json"
	"log"
	"net/http"
	"os"
)

var LISTEN_ADDRESS = os.Getenv("SERVICEPROXY_OPERATOR_LISTEN_ADDRESS")

func InitApi(store *InMemoryStore) {

	var addr string
	if len(LISTEN_ADDRESS) > 0 {
		addr = LISTEN_ADDRESS
	} else {
		addr = "127.0.0.1"
	}
	addr = fmt.Sprintf("%s:8090", LISTEN_ADDRESS)

	r := mux.NewRouter()
	app := App {
		Router: r,
		Store: store,
	}
	r.Handle("/api/get-all", http.HandlerFunc(app.GetAll))

	fmt.Printf("HTTP server listening on %s\n", addr)
	if err := http.ListenAndServe(addr, handlers.LoggingHandler(os.Stdout, r)); err != nil {
		log.Fatal(err)
	}
}

type App struct {
	Router *mux.Router
	Store *InMemoryStore
}

func (a *App) GetAll(w http.ResponseWriter, r *http.Request) {
	payload, err := json.Marshal(a.Store)
	if err != nil {
		log.Println("Unable to serialise InMemoryStore")
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
	w.Write(payload)
}
