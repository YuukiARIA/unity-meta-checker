### Dangling .meta paths
{{- range .DanglingMetaPaths }}
- {{ . }}
{{- end }}

### Asset paths without .meta
{{- range .MetalessAssetPaths }}
- {{ . }}
{{- end }}
