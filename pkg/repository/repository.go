package repository

type ReposAuthorizer interface {
}

type Repository struct {
	ReposAuthorizer
}

func NewRepository() *Repository {
	return &Repository{}
}
