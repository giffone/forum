package session

import (
	"context"
	"forum/internal/adapters/api"
	"forum/internal/constant"
	"forum/internal/model"
	"forum/internal/service"
)

type serviceSession struct {
	storageSession service.StorageSession
}

func NewService(storageSession service.StorageSession) api.ServiceSession {
	return &serviceSession{storageSession: storageSession}
}

func (ss *serviceSession) Create(ctx context.Context, dto *model.CreateSessionDTO) (error, string) {
	// create session
	if err := ss.storageSession.Create(ctx, dto); err == nil {
		return nil, ""
	}
	// if session already exist, delete it
	if err := ss.DeleteByLogin(ctx, dto.Login); err != nil {
		return err, constant.Code500
	}
	// create new session after delete
	if err := ss.storageSession.Create(ctx, dto); err != nil {
		return err, constant.Code500
	}
	return nil, ""
}

func (ss *serviceSession) GetByLogin(ctx context.Context, login string) (*model.Session, error) {

	return nil, nil
}

func (ss *serviceSession) DeleteByLogin(ctx context.Context, login string) error {
	return ss.storageSession.DeleteByLogin(ctx, login)
}

func (ss *serviceSession) DeleteExpired(ctx context.Context) error {

	return nil
}
