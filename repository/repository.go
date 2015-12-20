package repository

type Repository interface {
	Get(interface{}) error
	Save(interface{}) error
}
