package film_service

import (
	"context"
	"github.com/Alex1472/ozon-film-service/internal/service/film"
	film_service "github.com/Alex1472/ozon-film-service/pkg/film-service"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i Implementation) GetCategoryById(ctx context.Context, req *film_service.GetCategoryByIdRequest) (*film_service.GetCategoryByIdResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	cat, err := i.categoryService.GetCategoryByID(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, film.ErrCategoryNotFound) {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}
		return nil, errors.Wrap(err, "categoryService.GetCategoryByID")
	}

	resp := makeGetCategoryByIdResponse(cat)
	return resp, nil
}

func makeGetCategoryByIdResponse(cat *film.Category) *film_service.GetCategoryByIdResponse {
	pbCat := &film_service.Category{
		Id:   cat.ID,
		Name: cat.Name,
	}

	return &film_service.GetCategoryByIdResponse{
		Category: pbCat,
	}
}
