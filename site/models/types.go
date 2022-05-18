package models

type Cliente struct {
	Contatos []Contato `json:"contacts"`
}

type Contato struct {
	Nome    string `json:"name"`
	Celular string `json:"cellphone"`
}

type Usuario struct {
	ID    int64  `json:"id,omitempty"`
	Nome  string `json:"nome,omitempty"`
	Senha string `json:"senha,omitempty"`
}
