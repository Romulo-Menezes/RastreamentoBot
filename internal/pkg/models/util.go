package models

import (
	"log"
	"regexp"
)

func CheckCode(code string) bool {
	result, err := regexp.Match("([a-z]|[A-Z]){2}([0-9]){9}([a-z]|[A-Z]){2}", []byte(code))
	if err != nil {
		log.Printf("Ocorreu um erro ao verificar o codigo: %v\n", err)
	}
	return result
}
func GetLastModified(code string) string {
	obj := GetTrackingInformation(code)
	if obj.Events == nil {
		return ""
	}
	return obj.Events[0].CreationTime
}

func GetResume(code string) Event {
	obj := GetTrackingInformation(code)
	if obj.Events == nil {

	}
	return obj.Events[0]
}
