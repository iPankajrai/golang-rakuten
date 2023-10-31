package database

import (
	"contacts/models"

	"gorm.io/gorm"
)

type ContactDB struct {
	DB *gorm.DB // promoted field
}

func (c *ContactDB) Add(contact *models.Contact) (*models.Contact, error) {
	if err := c.DB.AutoMigrate(&models.Contact{}); err != nil {
		return nil, err
	}
	tx := c.DB.Create(contact)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return contact, nil
}

func (c *ContactDB) Get(id string) (*models.Contact, error) {
	contact := new(models.Contact)
	tx := c.DB.First(contact, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return contact, nil
}

func (c *ContactDB) Remove(id string) (int64, error) {
	tx := c.DB.Delete(&models.Contact{}, id)

	if tx.Error != nil {
		return -1, tx.Error
	}
	return tx.RowsAffected, nil
}
