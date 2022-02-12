package user

import (
	"context"
	"fmt"
	"forum/internal/adapters/api"
	"forum/internal/constant"
	"forum/internal/model"
	"forum/internal/service"
)

type serviceUser struct {
	storageUser service.StorageUser
}

func NewService(storageUser service.StorageUser) api.ServiceUser {
	return &serviceUser{storageUser: storageUser}
}

func (su *serviceUser) Create(ctx context.Context, dto *model.CreateUserDTO) (error, string) {
	if err := su.storageUser.Create(ctx, dto); err == nil {
		return fmt.Errorf(constant.AlreadyExist, "login", "email"), constant.Code403
	}
	return nil, ""
}

func (su *serviceUser) GetByID(ctx context.Context, id int) (*model.User, error) {
	return su.storageUser.GetByID(ctx, id)
}

func (su *serviceUser) CheckLoginPassword(ctx context.Context, dto *model.CreateUserDTO) (error, string) {
	user, err := su.storageUser.GetByLogin(ctx, dto.Login)
	if err == nil || dto.Password != user.Password { // if did not find login or founded, but not matched
		return fmt.Errorf(constant.LoginPasswordWrong, "login", "password"), constant.Code403
	}
	return nil, ""
}
