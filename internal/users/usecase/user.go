package usecase

import (
	"context"

	"github.com/hoag/go-social-feed/internal/models"
	"github.com/hoag/go-social-feed/internal/users"
	"github.com/hoag/go-social-feed/internal/users/repository"
)

func (uc impleUsecase) Create(ctx context.Context, input users.CreateInput) (models.User, error) {
	users, err := uc.repo.Create(ctx, repository.CreateOptions{
		UserName:     input.UserName,
		AvatarURL:    input.AvatarURL,
		Phone:        input.Phone,
		PasswordHash: input.PasswordHash,
		Birthday:     input.Birthday,
		Roles:        input.Roles,
		Permissions:  input.Permissions,
	})
	if err != nil {
		uc.l.Errorf(ctx, "users.usecase.Create.Create: %v", err)
		return models.User{}, err
	}

	return users, nil
}

func (uc impleUsecase) GetSessionUser(ctx context.Context, sc models.Scope) (models.User, error) {
	u, err := uc.repo.Detail(ctx, sc, sc.UserID)
	if err != nil {
		uc.l.Errorf(ctx, "users.usecase.GetSessionUser.Detail: %v", err)
		return models.User{}, nil
	}

	return u, nil
}

func (uc impleUsecase) Detail(ctx context.Context, sc models.Scope, id string) (models.User, error) {
	users, err := uc.repo.Detail(ctx, sc, id)
	if err != nil {
		uc.l.Errorf(ctx, "users.usecase.Detail.Detail: %v", err)
		return models.User{}, err
	}
	return users, nil
}

func (uc impleUsecase) List(ctx context.Context, sc models.Scope, input users.ListInput) ([]models.User, error) {
	userss, err := uc.repo.List(ctx, sc, repository.ListOptions{
		Filter: repository.Filter{
			ID:       input.ID,
			IDs:      input.IDs,
			UserName: input.UserName,
		},
	})
	if err != nil {
		uc.l.Errorf(ctx, "users.usecase.List.List: %v", err)
		return []models.User{}, err
	}
	return userss, nil
}

func (uc impleUsecase) Get(ctx context.Context, sc models.Scope, input users.GetInput) (users.GetOutput, error) {
	userss, paginator, err := uc.repo.Get(ctx, sc, repository.GetOptions{
		Filter: repository.Filter{
			ID:       input.ID,
			IDs:      input.IDs,
			UserName: input.UserName,
		},
		PagQuery: input.PagQuery,
	})
	if err != nil {
		uc.l.Errorf(ctx, "users.usecase.Get.Get: %v", err)
		return users.GetOutput{}, err
	}
	return users.GetOutput{
		Users:     userss,
		Paginator: paginator,
	}, nil
}

func (uc impleUsecase) Update(ctx context.Context, sc models.Scope, input users.UpdateInput) error {
	users, err := uc.repo.Detail(ctx, sc, input.ID)
	if err != nil {
		uc.l.Errorf(ctx, "users.usecase.Update.Detail: %v", err)
		return err
	}

	err = uc.repo.Update(ctx, sc, repository.UpdateOptions{
		User:      users,
		UserName:  input.UserName,
		AvatarURL: input.AvatarURL,
	})
	if err != nil {
		uc.l.Errorf(ctx, "users.usecase.Update.Update: %v", err)
		return err
	}
	return nil
}

func (uc impleUsecase) Delete(ctx context.Context, sc models.Scope, id string) error {
	err := uc.repo.Delete(ctx, sc, id)
	if err != nil {
		uc.l.Errorf(ctx, "users.usecase.Delete.Delete: %v", err)
		return err
	}
	return nil
}
