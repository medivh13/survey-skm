package usecases

import (
	checkoutUc "auth-skm/src/app/use_cases/checkout"
)

type AllUseCases struct {
	CheckoutUseCase checkoutUc.CheckoutUsecaseInterface
}
