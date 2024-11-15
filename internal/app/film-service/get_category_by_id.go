package film_service

import (
	"context"
	film_service "github.com/Alex1472/ozon-film-service/pkg/film-service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i Implementation) GetCategoryById(ctx context.Context, req *film_service.GetCategoryByIdRequest) (*film_service.GetCategoryByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategoryById not implemented")
}
