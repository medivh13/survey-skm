package responden_usecases

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : survey-skm
 */

import (
	"context"
	"log"
	dto "survey-skm/src/app/dtos/respondens"
	"survey-skm/src/domain/repositories"
)

type ResondenUsecaseInterface interface {
	CreateQuisionerData(ctx context.Context, data *dto.RespondenReqDTO) error
}

type respondenUseCase struct {
	RespondenRepo repositories.RespondenRepository
}

func NewRespondenUseCase(respondenRepo repositories.RespondenRepository) *respondenUseCase {
	return &respondenUseCase{
		RespondenRepo: respondenRepo,
	}
}

func (uc *respondenUseCase) CreateQuisionerData(ctx context.Context, data *dto.RespondenReqDTO) error {

	err := uc.RespondenRepo.CreateQuisionerData(ctx, data)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
