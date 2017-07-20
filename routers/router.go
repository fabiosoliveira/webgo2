package routers

import (
	"github.com/fabiosoliveira/webgo2/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.ContatoController{}, "get:App")
	beego.Router("/contato", &controllers.ContatoController{}, "get:GetAll;post:Post;put:Put")
	beego.Router("/contato/:id", &controllers.ContatoController{}, "delete:Delete;get:Get")
}
