package film_service

import (
	desc "github.com/Alex1472/ozon-film-service/pkg/film-service"
)

type Implementation struct {
	desc.UnimplementedFilmServiceServer
}

func NewCategoryService() desc.FilmServiceServer {
	return &Implementation{}
}
