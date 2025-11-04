package domain

type Storer interface {
	Add(c *Contact) error
	GetAll() []*Contact
}
