package misc

import (
	"context"
	"fmt"
	"github.com/coreos/go-oidc"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	v1Core "k8s.io/api/core/v1"
	"k8s.io/api/networking/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/json"
	"log"
	"net/http"
	"os"
	"strings"
)

var LISTEN_ADDRESS = os.Getenv("SERVICEPROXY_OPERATOR_LISTEN_ADDRESS")

func InitApi(store *InMemoryStore) {
	provider, err := oidc.NewProvider(context.Background(), "https://login.microsoftonline.com/73a99466-ad05-4221-9f90-e7142aa2f6c1/v2.0")
	if err != nil {
		log.Fatal(err)
	}

	authMiddleware := authenticationMiddleware{
		ClientID: "1fd40af5-f871-4502-834b-34c92ec9023f",
		Provider: provider,
	}


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
	r.Handle("/api/get-all", authMiddleware.Middleware(http.HandlerFunc(app.GetAll)))

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


type authenticationMiddleware struct {
	ClientID string
	Provider *oidc.Provider
}

func (amw *authenticationMiddleware) Middleware(next http.Handler) http.Handler {
	var verifier = amw.Provider.Verifier(&oidc.Config{ClientID: amw.ClientID})

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqToken := r.Header.Get("Authorization") //Authorization: Bearer a7ydfs87afasd8f990
		splitToken := strings.Split(reqToken, "Bearer")
		if len(splitToken) != 2 {
			http.Error(w, "Token doesn't seem right", http.StatusUnauthorized)
			return
		}

		reqToken = strings.TrimSpace(splitToken[1])

		idToken, err := verifier.Verify(r.Context(), reqToken)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Unable to verify token", http.StatusUnauthorized)
			return
		}

		var claims struct {
			Emails []string `json:"emails"`
		}
		if err := idToken.Claims(&claims); err != nil {
			fmt.Println(err)
			http.Error(w, "Unable to retrieve claims", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}