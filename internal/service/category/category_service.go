package category

import (
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

type Service struct {
	repository RepositoryInterface
}

type RepositoryInterface interface {
	FindAllCategories(_ context.Context) (Categories, error)
}

var ErrCategoryNotFound = errors.New("category not found")

func New(repository RepositoryInterface) Service {
	return Service{
		repository: repository,
	}
}

func (s Service) GetCategoryByID(ctx context.Context, id uint64) (*Category, error) {
	cats, err := s.repository.FindAllCategories(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "repository.FindAllCategories")
	}

	cat := cats.FilterById(id)
	if cat == nil {
		return nil, ErrCategoryNotFound
	}

	return cat, nil
}
