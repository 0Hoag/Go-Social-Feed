package usecase

import (
	"context"

	"github.com/hoag/go-social-feed/internal/auth"
	"github.com/hoag/go-social-feed/internal/users"
	"github.com/hoag/go-social-feed/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

func (uc impleUsecase) Login(ctx context.Context, input auth.LoginInput) (auth.LoginResponse, error) {
	u, err := uc.userUC.GetOne(ctx, users.Filter{
		Phone: input.Phone,
	})
	if err != nil {
		uc.l.Errorf(ctx, "auth.usecase.user.Login.GetOne: %v", err)
		return auth.LoginResponse{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(input.Password)); err != nil {
		uc.l.Errorf(ctx, "auth.usecase.user.Login.CompareHashAndPassword: %v", err)
		return auth.LoginResponse{}, auth.ErrInvalidCreds
	}

	jwtManager := jwt.NewManager(uc.cfg.JWT.SecretKey)

	token, err := jwtManager.Generate(u.ID.Hex(), nil, nil)
	if err != nil {
		uc.l.Errorf(ctx, "auth.usecase.user.Login.Login: %v", err)
		return auth.LoginResponse{}, err
	}

	return auth.LoginResponse{
		Token: token,
	}, nil
}
