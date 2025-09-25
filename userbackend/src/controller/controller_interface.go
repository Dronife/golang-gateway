package controller

import (
	serverdto "test/userbackend/config/server/dto"
)

type ControllerInterface interface {
	FetchAll() serverdto.Response
	Fetch(id int) serverdto.Response
	Detele(id int) serverdto.Response
}