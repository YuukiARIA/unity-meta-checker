package render

import (
	_ "embed"
	"io"
	"text/template"

	"github.com/YuukiARIA/unity-meta-checker/models"
)

//go:embed default_template.tpl
var defaultTemplateContent []byte

func GetDefaultTemplate() *template.Template {
	return template.Must(template.New("default").Parse(string(defaultTemplateContent)))
}

func LoadTemplate(path string) (*template.Template, error) {
	return template.ParseFiles(path)
}

func RenderResult(result *models.Result, t *template.Template, output io.Writer) error {
	return t.Execute(output, result)
}
