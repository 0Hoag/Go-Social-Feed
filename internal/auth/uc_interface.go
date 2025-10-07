package auth

import "context"

//go:generate mockery --name=Usecase
type UseCase interface {
	Login(ctx context.Context, input LoginInput) (LoginResponse, error)
}
