package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:BancoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:BancoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:BancoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:BancoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:BancoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:BancoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:BancoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:BancoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:BancoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:BancoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:CiiuClaseController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:CiiuClaseController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:CiiuClaseController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:CiiuClaseController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:CiiuClaseController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:CiiuClaseController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:CiiuClaseController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:CiiuClaseController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:CiiuClaseController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:CiiuClaseController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:CiiuDivisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:CiiuDivisionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:CiiuDivisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:CiiuDivisionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:CiiuDivisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:CiiuDivisionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:CiiuDivisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:CiiuDivisionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:CiiuDivisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:CiiuDivisionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:CiiuSubclaseController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:CiiuSubclaseController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:CiiuSubclaseController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:CiiuSubclaseController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:CiiuSubclaseController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:CiiuSubclaseController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:CiiuSubclaseController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:CiiuSubclaseController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:CiiuSubclaseController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:CiiuSubclaseController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:CiiuTipoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:CiiuTipoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:CiiuTipoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:CiiuTipoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:CiiuTipoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:CiiuTipoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:CiiuTipoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:CiiuTipoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:CiiuTipoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:CiiuTipoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:CiudadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:CiudadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:CiudadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:CiudadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:CiudadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:CiudadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:CiudadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:CiudadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:CiudadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:CiudadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:JefeDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:JefeDependenciaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:JefeDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:JefeDependenciaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:JefeDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:JefeDependenciaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:JefeDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:JefeDependenciaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:JefeDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:JefeDependenciaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:OrdenadorGastoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:OrdenadorGastoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:OrdenadorGastoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:OrdenadorGastoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:OrdenadorGastoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:OrdenadorGastoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:OrdenadorGastoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:OrdenadorGastoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:OrdenadorGastoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:OrdenadorGastoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:PuntoSalarialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:PuntoSalarialController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:PuntoSalarialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:PuntoSalarialController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:PuntoSalarialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:PuntoSalarialController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:PuntoSalarialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:PuntoSalarialController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:PuntoSalarialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:PuntoSalarialController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:RubrosDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:RubrosDependenciaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:RubrosDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:RubrosDependenciaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:RubrosDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:RubrosDependenciaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:RubrosDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:RubrosDependenciaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:RubrosDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:RubrosDependenciaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:RubrosOrdenadorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:RubrosOrdenadorController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:RubrosOrdenadorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:RubrosOrdenadorController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:RubrosOrdenadorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:RubrosOrdenadorController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:RubrosOrdenadorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:RubrosOrdenadorController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:RubrosOrdenadorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:RubrosOrdenadorController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:SalarioMinimoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:SalarioMinimoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:SalarioMinimoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:SalarioMinimoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:SalarioMinimoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:SalarioMinimoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:SalarioMinimoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:SalarioMinimoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:SalarioMinimoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:SalarioMinimoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:SniesAreaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:SniesAreaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:SniesAreaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:SniesAreaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:SniesAreaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:SniesAreaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:SniesAreaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:SniesAreaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:SniesAreaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:SniesAreaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:SniesNucleoBasicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:SniesNucleoBasicoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:SniesNucleoBasicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:SniesNucleoBasicoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:SniesNucleoBasicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:SniesNucleoBasicoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:SniesNucleoBasicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:SniesNucleoBasicoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:SniesNucleoBasicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:SniesNucleoBasicoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:SucursalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:SucursalController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:SucursalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:SucursalController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:SucursalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:SucursalController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:SucursalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:SucursalController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:SucursalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:SucursalController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:TipoDocumentoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:TipoDocumentoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:TipoDocumentoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:TipoDocumentoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:TipoDocumentoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:Tipo_entidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:Tipo_entidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:Tipo_entidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:Tipo_entidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:Tipo_entidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:Tipo_entidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:Tipo_entidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:Tipo_entidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:Tipo_entidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_amazon_crud/controllers:Tipo_entidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
