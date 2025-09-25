package serverdto

type Route struct {
	ControllerMethod func() Response
}

func NewRoute(controllerMethod func() Response) *Route {
	return &Route{ControllerMethod: controllerMethod}
}
