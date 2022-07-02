package respondens_dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type RespondenInterface interface {
	Validate() error
}

type RespondenReqDTO struct {
	Name         string                   `json:"name"`
	Email        string                   `json:"email"`
	Umur         int64                    `json:"umur"`
	PekerjaanID  int64                    `json:"pekerjaan_id"`
	PendidikanID int64                    `json:"pendidikan_id"`
	LayananID    int64                    `json:"layanan_id"`
	Quisioner    []*RespondenDetailReqDTO `json:"quisioner"`
}

type RespondenDetailReqDTO struct {
	PertanyaanID int64 `json:"pertanyaan_id"`
	JawabanID    int64 `json:"jawaban_id"`
	Nilai        int64 `json:"nilai"`
}

func (dto *RespondenReqDTO) Validate() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.Name, validation.Required),
		validation.Field(&dto.Email, validation.Required, is.Email),
		validation.Field(&dto.Umur, validation.Required, is.Int),
		validation.Field(&dto.PekerjaanID, validation.Required, is.Int),
		validation.Field(&dto.PendidikanID, validation.Required, is.Int),
		validation.Field(&dto.LayananID, validation.Required, is.Int),
	); err != nil {
		return err
	}
	return nil
}

type UserRespDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
