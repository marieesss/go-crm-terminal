package domain

type Storer interface {
	Add(c *Contact) error
	GetAll() map[int]*Contact
	Delete(id int) error
}
