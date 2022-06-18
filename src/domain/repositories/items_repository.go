package repositories

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : auth-skm
 */

import (
	"context"

	models "auth-skm/src/infra/models"
)

type ItemsRepository interface {
	ListItems(context context.Context, sku []string) ([]*models.Items, error)
}
