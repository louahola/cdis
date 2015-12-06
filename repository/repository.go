package repository

type Repository interface {
	Save(interface{}) error
}
