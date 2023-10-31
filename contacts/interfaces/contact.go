package interfaces

import "contacts/models"

type IContact interface {
	Add(*models.Contact) (*models.Contact, error)
	Get(string) (*models.Contact, error)
	Remove(string) (int64, error)
}
