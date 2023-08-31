package app

import (
	"belajar-go-rest-api/controller"
	"belajar-go-rest-api/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(AccountController controller.AccountController) *httprouter.Router {

	router := httprouter.New()

	router.GET("/api/account/:account", AccountController.FindyByAccount)
	router.POST("/api/account/:account/transfer", AccountController.Transfer)

	router.PanicHandler = exception.ErrorHandler

	return router
}
