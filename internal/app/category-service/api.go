package category_service

import (
	"github.com/Alex1472/ozon-category-service/internal/service/category"
	desc "github.com/Alex1472/ozon-category-service/pkg/category-service"
)

type Implementation struct {
	desc.UnimplementedCategoryServiceServer
	categoryService category.Service
}

func NewCategoryService(categoryService category.Service) desc.CategoryServiceServer {
	return &Implementation{
		categoryService: categoryService,
	}
}
