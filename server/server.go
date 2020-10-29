package server

import (
	"bytes"
	"fmt"
	"gopkg.in/yaml.v2"
	"html/template"
	"net/http"
)

type data struct {
	Variables string
	Template  string
	Result    string
	Error     error
}

type Values struct {
	items interface{}
}

var t *template.Template

func New(port string) error {
	var err error
	t, err = template.ParseFiles("server/template.html")

	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/render", renderHandler)

	return http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

func mainHandler(writer http.ResponseWriter, _ *http.Request) {
	if err := t.Execute(writer, data{}); err != nil {
		panic(err)
	}
}

func renderHandler(writer http.ResponseWriter, request *http.Request) {
	userVariablesStr := request.PostFormValue("variables")
	userTemplateStr := request.PostFormValue("template")
	var userVariables interface{}

	if err := yaml.Unmarshal([]byte(userVariablesStr), &userVariables); err != nil {
		panic(err)
	}

	userTemplate, err := template.New("userTemplate").Parse(userTemplateStr)

	if err != nil {
		panic(err)
	}

	data := data{
		Variables: userVariablesStr,
		Template:  userTemplateStr,
	}

	var tpl bytes.Buffer
	if err := userTemplate.Execute(&tpl, userVariables); err != nil {
		data.Error = err
	}

	data.Result = tpl.String()

	if err := t.Execute(writer, data); err != nil {
		panic(err)
	}
}
