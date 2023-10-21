package product

import (
	"github.com/Ctere1/golang-api/pkgs/storage"
)

func createProduct(product Product) (*Product, error) {
	// Create product in database
	err := storage.CreateProduct(product.Sku, product.Name, product.Price, product.Description)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func getProduct(productSku string) (*Product, error) {
	// Get product from database
	name, price, description, err := storage.GetProduct(productSku)
	if err != nil {
		return nil, err
	}

	// Create product object
	product := Product{
		Name:        name,
		Price:       price,
		Description: description,
		Sku:         productSku,
	}

	return &product, nil
}

func getProducts() ([]Product, error) {
	// Get products from database
	products, err := storage.GetAllProducts()
	if err != nil {
		return nil, err
	}

	// Convert []storage.Product to []Product
	var result []Product
	for _, p := range products {
		result = append(result, Product{
			Name:        p.Name,
			Price:       p.Price,
			Description: p.Description,
			Sku:         p.Sku,
		})
	}

	return result, nil
}

func deleteProduct(productSku string) error {
	// Delete product from database
	err := storage.DeleteProduct(productSku)
	if err != nil {
		return err
	}

	return nil
}

func updateProduct(product Product) (*Product, error) {
	// Update product in database
	err := storage.UpdateProduct(product.Sku, product.Name, product.Price, product.Description)
	if err != nil {
		return nil, err
	}

	return &product, nil
}
