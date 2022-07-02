package repositories

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : survey-skm
 */

import (
	"context"
	dtoLogin "survey-skm/src/app/dtos/login"
	dto "survey-skm/src/app/dtos/users"
	models "survey-skm/src/infra/models"
)

type UsersRepository interface {
	Register(ctx context.Context, data *dto.UserReqDTO) (*models.Users, error)
	GetUser(ctx context.Context, data *dtoLogin.LoginReqDTO) (*models.Users, error)
	UpdateToken(ctx context.Context, id int64, token string) error
}
