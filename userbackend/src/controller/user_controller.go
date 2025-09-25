package controller

import (
	serverdto "test/userbackend/config/server/dto"
)

type UserController struct {
}

func (userController *UserController) FetchAll() serverdto.Response {
	return serverdto.Response{
		ReturnMessage: []string{"This is user controller 111"},
		ReturnCode:    serverdto.RETURN_CODE_OK,
	}
}
