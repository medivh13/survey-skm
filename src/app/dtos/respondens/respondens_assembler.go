package respondens_dto

import (
	"fmt"
	models "survey-skm/src/infra/models"
)

func ToReturnQuisonerData(d []*models.ResultQuisonerByEachLayanan) []*QuisionerEachLayananRespDTO {
	fmt.Println(d)
	dataMap := make(map[string]*QuisionerEachLayananRespDTO)
	var resp []*QuisionerEachLayananRespDTO

	for _, val := range d {
		if _, ok := dataMap[val.Layanan]; !ok {
			dataMap[val.Layanan] = &QuisionerEachLayananRespDTO{
				Layanan: val.Layanan,
			}
			dataMap[val.Layanan].Result = append(dataMap[val.Layanan].Result, &DataResultEachLayananRespDTO{
				Responden:  val.Responden,
				Umur:       val.Umur,
				Pekerjaan:  val.Pekerjaan,
				Pendidikan: val.Pendidikan,
				Nilai:      val.Nilai,
			})
		} else {
			dataMap[val.Layanan].Result = append(dataMap[val.Layanan].Result, &DataResultEachLayananRespDTO{
				Responden:  val.Responden,
				Umur:       val.Umur,
				Pekerjaan:  val.Pekerjaan,
				Pendidikan: val.Pendidikan,
				Nilai:      val.Nilai,
			})
		}
	}

	for _, val := range dataMap {
		fmt.Println(val)
		resp = append(resp, val)
	}

	return resp

}
