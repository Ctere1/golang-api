package api

import (
	"log"
	"net/http"

	"github.com/Ctere1/golang-api/pkgs/category"
	"github.com/Ctere1/golang-api/pkgs/product"
	"github.com/gorilla/mux"
)

var (
	ListenAddress = "localhost"
	ListenPort    = ":443"
)

func StartRouter() {
	router := mux.NewRouter()

	// Product routes
	router.Handle("/api/v1/product", restApiAuth(product.ApiCreateProduct)).Methods(http.MethodPost)
	router.Handle("/api/v1/product", restApiAuth(product.ApiGetProducts)).Methods(http.MethodGet)
	router.Handle("/api/v1/product/{sku}", restApiAuth(product.ApiGetProduct)).Methods(http.MethodGet)
	router.Handle("/api/v1/product/{sku}", restApiAuth(product.ApiDeleteProduct)).Methods(http.MethodDelete)
	router.Handle("/api/v1/product", restApiAuth(product.ApiUpdateProduct)).Methods(http.MethodPut)

	// Category routes
	router.Handle("/api/v1/category", restApiAuth(category.ApiCreateCategory)).Methods(http.MethodPost)
	router.Handle("/api/v1/category", restApiAuth(category.ApiGetCategories)).Methods(http.MethodGet)
	router.Handle("/api/v1/category/{id}", restApiAuth(category.ApiGetCategory)).Methods(http.MethodGet)
	router.Handle("/api/v1/category/{id}", restApiAuth(category.ApiDeleteCategory)).Methods(http.MethodDelete)
	router.Handle("/api/v1/category", restApiAuth(category.ApiUpdateCategory)).Methods(http.MethodPut)

	//Not found handler
	router.NotFoundHandler = http.HandlerFunc(notFound)

	log.Println("Starting server on", ListenAddress+ListenPort)

	http.ListenAndServe(ListenAddress+ListenPort, router)

}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Not Found"))
}
