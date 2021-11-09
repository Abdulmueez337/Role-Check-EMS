package handlers

import (
	"ROles-Based/gen/restapi/operations/getting_role"
	"ROles-Based/service"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
)

type rolebase struct {
}

func RoleBase() getting_role.GetroleHandler {
	return &rolebase{}
}

func (rb *rolebase) Handle(params getting_role.GetroleParams) middleware.Responder {
	fmt.Println("Designation",params.Designation)
	fmt.Println("apiName",params.APIName)
	rbase,rmessage := service.CheckRole(params.Designation,params.APIName)
	if rbase == 200{
		return getting_role.NewGetroleOK().WithPayload(toUserGen(rbase,rmessage))
	}else if rbase == 403{
		return getting_role.NewGetroleForbidden().WithPayload(toUserGen2(rbase,rmessage))
	}else if rbase == 400 {
		return getting_role.NewGetroleBadRequest().WithPayload(toUserGen3(rbase,rmessage))
	}else {
		return getting_role.NewGetroleInternalServerError().WithPayload(toUserGen1(rbase,rmessage))
	}
}


