package storage

import (
	"fmt"
	"sync"

	"gorm.io/driver/postgres"
	_ "gorm.io/driver/sqlite"

	"gorm.io/gorm"
)

var (
	sync_category sync.Once
	db_category   *gorm.DB
)

type categories struct {
	ID   uint   `gorm:"primarykey"`
	Name string `gorm:"index:idx_name,type:hash;unique"`
}

type postgresCategory struct{}

// Create implements categoryTable.
func (postgresCategory) Create(name string) (string, error) {
	category := categories{Name: name}
	db := db_category.Create(&category)
	if db.Error != nil {
		return "", db.Error
	}

	return fmt.Sprint(category.ID), nil

}

// Delete implements categoryTable.
func (postgresCategory) Delete(id string) error {
	category := categories{}
	db := db_category.Where("id = ?", id).Delete(&category)
	// if category not found
	if db.RowsAffected == 0 {
		return fmt.Errorf("Category not found with id %s", id)
	}
	if db.Error != nil {
		return db.Error
	}
	return nil
}

// Get implements categoryTable.
func (postgresCategory) Get(id string) (category categories, err error) {
	db := db_category.Where("id = ?", id).Find(&category)
	if db.Error != nil {
		return categories{}, db.Error
	}
	return category, nil
}

// GetAll implements categoryTable.
func (postgresCategory) GetAll() (categories []categories, err error) {
	db := db_category.Find(&categories)
	if db.Error != nil {
		return nil, db.Error
	}
	return categories, nil
}

// Update implements categoryTable.
func (postgresCategory) Update(id string, name string) error {
	category := categories{}
	db := db_category.Where("id = ?", id).Find(&category)
	// if category not found
	if db.RowsAffected == 0 {
		return fmt.Errorf("Category not found with id %s", id)
	}
	if db.Error != nil {
		return db.Error
	}
	category.Name = name
	db = db_category.Save(&category)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

// Connect implements categoryTable.
func (p postgresCategory) Connect() *gorm.DB {
	var err error

	sync_category.Do(func() {
		db_category, err = gorm.Open(postgres.Open(Dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
	})
	return db_category
}

// Migrate implements categoryTable.
func (p postgresCategory) Migrate() {
	db_web := p.Connect()
	db_web.AutoMigrate(&categories{})
}
