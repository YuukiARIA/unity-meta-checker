package render

import (
	"io"
	"text/template"

	"github.com/YuukiARIA/unity-meta-checker/models"
)

const defaultTemplateContent = `### Dangling .meta paths
{{- range .DanglingMetaPaths }}
- {{ . }}
{{- end }}

### Asset paths without .meta
{{- range .MetalessAssetPaths }}
- {{ . }}
{{- end }}
`

func GetDefaultTemplate() *template.Template {
	return template.Must(template.New("default").Parse(defaultTemplateContent))
}

func LoadTemplate(path string) (*template.Template, error) {
	return template.ParseFiles(path)
}

func RenderResult(result *models.Result, t *template.Template, output io.Writer) error {
	return t.Execute(output, result)
}
