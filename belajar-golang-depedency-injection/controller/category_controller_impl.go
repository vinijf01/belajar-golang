package controller

import (
	"belajar-golang-depedency-injection/helper"
	"belajar-golang-depedency-injection/model/web"
	"belajar-golang-depedency-injection/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) *CategoryControllerImpl {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(r, &categoryCreateRequest)

	categoryRespon := controller.CategoryService.Create(r.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "ok",
		Data:   categoryRespon,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(r, &categoryUpdateRequest)

	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryUpdateRequest.Id = id

	categoryRespon := controller.CategoryService.Update(r.Context(), categoryUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "ok",
		Data:   categoryRespon,
	}

	helper.WriteToResponseBody(w, webResponse)

}

func (controller *CategoryControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	controller.CategoryService.Delete(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "ok",
	}

	helper.WriteToResponseBody(w, webResponse)

}

func (controller *CategoryControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	response := controller.CategoryService.FindById(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "ok",
		Data:   response,
	}

	helper.WriteToResponseBody(w, webResponse)

}

func (controller *CategoryControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryResponses := controller.CategoryService.FindAll(r.Context())

	webResponse := web.WebResponse{
		Code:   200,
		Status: "ok",
		Data:   categoryResponses,
	}

	helper.WriteToResponseBody(w, webResponse)

}
