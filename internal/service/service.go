package service

import (
	"log"

	"github.com/chenchi1009/go-starter-kit/internal/repository"
	"github.com/chenchi1009/go-starter-kit/pkg/jwt"
)

type Service struct {
	logger *log.Logger
	jwt    *jwt.JWT
}

func NewService(
	tm repository.Transaction,
	logger *log.Logger,
	jwt *jwt.JWT,
) *Service {
	return &Service{
		logger: logger,
		jwt:    jwt,
	}
}
