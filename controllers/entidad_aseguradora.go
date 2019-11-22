package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/core_amazon_crud/models"
)

// EntidadAseguradoraController operations for EntidadAseguradora
type EntidadAseguradoraController struct {
	beego.Controller
}

// URLMapping ...
func (c *EntidadAseguradoraController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create EntidadAseguradora
// @Param	body		body 	models.EntidadAseguradora	true		"body for EntidadAseguradora content"
// @Success 201 {object} models.EntidadAseguradora
// @Failure 400 the request contains incorrect syntax
// @router / [post]
func (c *EntidadAseguradoraController) Post() {
	var v models.EntidadAseguradora
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddEntidadAseguradora(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			logs.Error(err)
			//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
			c.Data["system"] = err
			c.Abort("400")
		}
	} else {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("400")
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get EntidadAseguradora by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.EntidadAseguradora
// @Failure 404 not found resource
// @router /:id [get]
func (c *EntidadAseguradoraController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get EntidadAseguradora
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.EntidadAseguradora
// @Failure 403
// @router / [get]
func (c *EntidadAseguradoraController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the EntidadAseguradora
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.EntidadAseguradora	true		"body for EntidadAseguradora content"
// @Success 200 {object} models.EntidadAseguradora
// @Failure 403 :id is not int
// @router /:id [put]
func (c *EntidadAseguradoraController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the EntidadAseguradora
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *EntidadAseguradoraController) Delete() {

}
