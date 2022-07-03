package models

/*
 * Author      : Jody (jody.almaida@gmail)
 * Modifier    :
 * Domain      : survey-skm
 */

import (
	"time"
)

type Respondens struct {
	ID           int64     `gorm:"id"`
	Name         string    `gorm:"display_name"`
	Email        string    `gorm:"email_address"`
	Umur         int64     `gorm:"umur"`
	PekerjaanID  int64     `gorm:"pekerjaan_id"`
	PendidikanID int64     `gorm:"pendidikan_id"`
	LayananID    int64     `gorm:"layanan_id"`
	CreatedAt    time.Time `gorm:"created_at"`
	UpdatedAt    time.Time `gorm:"updated_at"`
}

type RespondenDetails struct {
	ID           int64     `gorm:"id"`
	PertanyaanID int64     `gorm:"pertanyaan_id"`
	JawabanID    int64     `gorm:"jawaban_id"`
	Nilai        int64     `gorm:"nilai"`
	CreatedAt    time.Time `gorm:"created_at"`
	UpdatedAt    time.Time `gorm:"updated_at"`
}

type ResultQuisonerByEachLayanan struct {
	Layanan    string `gorm:"layanan"`
	Responden  string `gorm:"respondend"`
	Umur       string `gorm:"umur"`
	Pekerjaan  string `gorm:"pekerjaan"`
	Pendidikan string `gorm:"pendidikan"`
	Nilai      int64  `gorm:"nilai"`
}
