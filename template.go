package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"text/template"
)

type resume struct {
	About      string
	Experience []struct {
		Company string
		Team    string
		Start   string
		End     string
		Tasks   []struct {
			Overview  string
			Breakdown []string
		}
	}
	Technology []struct {
		Name        string
		Proficiency int
	}
}

func main() {

	resumeData, err := ioutil.ReadFile("resume.json")
	if err != nil {
		panic(err)
	}

	r := new(resume)
	if err = json.Unmarshal(resumeData, r); err != nil {
		panic(err)
	}

	tmpl, err := template.ParseFiles("index.tmpl")
	if err != nil {
		panic(err)
	}

	var renderedTmpl bytes.Buffer
	if err = tmpl.Execute(&renderedTmpl, r); err != nil {
		panic(err)
	}

	if err := ioutil.WriteFile("site/index.html", renderedTmpl.Bytes(), 0644); err != nil {
		panic(err)
	}
}
