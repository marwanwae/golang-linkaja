package exception

import (
	"belajar-go-rest-api/helper"
	"belajar-go-rest-api/model/web"
	"net/http"

	"github.com/go-playground/validator"
)

func ErrorHandler(writter http.ResponseWriter, request *http.Request, err interface{}) {

	if notFoundError(writter, request, err) {
		return
	}

	if balanceError(writter, request, err) {
		return
	}

	if validationErrors(writter, request, err) {
		return
	}

	internalServerError(writter, request, err)
}

func internalServerError(writter http.ResponseWriter, request *http.Request, err interface{}) {
	writter.Header().Set("Content-Type", "application/json")
	writter.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}

	helper.WriteResponseBody(writter, webResponse)
}

func notFoundError(writter http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)

	if ok {
		writter.Header().Set("Content-Type", "application/json")
		writter.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   exception.Error,
		}

		helper.WriteResponseBody(writter, webResponse)
		return true
	} else {
		return false
	}
}

func balanceError(writter http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(BalanceError)

	if ok {
		writter.Header().Set("Content-Type", "application/json")
		writter.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "TRANSFER FAILED",
			Data:   exception.Error,
		}

		helper.WriteResponseBody(writter, webResponse)
		return true
	} else {
		return false
	}
}

func validationErrors(writter http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)

	if ok {
		writter.Header().Set("Content-Type", "application/json")
		writter.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error(),
		}

		helper.WriteResponseBody(writter, webResponse)
		return true
	} else {
		return false
	}
}
