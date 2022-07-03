package responden_handlers

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : survey-skm
 */

import (
	"encoding/json"
	"net/http"

	dtoResponden "survey-skm/src/app/dtos/respondens"
	usecases "survey-skm/src/app/use_cases/responden"

	_ "net/http/pprof"
	common_error "survey-skm/src/infra/errors"
	"survey-skm/src/interface/rest/response"
)

type RespondenHandlerInterface interface {
	CreateQuisionerData(w http.ResponseWriter, r *http.Request)
	GetQuisionerDataEachLayanan(w http.ResponseWriter, r *http.Request)
}

type respondenHandler struct {
	response response.IResponseClient
	usecase  usecases.ResondenUsecaseInterface
}

func NewRespondenHandler(r response.IResponseClient, u usecases.ResondenUsecaseInterface) RespondenHandlerInterface {
	return &respondenHandler{
		response: r,
		usecase:  u,
	}
}

func (h *respondenHandler) CreateQuisionerData(w http.ResponseWriter, r *http.Request) {

	postDTO := dtoResponden.RespondenReqDTO{}
	err := json.NewDecoder(r.Body).Decode(&postDTO)
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}

	err = postDTO.Validate()
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}

	err = h.usecase.CreateQuisionerData(r.Context(), &postDTO)
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.UNKNOWN_ERROR, err))
		return
	}

	h.response.JSON(
		w,
		"Successful Create Quisioner",
		nil,
		nil,
	)
}

func (h *respondenHandler) GetQuisionerDataEachLayanan(w http.ResponseWriter, r *http.Request) {

	data, err := h.usecase.GetQuisionerDataEachLayanan(r.Context())
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.UNKNOWN_ERROR, err))
		return
	}

	h.response.JSON(
		w,
		"Successful Get Quisioner Data",
		data,
		nil,
	)
}
