package film_repository

import (
	"context"
	"github.com/Alex1472/ozon-film-service/internal/service/film"
)

var categories = film.Categories{
	{
		ID:   1,
		Name: "Toys",
	},
	{
		ID:   2,
		Name: "Laptops",
	},
	{
		ID:   3,
		Name: "Auto",
	},
}

type Repository struct {
}

func New() *Repository {
	return &Repository{}
}

func (r *Repository) FindAllCategories(_ context.Context) (film.Categories, error) {
	return categories, nil
}
