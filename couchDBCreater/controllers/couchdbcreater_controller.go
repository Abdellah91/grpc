/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	netHttp "net/http"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	dbv1 "github.com/Abdellah91/grpc/couchDBCreater/api/v1"
)

// CouchDBCreaterReconciler reconciles a CouchDBCreater object
type CouchDBCreaterReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=db.couchdb,resources=couchdbcreaters,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=db.couchdb,resources=couchdbcreaters/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=db.couchdb,resources=couchdbcreaters/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the CouchDBCreater object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.0/pkg/reconcile
func (r *CouchDBCreaterReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	l := log.FromContext(ctx)

	database := &dbv1.CouchDBCreater{}
	secret := &corev1.Secret{}
	r.Get(ctx, types.NamespacedName{Name: req.Name, Namespace: req.Namespace}, database)
	l.Info("Reconcile ", "database.Spec.Password.Secret: ", database.Spec.Password.Secret)

	r.Get(ctx, types.NamespacedName{Name: database.Spec.Password.Secret, Namespace: req.Namespace}, secret)
	url, DBName, userName := database.Spec.Url, database.Spec.Database, database.Spec.UserName
	DBPass := string(secret.Data[database.Spec.Password.Key])

	urlDBEndpoint := "http://" + userName + ":" + DBPass + "@" + url + "/" + DBName

	reqst, err := netHttp.NewRequest("PUT", urlDBEndpoint, nil)
	if err != nil {
		panic(err)
	}
	resp, err := netHttp.DefaultClient.Do(reqst)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *CouchDBCreaterReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&dbv1.CouchDBCreater{}).
		Complete(r)
}
