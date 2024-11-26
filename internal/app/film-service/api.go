package film_service

import (
	"github.com/Alex1472/ozon-film-service/internal/service/film"
	desc "github.com/Alex1472/ozon-film-service/pkg/film-service"
)

type Implementation struct {
	desc.UnimplementedFilmServiceServer
	categoryService film.Service
}

func NewCategoryService(categoryService film.Service) desc.FilmServiceServer {
	return &Implementation{
		categoryService: categoryService,
	}
}
