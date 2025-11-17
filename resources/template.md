{{range .Files}}
---
# {{.Path}}
---
{{.GetContent}}
---
{{end}}