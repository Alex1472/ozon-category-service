package film_service

import (
	"context"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	desc "github.com/Alex1472/ozon-film-service/pkg/film-service"
)

func (i *Implementation) SampleMethodV1(
	ctx context.Context,
	req *desc.SampleMethodV1Request,
) (*desc.SampleMethodV1Response, error) {
	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("SampleMethodV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &desc.SampleMethodV1Response{
		Value: &desc.Template{
			Id:   req.GetId(),
			Name: "Sample",
		},
	}, nil
}
