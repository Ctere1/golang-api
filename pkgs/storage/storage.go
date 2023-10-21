package storage

import (
	"gorm.io/gorm"
)

type productTable interface {
	Create(sku string, name string, price string, description string, categoryId string) error
	Update(sku string, name string, price string, description string, categoryId string) error
	Delete(sku string) error
	Get(sku string) (name string, price string, description string, categoryId string, err error)
	GetAll() (products []products, err error)
	Connect() *gorm.DB
	Migrate()
}

type categoryTable interface {
	Create(name string) (string, error)
	Update(id string, name string) error
	Delete(id string) error
	Get(id string) (category categories, err error)
	GetAll() (categories []categories, err error)
	Connect() *gorm.DB
	Migrate()
}

var (
	Dsn                    string
	productDatabaseEngine  productTable
	categoryDatabaseEngine categoryTable
)

func SetProductsStorageType(database string) productTable {
	return postgresProduct{}
}

func SetCategoriesStorageType(database string) categoryTable {
	return postgresCategory{}
}

func Initialize() {
	productDatabaseEngine = SetProductsStorageType("postgres")
	productDatabaseEngine.Migrate()
	categoryDatabaseEngine = SetCategoriesStorageType("postgres")
	categoryDatabaseEngine.Migrate()
}

func CreateProduct(sku string, name string, price string, description string, categoryId string) error {
	return productDatabaseEngine.Create(sku, name, price, description, categoryId)
}

func UpdateProduct(sku string, name string, price string, description string, categoryId string) error {
	return productDatabaseEngine.Update(sku, name, price, description, categoryId)
}

func DeleteProduct(sku string) error {
	return productDatabaseEngine.Delete(sku)
}

func GetProduct(sku string) (name string, price string, description string, categoryId string, err error) {
	return productDatabaseEngine.Get(sku)
}

func GetAllProducts() (products []products, err error) {
	return productDatabaseEngine.GetAll()
}

func CreateCategory(name string) (string, error) {
	return categoryDatabaseEngine.Create(name)
}

func UpdateCategory(id string, name string) error {
	return categoryDatabaseEngine.Update(id, name)
}

func DeleteCategory(id string) error {
	return categoryDatabaseEngine.Delete(id)
}

func GetCategory(id string) (category categories, err error) {
	return categoryDatabaseEngine.Get(id)
}

func GetAllCategories() (categories []categories, err error) {
	return categoryDatabaseEngine.GetAll()
}
