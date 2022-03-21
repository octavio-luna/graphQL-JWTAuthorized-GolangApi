package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Product struct {
	ID       uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Name     string `gorm:"size:100;not null" json:"name"`
	Category uint32 `gorm:"ForeignKey:ID" json:"category"`
	//server.DB.Model(&Product{}).AddForeignKey("id", "categories(id)", "CASCADE", "SET NULL")
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *Product) Prepare(category uint32) {
	p.ID = 0
	p.Name = html.EscapeString(strings.TrimSpace(p.Name))
	p.Category = category
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *Product) Validate() error {
	if p.Name == "" {
		return errors.New("Name required")
	}
	return nil
}

func (p *Product) SaveProduct(db *gorm.DB) (*Product, error) {

	var err error
	err = db.Debug().Create(&p).Error
	if err != nil {
		return &Product{}, err
	}
	return p, nil
}

func (p *Product) FindAllProducts(db *gorm.DB) (*[]Product, error) {
	var err error
	products := []Product{}
	err = db.Debug().Model(&Product{}).Limit(100).Find(&products).Error
	if err != nil {
		return &[]Product{}, err
	}
	return &products, err
}

func (p *Product) FindProductByID(db *gorm.DB, uid uint32) (*Product, error) {
	var err error
	err = db.Debug().Model(Product{}).Where("id = ?", uid).Take(&p).Error
	if err != nil {
		return &Product{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Product{}, errors.New("Product Not Found")
	}
	return p, err
}

func (p *Product) UpdateAProduct(db *gorm.DB, uid uint32) (*Product, error) {

	db = db.Debug().Model(&Product{}).Where("id = ?", uid).Take(&Product{}).UpdateColumns(
		map[string]interface{}{
			"name":      p.Name,
			"category":  p.Category,
			"update_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &Product{}, db.Error
	}
	// This is the display the updated user
	err := db.Debug().Model(&Product{}).Where("id = ?", uid).Take(&p).Error
	if err != nil {
		return &Product{}, err
	}
	return p, nil
}

func (p *Product) DeleteAProduct(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&Product{}).Where("id = ?", uid).Take(&Product{}).Delete(&Product{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
