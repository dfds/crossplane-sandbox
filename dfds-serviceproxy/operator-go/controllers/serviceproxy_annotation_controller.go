package controllers

import (
	"context"
	"fmt"
	"github.com/go-logr/logr"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"strings"
)

type ServiceProxyAnnotationReconciler struct {
	client.Client
	Log logr.Logger
	Scheme *runtime.Scheme
}

func (r *ServiceProxyAnnotationReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = r.Log.WithValues("serviceproxy", req.NamespacedName)

	svc := &v1.Service{}
	err := r.Get(ctx, req.NamespacedName, svc)

	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			fmt.Println("Service  resource not found. Ignoring since object must be deleted")
			return ctrl.Result{}, nil
		}
		// Error reading the object - requeue the request.
		fmt.Println(err, "Failed to get Service")
		return ctrl.Result{}, err
	}

	fmt.Println("Service discovered: ", svc.Name)
	spAnnotations := r.GetServiceProxyAnnotations(svc)
	if len(spAnnotations) > 0 {
		fmt.Println("This service contains ServiceProxyAnnotations:")
		for key, val := range spAnnotations {
			fmt.Println(key, ":", val)
		}
	}

	return ctrl.Result{}, nil
}

func (r *ServiceProxyAnnotationReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1.Service{}).
		Complete(r)
}

func (r *ServiceProxyAnnotationReconciler) GetServiceProxyAnnotations(svc *v1.Service) map[string]string {
	payload := map[string]string{}

	for key, val := range svc.Annotations {
		if strings.Contains(key, "dfds.serviceproxy.kubernetes.io") {
			payload[key] = val
		}
	}

	return payload
}