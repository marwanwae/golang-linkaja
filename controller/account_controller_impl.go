package controller

import (
	"belajar-go-rest-api/helper"
	"belajar-go-rest-api/model/web"
	"belajar-go-rest-api/services"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type AccountControllerImpl struct {
	AccountService services.AccountService
}

func NewAccountController(AccountService services.AccountService) AccountController {
	return &AccountControllerImpl{
		AccountService: AccountService,
	}
}

func (controller *AccountControllerImpl) Transfer(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	AccountId := params.ByName("account")
	account, err3 := strconv.Atoi(AccountId)
	helper.PanicIfError(err3)

	AccountCreateRequest := web.TransferRequest{}
	helper.ReadFromRequestBody(request, &AccountCreateRequest)

	controller.AccountService.Transfer(request.Context(), account, AccountCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "ok",
		Data:   nil,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *AccountControllerImpl) FindyByAccount(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	AccountId := params.ByName("account")
	account, err3 := strconv.Atoi(AccountId)
	helper.PanicIfError(err3)

	AccountRespnse := controller.AccountService.FindyById(request.Context(), account)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "ok",
		Data:   AccountRespnse,
	}

	helper.WriteResponseBody(writer, webResponse)
}
