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

func (repo *respondenRepository) GetQuisionerDataEachLayanan(ctx context.Context) ([]*models.ResultQuisonerByEachLayanan, error) {
	data := []*models.ResultQuisonerByEachLayanan{}

	q := repo.connection.WithContext(ctx)

	if err := q.Raw(`select 
	lyn.display_name as layanan,
	r.display_name as responden,
	r.umur,
	pk.display_name as pekerjaan,
	pd.display_name as pendidikan,
	sum(rd.nilai) as nilai
	from
	master.layanan_opds lyn 
	JOIN survey.respondens r ON lyn.id = r.layanan_id
	JOIN survey.responden_details rd ON r.id = rd.responden_id
	JOIN master.pekerjaans pk ON pk.id = r.pekerjaan_id
	JOIN master.pendidikans pd ON pd.id = r.pendidikan_id
	GROUP BY r.layanan_id, lyn.display_name, lyn.id, r.display_name, r.umur, pk.display_name, pd.display_name
	ORDER BY lyn.id`).Scan(&data).Error; err != nil {

		log.Println(err)
		return nil, err
	}

	return data, nil
}
