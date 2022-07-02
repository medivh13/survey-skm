package postgres

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : survey-skm
 */

import (
	"context"
	"log"

	dto "survey-skm/src/app/dtos/respondens"
	repositories "survey-skm/src/domain/repositories"
	models "survey-skm/src/infra/models"

	"gorm.io/gorm"
)

type respondenRepository struct {
	connection *gorm.DB
}

func NewRespondenRepository(db *gorm.DB) repositories.RespondenRepository {
	return &respondenRepository{
		connection: db,
	}
}

func (repo *respondenRepository) CreateQuisionerData(ctx context.Context, data *dto.RespondenReqDTO) error {
	respondenModel := models.Respondens{}
	respondenDetailModel := models.RespondenDetails{}
	q := repo.connection.WithContext(ctx)
	tx := q.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Raw(`INSERT INTO
	survey.respondens (
		display_name, email_address, umur, pekerjaan_id, pendidikan_id, layanan_id) 
		VALUES (?,?,?,?,?,?) RETURNING id`, data.Name, data.Email, data.Umur, data.PekerjaanID, data.PendidikanID, data.LayananID).Scan(&respondenModel).Error; err != nil {
		tx.Rollback()
		log.Println(err)
		return err
	}

	for _, val := range data.Quisioner {
		if err := tx.Raw(`INSERT INTO
	survey.responden_details (
		responden_id, pertanyaan_id, jawaban_id, nilai) 
		VALUES (?,?,?,?)`, respondenModel.ID, val.PertanyaanID, val.JawabanID, val.Nilai).Scan(&respondenDetailModel).Error; err != nil {
			tx.Rollback()
			log.Println(err)
			return err
		}
	}

	return tx.Commit().Error
}
