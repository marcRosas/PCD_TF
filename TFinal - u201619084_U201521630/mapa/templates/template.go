package templates

import (
	"bytes"
	"io/ioutil"
	"text/template"
)

const TemplatePath = "templates/"

func getTemplateFile(templateName string) (templateData string, err error) {
	data, err := ioutil.ReadFile(TemplatePath + templateName)
	if err != nil {
		return
	}
	templateData = string(data[:])
	return
}

func GetTemplate(templateName string) (templ *template.Template, err error) {
	templateData, err := getTemplateFile(templateName)
	if err != nil {
		return
	}
	templ, err = template.New(templateName).Parse(templateData)
	return
}

func GetTemplateWithData(templateName string, obj interface{}) (data []byte, err error) {
	templateData, err := GetTemplate(templateName)
	if err != nil {
		return
	}
	var buff bytes.Buffer
	err = templateData.Execute(&buff, obj)
	if err != nil {
		return
	}
	data = buff.Bytes()

	return
}
