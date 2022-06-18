package checkout_usecases

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : auth-skm
 */

import (
	"context"
	"errors"
	"testing"

	mockDTOitem "auth-skm/mocks/app/dtos/items"
	mockCheckoutRepo "auth-skm/mocks/domain/repository"

	dtoItem "auth-skm/src/app/dtos/items"
	models "auth-skm/src/infra/models"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockUsecase struct {
	mock.Mock
}

type UsecaseCheckoutTest struct {
	suite.Suite
	checkoutrepo *mockCheckoutRepo.MockItemsRepo

	usecase     CheckoutUsecaseInterface
	itemModels  []*models.Items
	dtoTest     *dtoItem.CheckoutReqDTO
	dtoTestFail *dtoItem.CheckoutReqDTO
	mockDTO     *mockDTOitem.MockItemDTO
}

func (suite *UsecaseCheckoutTest) SetupTest() {
	suite.checkoutrepo = new(mockCheckoutRepo.MockItemsRepo)
	suite.mockDTO = new(mockDTOitem.MockItemDTO)
	suite.usecase = NewCheckoutUseCase(suite.checkoutrepo)

	suite.itemModels = []*models.Items{
		&models.Items{
			ID:    1,
			SKU:   "Test",
			Name:  "Test",
			Price: 100,
			Qty:   10,
		},
	}

	suite.dtoTest = &dtoItem.CheckoutReqDTO{
		Data: []*dtoItem.SkuReqDTO{
			&dtoItem.SkuReqDTO{
				Sku: "Test",
				Qty: 3,
			},
		},
	}

}

func (uc *UsecaseCheckoutTest) TestCheckoutSuccess() {
	var sku = []string{"Test"}
	uc.checkoutrepo.Mock.On("ListItems", sku).Return(uc.itemModels, nil)
	_, err := uc.usecase.Checkout(context.Background(), uc.dtoTest)
	uc.Equal(nil, err)
}

func (uc *UsecaseCheckoutTest) TestCheckoutFailRetrieveItems() {
	var sku = []string{"Test"}
	uc.checkoutrepo.Mock.On("ListItems", sku).Return(nil, errors.New(mock.Anything))

	_, err := uc.usecase.Checkout(context.Background(), uc.dtoTest)
	uc.Error(errors.New(mock.Anything), err)
}

func TestUsecase(t *testing.T) {
	suite.Run(t, new(UsecaseCheckoutTest))
}
