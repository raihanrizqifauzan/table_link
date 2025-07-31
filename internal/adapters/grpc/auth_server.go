package grpc

import (
	"table_link/internal/application/auth"
)

type AuthServer struct {
	authService auth.Service
}
