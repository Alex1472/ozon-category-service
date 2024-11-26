package category_service

import (
	"context"
	"github.com/Alex1472/ozon-category-service/internal/service/category"
	category_service "github.com/Alex1472/ozon-category-service/pkg/category-service"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i Implementation) GetCategoryById(ctx context.Context, req *category_service.GetCategoryByIdRequest) (*category_service.GetCategoryByIdResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	cat, err := i.categoryService.GetCategoryByID(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, category.ErrCategoryNotFound) {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}
		return nil, errors.Wrap(err, "categoryService.GetCategoryByID")
	}

	resp := makeGetCategoryByIdResponse(cat)
	return resp, nil
}

func makeGetCategoryByIdResponse(cat *category.Category) *category_service.GetCategoryByIdResponse {
	pbCat := &category_service.Category{
		Id:   cat.ID,
		Name: cat.Name,
	}

	return &category_service.GetCategoryByIdResponse{
		Category: pbCat,
	}
}
