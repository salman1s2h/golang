package database

import (
	"demo/models"
	"fmt"

	"gorm.io/gorm"
)

type IContactDB interface {
	Create(*models.Contact) (*models.Contact, error)
	GetAll() ([]models.Contact, error)
	GetByID(id string) (*models.Contact, error)
	DeleteByID(id string) error
}

func NewContactDB(db *gorm.DB) IContactDB {
	return &ContactDB{db: db}
}

type ContactDB struct {
	db *gorm.DB
}

func (c *ContactDB) Create(contact *models.Contact) (*models.Contact, error) {
	if c.db != nil {
		c.db.AutoMigrate(&models.Contact{})
	}
	tx := c.db.Create((contact))
	if tx.Error != nil {
		return nil, tx.Error
	}
	return contact, nil
}

func (c *ContactDB) GetAll() ([]models.Contact, error) {
	var contacts []models.Contact
	tx := c.db.Find(&contacts)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return contacts, nil
}

func (c *ContactDB) DeleteByID(id string) error {
	tx := c.db.Delete(&models.Contact{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return fmt.Errorf("no data found with the following id:%v", id)
	}
	return nil
}

func (c *ContactDB) GetByID(id string) (*models.Contact, error) {
	contact := new(models.Contact)
	tx := c.db.First(contact, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return contact, nil
}
