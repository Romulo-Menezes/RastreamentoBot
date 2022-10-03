package models

import (
	"fmt"
	"log"
	"regexp"
	"time"
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

func FormatDate(obj Object) string {
	layout := "2006-01-02T15:04:05"
	data, err := time.Parse(layout, obj.Events[0].CreationTime)
	if err != nil {
		log.Printf("Ocorreu um erro ao coverter a data: %v\n", err)
	}
	return data.String()[:19]
}

func GetResume(name string, code string) (title string, description string) {
	obj := GetTrackingInformation(code)

	title = fmt.Sprintf("%v - %v", name, code)

	if obj.Events == nil {
		description = obj.Message
	} else {
		description = fmt.Sprintf("%v\n", obj.ToString(0))
	}
	return title, description
}

func GetHistory(name string, code string) (title string, description string) {
	title, description = GetResume(name, code)
	obj := GetTrackingInformation(code)

	for i := 1; i < len(obj.Events); i++ {
		description = fmt.Sprintf("%v\n%v\n", description, obj.ToString(i))
	}
	return title, description
}
