package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetTrackingInformation(code string) Object {
	var obj Objects

	link := fmt.Sprintf("https://proxyapp.correios.com.br/v1/sro-rastro/%s", code)

	response, err := http.Get(link)
	if err != nil {
		log.Fatalf("Erro ao buscar na api! %s", err.Error())
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Erro ao tratar a resposta! %s", err.Error())
	}

	json.Unmarshal(responseData, &obj)

	return obj.Objects[0]
}
