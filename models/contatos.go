package models

var contatos map[int]Contato
var cont int

func init() {
	cont = 1
	contatos = make(map[int]Contato)
	iniciar()
}

func AddContato(contato Contato) int {
	contato.Id = cont
	contatos[cont] = contato
	cont++
	return contato.Id
}

func GetContato(id int) Contato {
	return contatos[id]
}

func AlterarContato(contato Contato) int {
	contatos[contato.Id] = contato
	return contato.Id
}

func ApagarContato(id int) {
	delete(contatos, id)
}

func GetContatos() map[int]Contato {
	return contatos
}

func iniciar() {
	AddContato(Contato{1, "Carlos", "(75)98822-3030", "Carlos@hotmail.com"})
	AddContato(Contato{2, "Tássia Oliveira", "(75)98822-3930", "tassia@hotmail.com"})
	AddContato(Contato{3, "Fábio Souza", "(75)98822-3930", "fabio@gmail.com"})
}
