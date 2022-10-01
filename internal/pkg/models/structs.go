package models

type address struct {
	City string `json:"cidade"`
	UF   string `json:"uf"`
}

type localization struct {
	Address address `json:"endereco"`
	Type    string  `json:"tipo"`
}

type Event struct {
	Code           string       `json:"codigo"`
	Description    string       `json:"descricao"`
	CreationTime   string       `json:"dtHrCriado"`
	Localization   localization `json:"unidade"`
	TargetLocation localization `json:"unidadeDestino"`
}

type Object struct {
	Events []Event `json:"eventos"`
}

type Objects struct {
	Objects []Object `json:"objetos"`
}
