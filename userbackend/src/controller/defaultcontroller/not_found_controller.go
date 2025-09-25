package defaultcontroller

import (
	serverdto "test/userbackend/config/server/dto"
)

type DefaultController struct {
}

func (defaultController *DefaultController) FetchAll() serverdto.Response{
	return serverdto.Response{
		ReturnMessage: []string{"404 NOT FOUND"},
		ReturnCode: serverdto.RETURN_CODE_NOT_FOUND,
	}
}