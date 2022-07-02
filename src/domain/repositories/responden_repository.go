package repositories

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : survey-skm
 */

import (
	"context"

	dto "survey-skm/src/app/dtos/respondens"
)

type RespondenRepository interface {
	CreateQuisionerData(ctx context.Context, data *dto.RespondenReqDTO) error
}
