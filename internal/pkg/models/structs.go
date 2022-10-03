package models

import "fmt"

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
	Message string  `json:"mensagem"`
	Events  []Event `json:"eventos"`
}

type Objects struct {
	Objects []Object `json:"objetos"`
}

func (ad *address) toString() string {
	return fmt.Sprintf("%v - %v", ad.City, ad.UF)
}

func (l *localization) toString() string {
	return fmt.Sprintf("%v: %v", l.Type, l.Address.toString())
}

func (obj *Object) ToString(i int) string {
	return fmt.Sprintf("Data: %v\n%v\nStatus: %v", FormatDate(*obj), obj.Events[i].Localization.toString(), obj.Events[i].Description)
}
