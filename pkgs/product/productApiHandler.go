package product

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Product struct {
	Name        string
	Price       string
	Description string
	Sku         string
}

func ApiCreateProduct(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	// Create product
	result, err := createProduct(product)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			json.NewEncoder(w).Encode(map[string]string{"error": "Product already exists"})
		} else {
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		}
		return
	}

	// Return product as JSON
	json.NewEncoder(w).Encode(result)
}

func ApiGetProduct(w http.ResponseWriter, r *http.Request) {
	productSku := getSkuParamFromUrl(r)

	result, err := getProduct(productSku)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	// Return product as JSON
	json.NewEncoder(w).Encode(result)
}

func ApiGetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := getProducts()
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	if len(products) == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "No products found"})
		return
	}

	// Return products as JSON
	json.NewEncoder(w).Encode(products)

}

func ApiDeleteProduct(w http.ResponseWriter, r *http.Request) {
	productSku := getSkuParamFromUrl(r)
	err := deleteProduct(productSku)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	// Return success
	json.NewEncoder(w).Encode(map[string]string{"success": "Product deleted"})

}

func ApiUpdateProduct(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	// Update product
	result, err := updateProduct(product)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	// Return product as JSON
	json.NewEncoder(w).Encode(result)
}

func getSkuParamFromUrl(r *http.Request) string {
	params := mux.Vars(r)
	return params["sku"]
}
