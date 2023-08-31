package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AccountController interface {
	FindyByAccount(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Transfer(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
