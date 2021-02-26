package controllers

import (
	"context"
	"fmt"
	"strings"

	"github.com/dfds/crossplane-sandbox/dfds-serviceproxy/operator-go/misc"
	"github.com/go-logr/logr"
	v1net "k8s.io/api/networking/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type IngressProxyAnnotationReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
	Store  *misc.Store
}

func (r *IngressProxyAnnotationReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = r.Log.WithValues("ingressproxy", req.NamespacedName)

	ing := &v1net.Ingress{}
	err := r.Get(ctx, req.NamespacedName, ing)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			fmt.Println("Ingress resource not found. Cleaning up DynamoDB entry")
			r.Store.RemoveService(misc.ServiceProxyItem{ObjectName: req.Name, ObjectNamespace: req.Namespace, ObjectKind: "Ingress"})
			return ctrl.Result{}, nil
		}
		// Error reading the object - requeue the request.
		fmt.Println(err, "Failed to get Ingress")
		return ctrl.Result{}, err
	}

	fmt.Println("Ingress discovered: ", ing.Name)
	spAnnotations := r.GetIngressProxyAnnotations(ing)
	if len(spAnnotations) > 0 {
		fmt.Println("This ingress contains ServiceProxyAnnotations:")
		for key, val := range spAnnotations {
			fmt.Println(key, ":", val)
		}
		r.Store.PutService(misc.ServiceProxyItem{ObjectName: req.Name, ObjectNamespace: req.Namespace, ObjectKind: "Ingress"})
	}

	return ctrl.Result{}, nil
}

func (r *IngressProxyAnnotationReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1net.Ingress{}).
		Complete(r)
}

func (r *IngressProxyAnnotationReconciler) GetIngressProxyAnnotations(svc *v1net.Ingress) map[string]string {
	payload := map[string]string{}

	for key, val := range svc.Annotations {
		if strings.Contains(key, "dfds.serviceproxy.kubernetes.io") {
			payload[key] = val
		}
	}

	return payload
}
