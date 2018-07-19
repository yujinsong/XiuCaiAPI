package routers

import (
	"github.com/gorilla/mux"
	"github.com/XiuCai/XiuCaiAPI/controllers"
	"github.com/codegangsta/negroni"
)

func SetUserRouters(router *mux.Router) *mux.Router {
	userRouter := mux.NewRouter()
	userRouter.HandleFunc("/Xiucai/User/checkCode", controllers.CheckCode).Methods("POST")
	userRouter.HandleFunc("/Xiucai/User/checkSession", controllers.CheckSession).Methods("POST")

	router.PathPrefix("/User").Handler(negroni.New(
		//negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(userRouter),
	))
	return userRouter
}

