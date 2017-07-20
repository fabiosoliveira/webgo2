package validation

import "github.com/astaxie/beego/validation"

func init() {
	validation.SetDefaultMessage(map[string]string{
		"Required":     "Não pode esta vazio",
		"Min":          "Minimo é %d",
		"Max":          "Máximo é %d",
		"Range":        "Intervalo é %d para %d",
		"MinSize":      "O tamano mínimo é %d",
		"MaxSize":      "O tamanho máximo é %d",
		"Length":       "O comprimento necessário é %d",
		"Alpha":        "Deve ser caractere alfa válido",
		"Numeric":      "Deve ser caractere numérico válido",
		"AlphaNumeric": "Deve ser caracteres alfa ou numérico válidos",
		"Match":        "Deve corresponder %s",
		"NoMatch":      "Não deve corresponder %s",
		"AlphaDash":    "Deve ser válidos caracteres alfa ou numéricos ou dash(-_) caracteres",
		"Email":        "Deve ser um endereço de email válido",
		"IP":           "Deve ser um endereço IP válido",
		"Base64":       "Deve ser caracteres base64 válidos",
		"Mobile":       "Deve ser número de celular válido",
		"Tel":          "Deve ser número de telefone válido",
		"Phone":        "Deve ser número de telefone ou celular móvel válido",
		"ZipCode":      "Deve ser um CEP válido",
	})
}
