package server

import (
	serverdto "test/userbackend/config/server/dto"
	"test/userbackend/src/controller"
	"test/userbackend/src/controller/defaultcontroller"
)

func resolveRoute(path string) *serverdto.Route {
	switch path  {
	case "/user":
        uc := controller.UserController{}
        return serverdto.NewRoute(uc.FetchAll) 
	default:
		dc := defaultcontroller.DefaultController{}
        return serverdto.NewRoute(dc.FetchAll)  
	}
}