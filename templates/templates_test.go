package templates

import (
	"os"
	"testing"
	"text/template"
)

func TestLearnTemplates(t *testing.T) {
	// Examples are mostly from https://golang.org/pkg/text/template/ documentation.

	t.Run("Trivial", func(t *testing.T) {
		type Inventory struct {
			Material string
			Count    uint
		}
		sweaters := Inventory{"wool", 17}

		tmpl, err := template.New("test").
			Parse("{{.Count}} items are made of {{.Material}}.\n")
		if err != nil {
			t.Fatalf("parse template: %v", err)
		}

		if err := tmpl.Execute(os.Stdout, sweaters); err != nil {
			t.Fatalf("exec template: %v", err)
		}
	})

	t.Run("TrimWhitespaces", func(t *testing.T) {
		tmpl, err := template.New("test").
			Parse("{{23 -}} < {{- 35}}\n")
		if err != nil {
			t.Fatalf("parse template: %v", err)
		}

		if err := tmpl.Execute(os.Stdout, nil); err != nil {
			t.Fatalf("exec template: %v", err)
		}
	})

	t.Run("MyPipelineFunc", func(t *testing.T) {
		funcMap := template.FuncMap{
			"myDouble": myDouble,
		}

		tmpl, err := template.New("test").
			Funcs(funcMap).
			Parse("Double of 2 is {{ myDouble 2 }}.\n")
		if err != nil {
			t.Fatalf("parse template: %v", err)
		}

		if err := tmpl.Execute(os.Stdout, myDouble); err != nil {
			t.Fatalf("exec template: %v", err)
		}
	})

	t.Run("MyPipelineMethod", func(t *testing.T) {
		tmpl, err := template.New("test").
			Parse("Double of 2 is {{ .Double 2 }}.\n")
		if err != nil {
			t.Fatalf("parse template: %v", err)
		}

		if err := tmpl.Execute(os.Stdout, &MyPipelineMethodType{}); err != nil {
			t.Fatalf("exec template: %v", err)
		}
	})

	// A lot of pipelines are skipped. Add them later

	t.Run("Define", func(t *testing.T) {
		tmpl, err := template.New("test").
			Parse(`{{define "title"}}<h1>{{.}}</h1>{{end}}<html>
	<head></head>
	<body>
		{{template "title" .Title1}}<br>
		<p>some simple data</p><br>
		{{template "title" .Title2}}<br>
	</body>
</html>
`)
		if err != nil {
			t.Fatalf("parse template: %v", err)
		}

		if err := tmpl.Execute(os.Stdout, map[string]string{
			"Title1": "Title1", "Title2": "Title2",
		}); err != nil {
			t.Fatalf("exec template: %v", err)
		}
	})

	t.Run("ParseTree", func(t *testing.T) {
		title, err := template.New("title").
			Parse(`<h1>{{.}}</h1>`)
		if err != nil {
			t.Fatalf("parse template: %v", err)
		}

		mainTmpl, err := template.New("Main").
			Parse(`<html>
	<head></head>
	<body>
		{{template "title" .Title1}}<br>
		<p>some simple data</p><br>
		{{template "title" .Title2}}<br>
	</body>
</html>
	`)
		if err != nil {
			t.Fatalf("parse template: %v", err)
		}

		if mainTmpl, err = mainTmpl.AddParseTree("title", title.Tree); err != nil {
			t.Fatalf("parse template: %v", err)
		}

		if err := mainTmpl.ExecuteTemplate(os.Stdout, "Main", map[string]string{
			"Title1": "Title1", "Title2": "Title2",
		}); err != nil {
			t.Fatalf("exec template: %v", err)
		}

	})
}

type MyPipelineMethodType struct{}

func (*MyPipelineMethodType) Double(i int) int {
	return myDouble(i)
}

func myDouble(i int) int {
	return 2 * i
}
