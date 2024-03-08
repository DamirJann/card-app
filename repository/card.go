package repository

type Card struct {
	ID     int    `gorm:"id,primaryKey"`
	Name   string `gorm:"name"`
	Answer string `gorm:"answer"`
}

type CardFilter struct {
	ID   *int
	Name *string
}

type Repository interface {
	Find() ([]Card, error)
	Delete(filter CardFilter) ([]Card, error)
	Create(name, answer string) error
}

type NullRepository struct{}

func (n NullRepository) Find() ([]Card, error) {
	//TODO implement me
	panic("implement me")
}

func (n NullRepository) Delete(filter CardFilter) ([]Card, error) {
	//TODO implement me
	panic("implement me")
}

func (n NullRepository) Create(name, answer string) error {
	//TODO implement me
	panic("implement me")
}
