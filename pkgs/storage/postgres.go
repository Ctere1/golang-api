package storage

import (
	"sync"
	"time"

	"gorm.io/driver/postgres"
	_ "gorm.io/driver/sqlite"

	"gorm.io/gorm"
)

var (
	MaxOpenConns             int
	MaxIdleConns             int
	ConnMaxLifetimeInMinutes int
	database                 *gorm.DB
	once                     sync.Once
)

type products struct {
	ID          uint   `gorm:"primarykey"`
	Sku         string `gorm:"index:idx_sku,type:hash;unique"`
	Name        string
	Price       string `gorm:"default:null"`
	Description string `gorm:"default:null"`
}

type postgresProduct struct{}

// Create implements productTable.
func (postgresProduct) Create(sku string, name string, price string, description string) error {
	product := products{Sku: sku, Name: name, Price: price, Description: description}
	db := database.Create(&product)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

// Delete implements productTable.
func (postgresProduct) Delete(sku string) error {
	product := products{}
	db := database.Where("sku = ?", sku).Delete(&product)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

// Get implements productTable.
func (postgresProduct) Get(sku string) (name string, price string, description string, err error) {
	product := products{}
	db := database.Where("sku = ?", sku).First(&product)
	if db.Error != nil {
		return "", "", "", db.Error
	}
	return product.Name, product.Price, product.Description, nil
}

// Update implements productTable.
func (postgresProduct) Update(sku string, name string, price string, description string) error {
	// If sku is not found, returns error
	_, _, _, err := postgresProduct{}.Get(sku)
	if err != nil {
		return err
	}

	// Update product
	product := products{Sku: sku, Name: name, Price: price, Description: description}
	db := database.Where("sku = ?", sku).Updates(&product)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

// Get All products
func (postgresProduct) GetAll() (products []products, err error) {
	db := database.Find(&products)
	if db.Error != nil {
		return nil, db.Error
	}
	return products, nil
}

func (p postgresProduct) Connect() *gorm.DB {
	var err error
	once.Do(func() {

		database, err = gorm.Open(postgres.Open(Dsn), &gorm.Config{PrepareStmt: true})
		if err != nil {
			panic("Database connection failed")
		}

	})

	sqlDB, err := database.DB()
	sqlDB.SetMaxOpenConns(MaxOpenConns)
	sqlDB.SetMaxIdleConns(MaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Duration(ConnMaxLifetimeInMinutes) * time.Minute)

	return database
}

func (p postgresProduct) Migrate() {
	db := p.Connect()
	db.AutoMigrate(&products{})
}
