package service

import (
	"context"
	"log"

	"github.com/ersa97/new-grpc/server/data"
)

type Server struct {
}

func (s *Server) GetAll(ctx context.Context, req *data.AllRequest) (*data.AllResponse, error) {
	log.Println("test get All")

	return &data.AllResponse{
		Message: req.Message,
		Data:    req.Data,
	}, nil
}
