package handlers

import "github.com/ainmtsn1999/digitalent_final_hacktiv8/services"

type HttpServer struct {
	app services.ServiceInterface
}

func NewHttpServer(app services.ServiceInterface) HttpServer {
	return HttpServer{app: app}
}
