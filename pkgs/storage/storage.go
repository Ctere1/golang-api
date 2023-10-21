package storage

import (
	"gorm.io/gorm"
)

type productTable interface {
	Create(sku string, name string, price string, description string) error
	Update(sku string, name string, price string, description string) error
	Delete(sku string) error
	Get(sku string) (name string, price string, description string, err error)
	GetAll() (products []products, err error)
	Connect() *gorm.DB
	Migrate()
}

var (
	Dsn                   string = "host=localhost user=postgres password=admin dbname=postgres port=5432 sslmode=disable TimeZone=Europe/Istanbul"
	productDatabaseEngine productTable
)

func SetStorageType(database string) productTable {
	return postgresProduct{}
}

func Initialize() {
	productDatabaseEngine = SetStorageType("postgres")
	productDatabaseEngine.Migrate()
}

func CreateProduct(sku string, name string, price string, description string) error {
	return productDatabaseEngine.Create(sku, name, price, description)
}

func UpdateProduct(sku string, name string, price string, description string) error {
	return productDatabaseEngine.Update(sku, name, price, description)
}

func DeleteProduct(sku string) error {
	return productDatabaseEngine.Delete(sku)
}

func GetProduct(sku string) (name string, price string, description string, err error) {
	return productDatabaseEngine.Get(sku)
}

func GetAllProducts() (products []products, err error) {
	return productDatabaseEngine.GetAll()
}
