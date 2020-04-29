package models

type People struct {
	Name     string `schema:"nome" json:"nome"`
	MName    string `schema:"nome_mae" json:"nome_mae"`
	CPF      string `schema:"cpf" json:"cpf"`
	RG       string `schema:"rg" json:"rg"`
	Sex      string `schema:"sexo" json:"sexo"`
	DataNasc string `schema:"data_nasc" json:"data_nasc"`
}
