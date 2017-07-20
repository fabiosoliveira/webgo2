package models

import (
	"regexp"

	"github.com/astaxie/beego/validation"
)

type Contato struct {
	Id       int    `json:"id"`
	Nome     string `json:"nome" valid:"Required;MinSize(3);MaxSize(60)"`
	Telefone string `json:"telefone"`
	Email    string `json:"email" valid:"Email"`
}

func (c *Contato) Valid(v *validation.Validation) {
	if m, _ := regexp.MatchString("^[a-zA-ZãÃáÁàÀêÊéÉèÈíÍìÌôÔõÕóÓòÒúÚùÙûÛçÇºª' ']+$", c.Nome); !m {
		v.SetError("Nome.Alpha", "Deve ser caracteres alfabéticos válidos")
	}

	if m, _ := regexp.MatchString("\\([0-9]{2}\\)[2-9](([0-9]{3})|([0-9]{4}))-[0-9]{4}", c.Telefone); !m {
		v.SetError("Telefone.Phone", "Deve ser número de telefone ou celular móvel válido")
	}
}
