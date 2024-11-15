package film_service

import (
	desc "github.com/Alex1472/ozon-film-service/pkg/film-service"
)

type Implementation struct {
	desc.UnimplementedSampleServiceServer
}

func NewSampleService() desc.SampleServiceServer {
	return &Implementation{}
}
