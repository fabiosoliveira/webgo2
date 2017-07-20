package controllers

import (
	"encoding/json"
	"github.com/fabiosoliveira/webgo/models"
	"github.com/astaxie/beego/validation"

	_ "github.com/fabiosoliveira/webgo/validation"

	"github.com/astaxie/beego"

	"strconv"
)

type ContatoController struct {
	beego.Controller
}

func (c *ContatoController) GetAll() {
	c.Data["json"] = models.GetContatos()
	c.ServeJSON()
}

func (c *ContatoController) Get() {
	if id, err := c.GetInt(":id"); err != nil {
		c.Data["json"] = map[string]interface{}{"ContatoId": "erro, tipo inválido"}
	} else {
		contato := models.GetContato(id)
		c.Data["json"] = contato
	}
	c.ServeJSON()
}

func (c *ContatoController) App() {
	c.TplName = "index.html"
	c.Render()
}

func (c *ContatoController) Post() {
	c.tratarEnvioForm(models.AddContato)
}

func (c *ContatoController) Put() {
	c.tratarEnvioForm(models.AlterarContato)
}

func (c *ContatoController) tratarEnvioForm(f func(contato models.Contato) int) {
	var ob models.Contato
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)

	valid := validation.Validation{}
	ob.Valid(&valid) // validação personalizada

	b, _ := valid.Valid(&ob)

	c.TplName = "msg_info.tpl"
	if !b {
		c.Data["Erros"] = valid.Errors
		c.Data["Tipo"] = "danger"
	} else {
		c.Data["Single"] = true
		c.Data["Tipo"] = "success"
		c.Data["Key"] = "objectid " + strconv.Itoa(f(ob))
		c.Data["Message"] = "adicionado com sucesso"
	}
	c.Render()
}

func (c *ContatoController) Delete() {
	c.TplName = "msg_info.tpl"
	if id, err := c.GetInt(":id"); err != nil {
		c.Data["Single"] = true
		c.Data["Tipo"] = "danger"
		c.Data["Key"] = "objectid " + strconv.Itoa(id)
		c.Data["Message"] = "tipo inválido"
	} else {
		models.ApagarContato(id)
		c.Data["Single"] = true
		c.Data["Tipo"] = "success"
		c.Data["Key"] = "objectid " + strconv.Itoa(id)
		c.Data["Message"] = "removido com sucesso"
	}
	c.Render()
}
