package misc

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	v1Core "k8s.io/api/core/v1"
	"k8s.io/api/networking/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	payload, err := json.Marshal(StoreToGetAllResponse(a.Store))
	if err != nil {
		log.Println("Unable to serialise InMemoryStore")
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
	w.Write(payload)
}

type GetAllResponse struct {
	Ingress []IngressDto
	Service []ServiceDto
}

func StoreToGetAllResponse(store *InMemoryStore) GetAllResponse{
	payload := GetAllResponse{
		Ingress: []IngressDto{},
		Service: []ServiceDto{},
	}

	for _, v := range store.Services {
		val := ServiceDto{
			Kind: v.Kind,
			ApiVersion: v.APIVersion,
			Metadata: v.ObjectMeta,
			Spec: v.Spec,
			Status: v.Status,
		}

		val.Metadata.ManagedFields = []v1.ManagedFieldsEntry{}
		delete(val.Metadata.Annotations, "kubectl.kubernetes.io/last-applied-configuration")

		payload.Service = append(payload.Service, val)
	}

	for _, v := range store.Ingresses {
		val := IngressDto{
			Kind: v.Kind,
			ApiVersion: v.APIVersion,
			Metadata: v.ObjectMeta,
			Spec: v.Spec,
			Status: v.Status,
		}

		val.Metadata.ManagedFields = []v1.ManagedFieldsEntry{}
		delete(val.Metadata.Annotations, "kubectl.kubernetes.io/last-applied-configuration")

		payload.Ingress = append(payload.Ingress, val)
	}

	return payload
}

type IngressDto struct {
	Kind string
	ApiVersion string
	Metadata v1.ObjectMeta
	Spec v1beta1.IngressSpec
	Status v1beta1.IngressStatus
}

type ServiceDto struct {
	Kind string
	ApiVersion string
	Metadata v1.ObjectMeta
	Spec v1Core.ServiceSpec
	Status v1Core.ServiceStatus
}