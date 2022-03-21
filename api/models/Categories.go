package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Categories struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"size:100;not null;unique" json:"name"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (c *Categories) Prepare() {
	c.ID = 0
	c.Name = html.EscapeString(strings.TrimSpace(c.Name))
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
}

func (c *Categories) Validate(action string) error {
	if c.Name == "" {
		return errors.New("Required Name")
	}
	return nil
}

func (c *Categories) SaveCategory(db *gorm.DB) (*Categories, error) {

	var err error
	err = db.Debug().Create(&c).Error
	if err != nil {
		return &Categories{}, err
	}
	return c, nil
}

func (c *Categories) FindAllCategories(db *gorm.DB) (*[]Categories, error) {
	var err error
	categories := []Categories{}
	err = db.Debug().Model(&Categories{}).Limit(100).Find(&categories).Error
	if err != nil {
		return &[]Categories{}, err
	}
	return &categories, err
}

func (c *Categories) FindCategoryByID(db *gorm.DB, uid uint32) (*Categories, error) {
	var err error
	err = db.Debug().Model(Categories{}).Where("id = ?", uid).Take(&c).Error
	if err != nil {
		return &Categories{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Categories{}, errors.New("User Not Found")
	}
	return c, err
}

func (c *Categories) UpdateCategory(db *gorm.DB, uid uint32) (*Categories, error) {

	db = db.Debug().Model(&Categories{}).Where("id = ?", uid).Take(&Categories{}).UpdateColumns(
		map[string]interface{}{
			"name":      c.Name,
			"update_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &Categories{}, db.Error
	}
	// This is the display the updated user
	err := db.Debug().Model(&Categories{}).Where("id = ?", uid).Take(&c).Error
	if err != nil {
		return &Categories{}, err
	}
	return c, nil
}

func (c *Categories) DeleteACategory(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&Categories{}).Where("id = ?", uid).Take(&Categories{}).Delete(&Categories{})

	if db.Error != nil {
		return 0, db.Error
	}

	return db.RowsAffected, nil
}
