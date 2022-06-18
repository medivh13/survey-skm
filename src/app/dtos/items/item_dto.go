package items_dto

import validation "github.com/go-ozzo/ozzo-validation"

type ItemInterface interface {
	Validate() error
}

type CheckoutReqDTO struct {
	Data []*SkuReqDTO
}

type SkuReqDTO struct {
	Sku string `json:"sku"`
	Qty int64  `json:"qty"`
}

func (dto *CheckoutReqDTO) Validate() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.Data, validation.Required),
	); err != nil {
		return err
	}
	return nil
}
